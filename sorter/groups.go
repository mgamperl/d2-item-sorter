package sorter

import (
	"io/ioutil"
	"sort"

	"gopkg.in/yaml.v2"
)

func Uint64ToBool(i uint64) bool {
	return i == 1
}

func Bool(v bool) *bool       { return &v }
func Int(v int) *int          { return &v }
func String(v string) *string { return &v }

type GroupFilter struct {
	IncludedCategories *[]string `yaml:"included,omitempty"`
	ExcludedCategories *[]string `yaml:"excluded,omitempty"`
	MinItemLevel       *int      `yaml:"minItemLevel,omitempty"`
	MaxItemLevel       *int      `yaml:"maxItemLevel,omitempty"`
	ItemType           *string   `yaml:"itemType,omitempty"`
	Identified         *bool     `yaml:"identified,omitempty"`
	MinTotalSockets    *int      `yaml:"minTotalSockets,omitempty"`
	MaxTotalSockets    *int      `yaml:"maxTotalSockets,omitempty"`
	MinFilledSockets   *int      `yaml:"minFilledSockets,omitempty"`
	MaxFilledSockets   *int      `yaml:"maxFilledSockets,omitempty"`
	UniqueItems        *bool     `yaml:"uniqueItems,omitempty"`
	SetItems           *bool     `yaml:"setItems,omitempty"`
	RareItems          *bool     `yaml:"rareItems,omitempty"`
	MagicItems         *bool     `yaml:"magicItems,omitempty"`
	NormalItems        *bool     `yaml:"normalItems,omitempty"`
	Etheral            *bool     `yaml:"etheral,omitempty"`
}

type SortMode string

const (
	SortModeName       SortMode = "name"
	SortModeItemLevel  SortMode = "itemLevel"
	SortModeRuneNumber SortMode = "runeNumber"
	SortModeGemQuality SortMode = "gemQuality"
)

type SortFunctionType func(items []SortableItem) func(i, j int) bool

var sortModeFunctionMap = map[SortMode]SortFunctionType{
	SortModeName: func(items []SortableItem) func(i, j int) bool {
		return func(i, j int) bool {
			criteriaA := items[i].GetName()
			criteriaB := items[j].GetName()
			return criteriaA < criteriaB
		}
	},
	SortModeRuneNumber: func(items []SortableItem) func(i, j int) bool {
		return func(i, j int) bool {
			criteriaA := getRuneNumber(items[i].D2sItem.Type)
			criteriaB := getRuneNumber(items[j].D2sItem.Type)
			return criteriaA < criteriaB
		}
	},
	SortModeGemQuality: func(items []SortableItem) func(i, j int) bool {
		return func(i, j int) bool {
			criteriaA := GemSortMap[items[i].AdditionalProps.ItemSubCategory]
			criteriaB := GemSortMap[items[j].AdditionalProps.ItemSubCategory]
			return criteriaA < criteriaB
		}
	},
	SortModeItemLevel: func(items []SortableItem) func(i, j int) bool {
		return func(i, j int) bool {

			sortMap := [][]string{
				{string(rune(items[i].D2sItem.Identified)), string(rune(items[j].D2sItem.Identified)), string(rune(0))},
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
				{string(rune(items[i].D2sItem.Quantity)), string(rune(items[j].D2sItem.Quantity)), string(rune(0))},
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
			return items[i].D2sItem.Type < items[j].D2sItem.Type
		}
	},
}

type ItemGroup struct {
	Name          string         `yaml:"name,omitempty"`
	IndexType     StashPageType  `yaml:"indexType,omitempty"`
	GroupFilter   []GroupFilter  `yaml:"groupFilter,omitempty"`
	MatchingItems []SortableItem `yaml:"-"`
	StashPages    []StashPage    `yaml:"-"`
	SortMode      SortMode       `yaml:"sortMode,omitempty"`
}

func (g *ItemGroup) addItem(item SortableItem) {
	//fmt.Printf("adding item to group \n")
	g.MatchingItems = append(g.MatchingItems, item)
	//fmt.Printf("%d items in group \n", len(g.MatchingItems))
}

func (g *ItemGroup) clear() {
	g.MatchingItems = []SortableItem{}
}

func (g *ItemGroup) createStashPages() []StashPage {
	return createStashPagesForItems(g.MatchingItems, g.IndexType, g.Name, g.SortMode)
}

func loadGroupConfigFile(filepath string) []ItemGroup {
	groups := []ItemGroup{}
	//test := Test{A: 1, B: Int(1), c: Int(1)}
	configFileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic("can not open config file: " + filepath)
	}

	err = yaml.Unmarshal(configFileData, &groups)

	if err != nil {
		panic("Error parsing group config file: " + err.Error())
	}

	return groups
}

func createStashPagesForItems(items []SortableItem, indexType StashPageType, name string, sortMode SortMode) []StashPage {
	newStashPages := []StashPage{}
	currentStashPage := StashPage{Nr: 0, Name: name, Items: []SortableItem{}, Type: indexType, spaces: [10][10]int{}}

	if sortModeFunctionMap[sortMode] != nil {
		sort.SliceStable(items, sortModeFunctionMap[sortMode](items))
	} else {
		sort.SliceStable(items, sortModeFunctionMap[SortModeItemLevel](items))
	}

	/*if sortMode == SortModeGemQuality {
		sort.SliceStable(items, )
	} else if sortMode == SortModeRuneNumber {
		sort.SliceStable(items, )
	} else if sortMode == SortModeName {
		sort.SliceStable(items, )
	} else {
		sort.SliceStable(items,)
	}*/

	for _, item := range items {

		isInserted := currentStashPage.insertItem(item)
		//fmt.Printf("inserted item: %s on page: %v at %d %d\n", item.D2sItem.Type, isInserted, item.D2sItem.PositionX, item.D2sItem.PositionY)

		if !isInserted { //means the page is full

			//check previous pages
			newStashPages = append(newStashPages, currentStashPage)
			currentStashPage = StashPage{Nr: 0, Name: name, Items: []SortableItem{}, Type: NormalPage, spaces: [10][10]int{}}
			currentStashPage.insertItem(item)
		}

	}
	if len(currentStashPage.Items) != 0 {
		newStashPages = append(newStashPages, currentStashPage)
	}
	return newStashPages
}

func filterMatches(filter GroupFilter, item SortableItem) bool {
	if filter.UniqueItems != nil && *filter.UniqueItems != item.IsUnique {
		return false
	}

	if filter.SetItems != nil && *filter.SetItems != item.IsSetItem {
		return false
	}

	if filter.RareItems != nil && *filter.RareItems != item.isRare {
		return false
	}

	if filter.Identified != nil && *filter.Identified != Uint64ToBool(item.D2sItem.Identified) {
		return false
	}

	if filter.ExcludedCategories != nil && len(*filter.ExcludedCategories) > 0 {
		for _, filterCategory := range *filter.ExcludedCategories {
			if filterCategory == string(item.AdditionalProps.ItemCategory) ||
				filterCategory == string(item.AdditionalProps.ItemSubCategory) ||
				filterCategory == string(item.AdditionalProps.ItemSubSubCategory) {
				return false
			}
		}
	}

	if filter.IncludedCategories != nil && len(*filter.IncludedCategories) > 0 {
		for _, filterCategory := range *filter.IncludedCategories {
			if filterCategory == string(item.AdditionalProps.ItemCategory) ||
				filterCategory == string(item.AdditionalProps.ItemSubCategory) ||
				filterCategory == string(item.AdditionalProps.ItemSubSubCategory) {
			} else {
				return false
			}
		}
	}

	if filter.MinFilledSockets != nil && *filter.MinFilledSockets > 0 {
		if int(item.D2sItem.NrOfItemsInSockets) < *filter.MinFilledSockets {
			return false
		}
	}

	if filter.MaxFilledSockets != nil && *filter.MaxFilledSockets > 0 {
		if int(item.D2sItem.NrOfItemsInSockets) > *filter.MaxFilledSockets {
			return false
		}
	}

	if filter.MaxTotalSockets != nil && *filter.MaxTotalSockets > 0 {
		if int(item.D2sItem.TotalNrOfSockets) > *filter.MaxTotalSockets {
			return false
		}
	}

	if filter.MinTotalSockets != nil && *filter.MinTotalSockets > 0 {
		if int(item.D2sItem.TotalNrOfSockets) < *filter.MinTotalSockets {
			return false
		}
	}

	return true
}

func belongsToGroup(group ItemGroup, item SortableItem) bool {
	for _, filter := range group.GroupFilter {
		if filterMatches(filter, item) {
			return true
		}
	}
	return false
}

func groupItems(groups []ItemGroup, items []SortableItem) []StashPage {

	remainingItems := []SortableItem{}
	_ = remainingItems

	for idx := range groups {
		groups[idx].clear()
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
				groups[idx].addItem(item)
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
	allPages := []StashPage{}
	for _, group := range groups {
		//fmt.Printf("\tcreating stash pages for %s and %d items sorted by %s\n", group.Name, len(group.MatchingItems), group.SortMode)
		allPages = append(allPages, group.createStashPages()...)
	}

	//create pages for all remaining items
	//fmt.Printf("creating stash pages for %d remaining items\n", len(remainingItems))
	allPages = append(allPages, createStashPagesForItems(remainingItems, MainIndexPage, "Misc", SortModeItemLevel)...)

	//return all stash pages created
	return allPages
}
