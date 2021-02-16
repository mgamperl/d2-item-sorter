package domain

import (
	"github.com/nokka/d2s"
)

type EquipmentLocation string

const (
	LocationInventory     EquipmentLocation = "inventory"
	LocationEquipped      EquipmentLocation = "equipped"
	LocationBelt          EquipmentLocation = "belt"
	LocationSocketed      EquipmentLocation = "socketed"
	LocationCursor        EquipmentLocation = "cursor"
	LocationSharedStash   EquipmentLocation = "sharedStash"
	LocationPersonalStash EquipmentLocation = "personalStash"
	LocationCube          EquipmentLocation = "cube"
	LocationMercenary     EquipmentLocation = "mercenary"
)

//ItemType represents the ItemType
type ItemType struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Quality string `json:"quality"`
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
	ID       string            `json:"id"`       //id to identify an item (ie. which kind of unique, which kind of rare)
	Name     string            `json:"name"`     //display name for the item
	Quality  ItemQuality       `json:"quality"`  //Normal, Magic, Rare, Unique, Set, ..
	Type     ItemType          `json:"type"`     //Which kind of item (Cap, Greaves, Hand Axe)
	Location EquipmentLocation `json:"location"` //where is the item (inventory, stash, equipped, corpse, merc, ...)
	ImageID  string            `json:"imageId"`  //which id to use for showing the right image
	Tags     []string          `json:"tags"`     //tags for item categorization
	Source   string            `json:"source"`   //source of the item (character or shared stash)
	origItem d2s.Item          `json:"-"`
}
