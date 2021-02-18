package sorter

import (
	"d2-item-sorter/pkg/config"
	"d2-item-sorter/pkg/data"
	"d2-item-sorter/pkg/domain"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Bool(v bool) *bool       { return &v }
func Int(v int) *int          { return &v }
func String(v string) *string { return &v }

type SortFunctionType func(items []domain.Item) func(i, j int) bool

var sortModeFunctionMap = map[data.SortMode]SortFunctionType{
	data.SortModeName: func(items []domain.Item) func(i, j int) bool {
		return func(i, j int) bool {
			criteriaA := items[i].Name
			criteriaB := items[j].Name
			return criteriaA < criteriaB
		}
	},
	data.SortModeRuneNumber: func(items []domain.Item) func(i, j int) bool {
		return func(i, j int) bool {
			criteriaA := getRuneNumber(items[i].Type.ID)
			criteriaB := getRuneNumber(items[j].Type.ID)
			return criteriaA < criteriaB
		}
	},
	data.SortModeGemQuality: func(items []domain.Item) func(i, j int) bool {
		return func(i, j int) bool {
			criteriaA := data.GemSortMap[items[i].SubCategory]
			criteriaB := data.GemSortMap[items[j].SubCategory]
			return criteriaA < criteriaB
		}
	},
	data.SortModeItemLevel: func(items []domain.Item) func(i, j int) bool {
		return func(i, j int) bool {

			sortMap := [][]string{
				/*{string(rune(items[i].D2sItem.Identified)), string(rune(items[j].D2sItem.Identified)), string(rune(0))},
				{string(rune(items[i].D2sItem.Ethereal)), string(rune(items[j].D2sItem.Ethereal)), string(rune(0))},
				{string(rune(items[i].D2sItem.Personalized)), string(rune(items[j].D2sItem.Personalized)), string(rune(0))},
				{string(rune(items[i].D2sItem.RunewordID)), string(rune(items[j].D2sItem.RunewordID)), string(rune(1))},
				{string(rune(items[i].D2sItem.SetID)), string(rune(items[j].D2sItem.SetID)), string(rune(0))},
				{string(rune(items[i].D2sItem.NrOfItemsInSockets)), string(rune(items[j].D2sItem.NrOfItemsInSockets)), string(rune(0))},
				{string(rune(items[i].D2sItem.TotalNrOfSockets)), string(rune(items[j].D2sItem.TotalNrOfSockets)), string(rune(0))},
				{string(rune(items[i].D2sItem.Quality)), string(rune(items[j].D2sItem.Quality)), string(rune(0))},
				{string(rune(items[i].D2sItem.LowQualityID)), string(rune(items[j].D2sItem.LowQualityID)), string(rune(0))},
				{string(rune(items[i].D2sItem.TypeID)), string(rune(items[j].D2sItem.TypeID)), string(rune(1))},
				{string(items[i].D2sItem.Type), string(items[j].D2sItem.Type), string(rune(1))},
				{string(rune(items[i].D2sItem.Level)), string(rune(items[j].D2sItem.Level)), string(rune(1))},
				{string(rune(items[i].D2sItem.CurrentDurability)), string(rune(items[j].D2sItem.CurrentDurability)), string(rune(0))},
				{string(rune(items[i].D2sItem.Quantity)), string(rune(items[j].D2sItem.Quantity)), string(rune(0))},*/
			}

			for _, criterias := range sortMap {
				if criterias[0] != criterias[1] {
					if criterias[2] == "0" {
						return criterias[0] < criterias[1]
					} else {
						return criterias[0] > criterias[1]
					}
				}
			}

			return items[i].Type.ID < items[j].Type.ID
		}
	},
}

func isIgnoredPage(name string) bool {
	return strings.HasPrefix(name, "(i)")
}

func isPickupPage(name string) bool {
	return strings.HasPrefix(name, "(p)")
}

func isSharedPickupPage(name string) bool {
	return strings.HasPrefix(name, "(sp)")
}

func getRuneNumber(id string) int {
	runeNumber, err := strconv.ParseInt(id[1:], 10, 0)
	if err != nil {
		return 0
	}
	return int(runeNumber)
}

func SortItems(s domain.Stash, groups []domain.ItemGroup) []domain.StashPage {
	itemsToSort := []domain.Item{}
	itemsToIgnoreCount := 0
	newStashPages := []domain.StashPage{}

	//fmt.Printf("grouping %d pages\n", len(s.stashPages))

	for idx, page := range s.StashPages {

		if isIgnoredPage(page.Name) {
			//ignore this page (just add to new pages without any changes)
			newStashPages = append(newStashPages, page)
			fmt.Printf("\t page %d with %d items will be ignored\n", idx, len(page.Items))
			itemsToIgnoreCount += len(page.Items)
		} else {

			if isPickupPage(page.Name) {
				//this is a page we will pickup items but keep the page itself
				newStashPages = append(newStashPages, domain.NewStashPage(0, page.Name, page.Type, config.GetConfig().StashSize.Width, config.GetConfig().StashSize.Height))
				fmt.Printf("\t page %d with %d items will be used as pickup page\n", idx, len(page.Items))
			}
			//if we get until here it means we take the items of the page and add them to all items (base for sorting)
			itemsToSort = append(itemsToSort, page.Items...)
		}
	}

	fmt.Printf("found %d items (%d to sort and %d to ignore)\n", len(itemsToSort)+itemsToIgnoreCount, len(itemsToSort), itemsToIgnoreCount)
	newStashPages = append(newStashPages, GroupItems(groups, itemsToSort)...)

	return newStashPages
}

func createStashPagesForItems(items []domain.Item, indexType data.StashPageType, name string, sortMode data.SortMode) []domain.StashPage {
	pageCount := 0
	newStashPages := []domain.StashPage{}
	currentStashPage := domain.NewStashPage(pageCount, name, indexType, config.GetConfig().StashSize.Width, config.GetConfig().StashSize.Height)
	pageCount++

	if sortModeFunctionMap[sortMode] != nil {
		sort.SliceStable(items, sortModeFunctionMap[sortMode](items))
	} else {
		sort.SliceStable(items, sortModeFunctionMap[data.SortModeItemLevel](items))
	}

	for _, item := range items {

		isInserted := currentStashPage.InsertItem(item)
		//fmt.Printf("inserted item: %s on page: %v\n", item.Type, isInserted)

		if !isInserted { //means the page is full

			//check previous pages
			newStashPages = append(newStashPages, currentStashPage)
			fmt.Printf("added page %d with %d items\n", currentStashPage.Nr, len(currentStashPage.Items))
			currentStashPage = domain.NewStashPage(pageCount, name, indexType, config.GetConfig().StashSize.Width, config.GetConfig().StashSize.Height)
			pageCount++
			currentStashPage.InsertItem(item)
		}

	}
	if len(currentStashPage.Items) != 0 {
		newStashPages = append(newStashPages, currentStashPage)
		fmt.Printf("added page %d with %d items\n", currentStashPage.Nr, len(currentStashPage.Items))
	}
	return newStashPages
}

func filterMatches(filter domain.GroupFilter, item domain.Item) bool {
	if filter.UniqueItems != nil && *filter.UniqueItems != (item.Quality == domain.ItemQualityUnique) {
		return false
	}

	if filter.SetItems != nil && *filter.SetItems != (item.Quality == domain.ItemQualitySet) {
		return false
	}

	if filter.RareItems != nil && *filter.RareItems != (item.Quality == domain.ItemQualityRare) {
		return false
	}

	if filter.Identified != nil && *filter.Identified != item.Identified {
		return false
	}

	if filter.ExcludedCategories != nil && len(*filter.ExcludedCategories) > 0 {
		for _, filterCategory := range *filter.ExcludedCategories {
			if filterCategory == string(item.Category) ||
				filterCategory == string(item.SubCategory) ||
				filterCategory == string(item.SubSubCategory) {
				return false
			}
		}
	}

	if filter.IncludedCategories != nil && len(*filter.IncludedCategories) > 0 {
		for _, filterCategory := range *filter.IncludedCategories {
			if filterCategory == string(item.Category) ||
				filterCategory == string(item.SubCategory) ||
				filterCategory == string(item.SubSubCategory) {
			} else {
				return false
			}
		}
	}

	if filter.MinFilledSockets != nil && *filter.MinFilledSockets > 0 {
		if int(len(item.SocketedWith)) < *filter.MinFilledSockets {
			return false
		}
	}

	if filter.MaxFilledSockets != nil && *filter.MaxFilledSockets > 0 {
		if int(len(item.SocketedWith)) > *filter.MaxFilledSockets {
			return false
		}
	}

	if filter.MaxTotalSockets != nil && *filter.MaxTotalSockets > 0 {
		if int(item.Sockets) > *filter.MaxTotalSockets {
			return false
		}
	}

	if filter.MinTotalSockets != nil && *filter.MinTotalSockets > 0 {
		if int(item.Sockets) < *filter.MinTotalSockets {
			return false
		}
	}

	return true
}

func belongsToGroup(group domain.ItemGroup, item domain.Item) bool {
	for _, filter := range group.GroupFilter {
		if filterMatches(filter, item) {
			return true
		}
	}
	return false
}

func GroupItems(groups []domain.ItemGroup, items []domain.Item) []domain.StashPage {

	remainingItems := []domain.Item{}
	_ = remainingItems

	for idx := range groups {
		groups[idx].Clear()
	}

	assigned := false
	//assign items to groups
	for _, item := range items {

		/*if item.D2sItem.ID != 0 {
			fmt.Printf("%d %s\n", item.D2sItem.ID, item.getName())
		}*/

		assigned = false
		for idx, group := range groups {

			if belongsToGroup(group, item) {
				groups[idx].AddItem(item)
				assigned = true
				break
			}
		}

		if assigned == false {
			remainingItems = append(remainingItems, item)
			assigned = true
		}
	}

	//sort items in groups

	//create stash pages for each group
	allPages := []domain.StashPage{}
	for _, group := range groups {
		fmt.Printf("\tcreating stash pages for %s and %d items sorted by %s\n", group.Name, len(group.MatchingItems), group.SortMode)
		allPages = append(allPages, createStashPagesForItems(group.MatchingItems, group.IndexType, group.Name, group.SortMode)...)
	}

	//create pages for all remaining items
	fmt.Printf("creating stash pages for %d remaining items\n", len(remainingItems))
	allPages = append(allPages, createStashPagesForItems(remainingItems, data.MainIndexPage, "Misc", data.SortModeItemLevel)...)

	//return all stash pages created
	return allPages
}

func TransferPagesTo(source *domain.Stash, target *domain.Stash) {

	ignoredItems := 0
	_ = ignoredItems
	sourceStashPages := []domain.StashPage{}
	targetStashPages := []domain.StashPage{}

	for idx, page := range source.StashPages {

		if isIgnoredPage(page.Name) {
			//ignore this page (just add to new pages without any changes)
			sourceStashPages = append(sourceStashPages, page)
			fmt.Printf("page %d with %d items will be ignored\n", idx, len(page.Items))
			ignoredItems += len(page.Items)
		} else {

			if isPickupPage(page.Name) {
				//this is a page we will pickup items but keep the page itself
				sourceStashPages = append(sourceStashPages, domain.NewStashPage(0, page.Name, page.Type, config.GetConfig().StashSize.Width, config.GetConfig().StashSize.Height))
				fmt.Printf("page %d with %d items will be used as pickup page\n", idx, len(page.Items))
			}
			//if we get until here it means we take the whole page and put it into the target
			targetStashPages = append(targetStashPages, page)
		}
	}

	//add new (empty) stash pages to our source
	source.StashPages = sourceStashPages
	//move filled pages to our target
	target.StashPages = append(target.StashPages, targetStashPages...)

	fmt.Printf("%d stash pages kept in source and %d pages for target\n", len(source.StashPages), len(targetStashPages))
}
