package sorter

import (
	"strconv"

	"d2-item-sorter/utils"

	"github.com/nokka/d2s"
)

type SortableItem struct {
	D2sItem         d2s.Item          `json:"details"`
	AdditionalProps ItemPropertyData  `json:"additionalData"`
	IsUnique        bool              `json:"isUnique"`
	IsSetItem       bool              `json:"isSetItem"`
	isRare          bool              `json:"isRare"`
	isMagic         bool              `json:"isMagic"`
	isNormal        bool              `json:"isNormal"`
	isRuneword      bool              `json:"isRuneword"`
	Location        EquipmentLocation `json:"location"`
}

func (i *SortableItem) GetName() string {

	eth := ""
	sockets := ""
	prefixes := ""
	suffixes := ""

	if len(i.D2sItem.MagicPrefixName) > 0 {
		prefixes = i.D2sItem.MagicPrefixName + i.D2sItem.RareName + " "
	}

	if len(i.D2sItem.RareName2+i.D2sItem.MagicSuffixName) > 0 {
		suffixes = " of " + i.D2sItem.RareName2 + i.D2sItem.MagicSuffixName
	}

	if len(i.D2sItem.RareName) > 0 {
		prefixes = i.D2sItem.RareName
	}

	if len(i.D2sItem.RareName2) > 0 {
		suffixes = " " + i.D2sItem.RareName2
	}

	if i.D2sItem.Ethereal == 1 {
		eth = "[ETH] "
	}

	if i.D2sItem.TotalNrOfSockets > 0 {
		sockets = " (" + strconv.Itoa(int(i.D2sItem.TotalNrOfSockets)) + ")"
	}

	if i.IsUnique {
		return eth + i.D2sItem.UniqueName + " / " + i.D2sItem.TypeName
	}

	if i.IsSetItem {
		return eth + i.D2sItem.SetName
	}

	if i.isRare || i.isMagic {
		return eth + prefixes + suffixes + " / " + i.D2sItem.TypeName
	}

	if i.isRuneword {
		return eth + i.D2sItem.RunewordName + " / " + i.D2sItem.TypeName
	}

	return eth + prefixes + i.D2sItem.TypeName + suffixes + sockets
}

func (i *SortableItem) SetItem(item d2s.Item, stashType StashType, mercenary bool) {
	i.D2sItem = item

	if len(item.UniqueName) > 0 {
		i.IsUnique = true
	}

	if len(item.RareName) > 0 || len(item.RareName2) > 0 {
		i.isRare = true
	}

	if len(item.SetName) > 0 {
		i.IsSetItem = true
	}

	if len(item.RunewordName) > 0 {
		i.isRuneword = true
	}

	i.Location = EquipmentPositionMap[byte(i.D2sItem.LocationID)]
	if i.Location == LocationInventory {
		if stashType == StashTypeShared {
			i.Location = LocationSharedStash
		} else {
			i.Location = LocationPersonalStash
			value, exists := PanelPositionMap[byte(i.D2sItem.AltPositionID)]
			if exists {
				i.Location = value
			}

		}

	}

	if mercenary {
		i.Location = LocationMercenary
	}

}

func (i *SortableItem) SetItemFromSharedStash(item d2s.Item) {
	i.SetItem(item, StashTypeShared, false)
}

func (i *SortableItem) SetItemFromCharacter(item d2s.Item) {
	i.SetItem(item, StashTypeCharacter, false)
}

func (i *SortableItem) SetItemFromMercenary(item d2s.Item) {
	i.SetItem(item, StashTypeCharacter, true)
}

func (i *SortableItem) setPosition(x int, y int) {
	i.D2sItem.PositionX = uint64(x)
	i.D2sItem.PositionY = uint64(y)
	//fmt.Printf("set position X to %d\n", i.D2sItem.PositionX)
	i.D2sItem.OriginalData = utils.WriteBits(i.D2sItem.OriginalData, 65, 4, int(i.D2sItem.PositionX))
	//fmt.Printf("set position Y to %d\n", i.D2sItem.PositionY)
	i.D2sItem.OriginalData = utils.WriteBits(i.D2sItem.OriginalData, 69, 4, int(i.D2sItem.PositionY))
}

type StashPageType string

const (
	MainIndexPage StashPageType = "MainIndex"
	IndexPage     StashPageType = "Index"
	NormalPage    StashPageType = "Normal"
)

var StashPageByteMap = map[byte]StashPageType{
	0x07: MainIndexPage,
	0x06: MainIndexPage,
	0x03: IndexPage,
	0x02: IndexPage,
	0x01: NormalPage,
	0x00: NormalPage,
}

var StashPageIndexMap = map[StashType]map[StashPageType]byte{
	StashTypeShared: {
		MainIndexPage: 0x07,
		IndexPage:     0x03,
		NormalPage:    0x01,
	},
	StashTypeCharacter: {
		MainIndexPage: 0x06,
		IndexPage:     0x02,
		NormalPage:    0x00,
	},
}

type StashPage struct {
	Nr                 uint16         `json:"number"`
	Name               string         `json:"name"`
	Items              []SortableItem `json:"items"`
	Type               StashPageType  `json:"type"`
	spaces             [10][10]int
	isIgnoredPage      bool
	isPickupPage       bool
	isSharedPickupPage bool
}

func (p *StashPage) isCollision(x_position int, x_size int, y_position int, y_size int) bool {
	for x := 0; x < x_size; x++ {
		for y := 0; y < y_size; y++ {
			if x_position+x > 9 || y_position+y > 9 || p.spaces[x_position+x][y_position+y] > 0 {
				return true
			}
		}
	}
	return false
}

func (p *StashPage) allocateItem(x_position int, x_size int, y_position int, y_size int) {
	for x := 0; x < x_size; x++ {
		for y := 0; y < y_size; y++ {
			p.spaces[x_position+x][y_position+y] += 1
		}
	}
}

func (p *StashPage) insertItem(item SortableItem) bool {
	allocated := false
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if !allocated && !p.isCollision(x, item.AdditionalProps.SizeX, y, item.AdditionalProps.SizeY) {
				item.setPosition(x, y)
				p.Items = append(p.Items, item)
				p.allocateItem(x, item.AdditionalProps.SizeX, y, item.AdditionalProps.SizeY)
				allocated = true
			}
		}
	}
	return allocated
}

func (p *StashPage) insertItemAtColumn(item SortableItem, column int) bool {
	allocated := false
	x := column
	for y := 0; y < 10; y++ {
		if !allocated && !p.isCollision(x, item.AdditionalProps.SizeX, y, item.AdditionalProps.SizeY) {
			item.setPosition(x, y)
			p.Items = append(p.Items, item)
			p.allocateItem(x, item.AdditionalProps.SizeX, y, item.AdditionalProps.SizeY)
			allocated = true
		}
	}
	return allocated
}
