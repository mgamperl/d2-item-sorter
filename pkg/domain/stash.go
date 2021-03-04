package domain

import (
	"bufio"
	"d2-item-sorter/pkg/data"
)

type SharedGold struct {
	Amount int    `json:"amount"`
	Bytes  []byte `json:"-"`
}

type Stash struct {
	StashType  data.StashType `json:"type"`
	SharedGold SharedGold     `json:"sharedGold"`
	StashPages []StashPage    `json:"pages"`
	ItemCount  int            `json:"itemCount"`
	Filepath   string         `json:"filename"`
	ExtraByte  byte           `json:"-"`
	Bfr        *bufio.Reader  `json:"-"`
	Version    string         `json:"-"`
}

func (s *Stash) GetAllItems() []Item {
	allItems := []Item{}
	for _, stashPage := range s.StashPages {
		allItems = append(allItems, stashPage.Items...)
	}
	return allItems
}
