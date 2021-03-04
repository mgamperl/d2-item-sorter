package domain

import (
	"d2-item-sorter/pkg/data"

	"github.com/nokka/d2s"
)

type Character struct {
	Name                  string `json:"name"`
	Level                 int    `json:"level"`
	Class                 string `json:"class"`
	Items                 []Item `json:"items"`
	SaveFilename          string `json:"filename"`
	PersonalStashFilename string `json:"stashFilename"`
	D2Char                d2s.Character
	PersonalStash         *Stash `json:"personalStash"`
}

func (c *Character) GetInventoryItems() []Item {
	return ItemFilterFunc(c.Items, func(i Item) bool {
		return i.StoredAt.Name == data.LocationInventory
	})
}

func (c *Character) GetMercItems() []Item {
	return ItemFilterFunc(c.Items, func(i Item) bool {
		return i.StoredAt.Name == data.LocationMercenary
	})
}

func (c *Character) GetEquippedItems() []Item {
	return ItemFilterFunc(c.Items, func(i Item) bool {
		return i.StoredAt.Name == data.LocationEquipped
	})
}

func (c *Character) GetBeltItems() []Item {
	return ItemFilterFunc(c.Items, func(i Item) bool {
		return i.StoredAt.Name == data.LocationBelt
	})
}
