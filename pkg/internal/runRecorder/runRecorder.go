package runRecorder

import (
	"d2-item-sorter/pkg/domain"
	"d2-item-sorter/pkg/internal/reader"
	"d2-item-sorter/pkg/internal/sorter"
	"d2-item-sorter/pkg/internal/writer"
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
)

var running = false
var runTimestamp string

var currentRun *domain.Run = &domain.Run{}
var previousRun *domain.Run = &domain.Run{}
var runNr int = 0
var runs []domain.Run = []domain.Run{}

var recorderChannel = make(chan bool)

func StopRunRecorder() {
	running = false
	recorderChannel <- true
}

func StartRunRecorder(stash *domain.Stash, groups []domain.ItemGroup) bool {
	if running == false {
		running = true
		runTimestamp = time.Now().Format("20060102_150405_000")

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			fmt.Println("ERROR", err)
		}
		defer watcher.Close()

		// read all items one time initially and use this as the "previousRun" to see what items the first run might bring
		stash = sortStash(stash, groups)

		previousRun.Items = stash.GetAllItems()
		currentRun.StartTime = time.Now()

		err = watcher.Add(stash.Filepath)
		if err != nil {
			panic("Couldn't watch stash file")
		}

		done := make(chan bool)
		go executeWatch(watcher, stash, groups)

		<-done

		return true
	} else {
		fmt.Printf("RunRecorder already running\n")
		return false
	}
}

func executeWatch(watcher *fsnotify.Watcher, stash *domain.Stash, groups []domain.ItemGroup) {
	fmt.Printf("RunRecorder Started\n")
	for {
		select {
		case <-recorderChannel:
			fmt.Printf("RunRecorder stopped\n")
			return
		case event, ok := <-watcher.Events:
			_ = ok
			if event.Op&fsnotify.Remove == fsnotify.Remove {

				watcher.Remove(stash.Filepath)
				handleRun(stash, groups)
				watcher.Add(stash.Filepath)

			} else if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Printf("---------------------------------------------------------------")
				fmt.Printf("write operation on %s", stash.Filepath)
			}

		case err := <-watcher.Errors:
			fmt.Println("ERROR", err)
		}
	}
}

func sortStash(prevStash *domain.Stash, groups []domain.ItemGroup) *domain.Stash {
	stash := reader.ReloadAndReturnStash(*prevStash)
	//sharedStash.Reload()
	stash.StashPages = sorter.SortItems(stash, groups)
	writer.WriteStashToFile(stash, "")
	return &stash
}
