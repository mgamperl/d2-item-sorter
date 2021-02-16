package internal

import (
	"bufio"
	"d2-item-sorter/pkg/data"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

type Run struct {
	Nr           int            `json:"nr"`
	NewItems     []SortableItem `json:"newItems"`
	RemovedItems []SortableItem `json:"removedItems"`
	Items        []SortableItem `json:"items"`
	StartTime    time.Time      `json:"start"`
	EndTime      time.Time      `json:"end"`
}

func (r *Run) getDuration() time.Duration {
	return r.EndTime.Sub(r.StartTime)
}

var running = false

var globalGroups []ItemGroup
var runTimestamp string
var runNr int = 1

var runs = []Run{}
var previousRun = &Run{Nr: -1}
var currentRun = &Run{Nr: runNr, StartTime: time.Now()}

//SHARED_STASH_FILENAME is the filename of the shared stash file
const SHARED_STASH_FILENAME string = "_LOD_SharedStashSave.sss"

// StartRunRecorder starts the Run Recorder which watches the shared stash to automatically sort all items
func StartRunRecorder(stash *Stash, groups []ItemGroup) bool {
	if running == false {

		running = true
		runTimestamp = time.Now().Format("20060102_150405_000")

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			fmt.Println("ERROR", err)
		}
		defer watcher.Close()

		// read all items one time initially and use this as the "previousRun" to see what items the first run might bring
		stash.Reload()
		stash.SortItems(groups)
		stash.Save("")

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
					//fmt.Printf("EVENT! %v - %#v\n", ok, event)
					if event.Op&fsnotify.Remove == fsnotify.Remove {

						watcher.Remove(stash.Filepath)
						handleRun(stash, groups, runTimestamp)
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

func handleRun(stash *Stash, groups []ItemGroup, backupTimestamp string) {
	currentRun.EndTime = time.Now()

	fmt.Printf("---------------------------------------------------------------\n")
	fmt.Printf(" RUN #%d\n", currentRun.Nr)
	fmt.Printf("---------------------------------------------------------------\n")
	//fmt.Printf("remove operation on %s", stashFilename)

	time.Sleep(time.Millisecond * 200)

	if _, err := os.Stat(stash.Filepath); os.IsNotExist(err) {
		fmt.Printf("%s does not exist", stash.Filepath)
		return
	}

	stash.Reload()
	stash.SortItems(groups)
	stash.Save("")

	currentRun.Items = stash.GetAllItems()

	//TODO: Enable LOOT Filter again
	//lootfilter.WriteBHConfig(utils.GetD2InstallPath(), currentRun.Items)

	newItems, deletedItems := stash.getNewItemsComparedTo(previousRun.Items)

	for _, item := range newItems {

		if newUniqueMap, ok := data.UniqueItems[item.D2sItem.UniqueID]; ok == true {
			fmt.Printf("\tNew Item: %s (%s) - %s %s\n", item.GetName(), item.D2sItem.Type, newUniqueMap.Name, newUniqueMap.ItemType)
		} else {
			fmt.Printf("\tNew Item: %s (%s)\n", item.GetName(), item.D2sItem.Type)
		}
	}

	for _, item := range deletedItems {
		fmt.Printf("\tDeleted Item: %s (%s)\n", item.GetName(), item.D2sItem.Type)
	}

	fmt.Printf("and took %v\n", currentRun.getDuration())

	previousRun = currentRun
	runs = append(runs, *previousRun)

	previousRun.Items = stash.GetAllItems()

	result, _ := json.Marshal(runs)
	file, _ := os.Create("c:/temp/runReport.json")
	file.Write(result)
	file.Close()

	runNr++
	currentRun = &Run{Nr: runNr, StartTime: time.Now(), NewItems: []SortableItem{}, RemovedItems: []SortableItem{}}
}

/*func getAllItems() []SortableItem {

	allItems := []SortableItem{}
	saveDir, err := getSaveDir()
	if err != nil {
		panic("can not find Diablo2 Directory")
	}

	sharedStash := ParseSharedStash(saveDir)
	characters, _ := ParseCharacters(saveDir)
	fmt.Printf("%d characters found\n", len(characters))

	for _, char := range characters {
		allItems = append(allItems, char.GetAllItems()...)
	}

	allItems = append(allItems, sharedStash.GetAllItems()...)

	fmt.Printf("%d items found\n", len(allItems))

	return allItems
}*/

func StopRunRecorder() {
	running = false
}

func ParseSharedStash(saveDir string) Stash {
	stashFilename := filepath.Join(saveDir, SHARED_STASH_FILENAME)
	stash := Stash{}
	stash.ParseStash(stashFilename)
	return stash
}

func backupAndWrite(path string, bytesToWrite []byte, backupSuffix string) error {
	if len(backupSuffix) == 0 {
		backupSuffix = ".backup_" + time.Now().Format("20060102_150405_000")
	}

	os.Rename(path, path+backupSuffix)

	file, err := os.Create(path)

	if err != nil {
		fmt.Printf("Error while opening file to write: %s", path)
		return err
	}

	defer file.Close()

	bfw := bufio.NewWriter(file)

	_, err = bfw.Write(bytesToWrite)

	bfw.Flush()

	return nil
}
