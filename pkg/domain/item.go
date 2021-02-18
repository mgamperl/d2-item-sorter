package domain

import (
	"d2-item-sorter/pkg/data"

	"github.com/nokka/d2s"
)

//Type represents the Type
type Type struct {
	ID      string               `json:"id"`
	Name    string               `json:"name"`
	Quality data.ItemTypeQuality `json:"quality"`
	SizeX   int                  `json:"sizeX"`
	SizeY   int                  `json:"sizeY"`
}

//ItemQuality representing the item quality constants
type ItemQuality string

const (
	//ItemQualityNormal are normal items (white)
	ItemQualityNormal ItemQuality = "Normal"
	//ItemQualityMagic are magic items (blue)
	ItemQualityMagic ItemQuality = "Magic"
	//ItemQualityRare are rare items (yellow)
	ItemQualityRare ItemQuality = "Rare"
	//ItemQualitySet are rare items (green)
	ItemQualitySet ItemQuality = "Set"
	//ItemQualityUnique are unique items (golden)
	ItemQualityUnique   ItemQuality = "Unique"
	ItemQualityRuneword ItemQuality = "Runeword"
)

//Item represents an ingame item, used for API output and also for internal calcuations
type Item struct {
	ID             string                      `json:"id"`       //id to identify an item (ie. which kind of unique, which kind of rare)
	Name           string                      `json:"name"`     //display name for the item
	Quality        ItemQuality                 `json:"quality"`  //Normal, Magic, Rare, Unique, Set, ..
	Type           Type                        `json:"type"`     //Which kind of item (Cap, Greaves, Hand Axe)
	StoredAt       ItemLocation                `json:"location"` //where is the item (inventory, stash, equipped, corpse, merc, ...)
	ImageID        string                      `json:"imageId"`  //which id to use for showing the right image
	Tags           []string                    `json:"tags"`     //tags for item categorization
	Source         string                      `json:"source"`   //source of the item (character or shared stash)
	Identified     bool                        `json:"identified"`
	Ethereal       bool                        `json:"etheral"`
	Sockets        int                         `json:"sockets"`
	SocketedWith   []Item                      `json:"socketedWith"`
	origItem       d2s.Item                    `json:"-"`
	Category       data.ItemCategoryName       `json:"category"`
	SubCategory    data.ItemSubCategoryName    `json:"subCategory"`
	SubSubCategory data.ItemSubSubCategoryName `json:"subSubCategory"`
}

type ItemLocation struct {
	X    int                    `json:"x"`
	Y    int                    `json:"y"`
	Name data.EquipmentLocation `json:"location"`
}

func (i *Item) SetPosition(x int, y int) {
	i.StoredAt.X = x
	i.StoredAt.Y = y
}

func (i *Item) GetOriginalData() []byte {
	return i.origItem.OriginalData
}

func (i *Item) SetOriginalItem(item d2s.Item) {
	i.origItem = item
}

func ItemFilterFunc(ss []Item, test func(Item) bool) (ret []Item) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
