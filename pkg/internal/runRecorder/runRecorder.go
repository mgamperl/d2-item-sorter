package runRecorder

import (
	"d2-item-sorter/pkg/domain"
	"d2-item-sorter/pkg/internal/reader"
	"d2-item-sorter/pkg/internal/sorter"
	"d2-item-sorter/pkg/internal/writer"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

var running = false
var runTimestamp string

var currentRun *domain.Run = &domain.Run{}
var previousRun *domain.Run = &domain.Run{}
var runNr int = 0
var runs []domain.Run = []domain.Run{}

func StopRunRecorder() {
	running = false
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
		stash := reader.ReloadAndReturnStash(*stash)
		//sharedStash.Reload()
		stash.StashPages = sorter.SortItems(stash, groups)
		writer.WriteStashToFile(stash, "")

		previousRun.Items = stash.GetAllItems()

		currentRun.StartTime = time.Now()

		err = watcher.Add(stash.Filepath)
		if err != nil {
			panic("Couldn't watch stash file")
		}

		done := make(chan bool)

		go func() {
			fmt.Printf("Run Recorder Started\n")
			for running == true {
				select {
				case <-done:
					fmt.Println("RunRecorder stopped")
					return
				case event, ok := <-watcher.Events:
					_ = ok
					fmt.Printf("running %v \n", running)
					//fmt.Printf("EVENT! %v - %#v\n", ok, event)
					if event.Op&fsnotify.Remove == fsnotify.Remove {

						watcher.Remove(stash.Filepath)
						handleRun(&stash, groups, runTimestamp)
						watcher.Add(stash.Filepath)

					} else if event.Op&fsnotify.Write == fsnotify.Write {
						fmt.Printf("---------------------------------------------------------------")
						fmt.Printf("write operation on %s", stash.Filepath)

						//SortStash(stashFilename, groups, runTimestamp)
					}

				case err := <-watcher.Errors:
					fmt.Println("ERROR", err)
				}
			}
			fmt.Printf("Run Recorder Stopped\n")
		}()

		<-done

		return true
	} else {
		fmt.Printf("RunRecorder already running\n")
		return false
	}
}

func handleRun(stash *domain.Stash, groups []domain.ItemGroup, backupTimestamp string) {
	currentRun.EndTime = time.Now()
	currentRun.Duration = currentRun.EndTime.Sub(currentRun.StartTime)

	fmt.Printf("---------------------------------------------------------------\n")
	fmt.Printf(" RUN #%d\n", currentRun.Counter)
	fmt.Printf("---------------------------------------------------------------\n")
	//fmt.Printf("remove operation on %s", stashFilename)

	time.Sleep(time.Millisecond * 200)

	if _, err := os.Stat(stash.Filepath); os.IsNotExist(err) {
		fmt.Printf("%s does not exist", stash.Filepath)
		return
	}

	s := reader.ReloadAndReturnStash(*stash)
	//sharedStash.Reload()
	s.StashPages = sorter.SortItems(s, groups)
	writer.WriteStashToFile(s, "")

	currentRun.Items = s.GetAllItems()

	//TODO: Enable LOOT Filter again
	//lootfilter.WriteBHConfig(utils.GetD2InstallPath(), currentRun.Items)

	newItems, deletedItems := sorter.CompareItems(currentRun.Items, previousRun.Items)

	for _, item := range newItems {
		fmt.Printf("\tAdded %s Item: %s (%s/%s) in %s\n", item.Quality, item.Name, item.Type.Name, item.Type.Quality, item.Source)
	}

	for _, item := range deletedItems {
		fmt.Printf("\tRemoved %s Item: %s (%s/%s) in %s\n", item.Quality, item.Name, item.Type.Name, item.Type.Quality, item.Source)
	}

	fmt.Printf("%d new and %d deleted items (%d/%d)\n", len(newItems), len(deletedItems), len(currentRun.Items), len(previousRun.Items))
	fmt.Printf("and took %v\n", currentRun.Duration)

	previousRun = currentRun
	runs = append(runs, *previousRun)

	previousRun.Items = s.GetAllItems()

	result, _ := json.Marshal(runs)
	file, _ := os.Create("c:/temp/runReport.json")
	file.Write(result)
	file.Close()

	runNr++
	currentRun = &domain.Run{Counter: runNr, StartTime: time.Now(), NewItems: []domain.Item{}}
}
