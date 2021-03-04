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
)

func handleRun(stash *domain.Stash, groups []domain.ItemGroup) {
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
