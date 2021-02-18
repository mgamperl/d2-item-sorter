package domain

import (
	"d2-item-sorter/pkg/data"
)

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

type ItemGroup struct {
	Name          string             `yaml:"name,omitempty"`
	IndexType     data.StashPageType `yaml:"indexType,omitempty"`
	GroupFilter   []GroupFilter      `yaml:"groupFilter,omitempty"`
	MatchingItems []Item             `yaml:"-"`
	StashPages    []StashPage        `yaml:"-"`
	SortMode      data.SortMode      `yaml:"sortMode,omitempty"`
}

func (g *ItemGroup) AddItem(item Item) {
	//fmt.Printf("adding item to group \n")
	g.MatchingItems = append(g.MatchingItems, item)
	//fmt.Printf("%d items in group \n", len(g.MatchingItems))
}

func (g *ItemGroup) Clear() {
	g.MatchingItems = []Item{}
}
