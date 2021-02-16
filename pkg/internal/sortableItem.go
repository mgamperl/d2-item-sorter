package internal

import (
	"d2-item-sorter/pkg/data"
	"d2-item-sorter/pkg/domain"
	"d2-item-sorter/pkg/internal/utils"
	"fmt"
	"strconv"

	"github.com/nokka/d2s"
)

type SortableItem struct {
	ID              string                   `json:"id"`
	D2sItem         d2s.Item                 `json:"details"`
	AdditionalProps data.ItemPropertyData    `json:"additionalData"`
	IsUnique        bool                     `json:"isUnique"`
	IsSetItem       bool                     `json:"isSetItem"`
	isRare          bool                     `json:"isRare"`
	isMagic         bool                     `json:"isMagic"`
	isNormal        bool                     `json:"isNormal"`
	isRuneword      bool                     `json:"isRuneword"`
	Location        domain.EquipmentLocation `json:"location"`
	ImageId         string                   `json:"imageId"`
	Name            string                   `json:"name"`
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

func (i *SortableItem) SetItem(item d2s.Item, stashType domain.StashType, mercenary bool) {
	i.D2sItem = item

	i.ImageId = fmt.Sprintf("%s", i.D2sItem.Type)

	if len(item.UniqueName) > 0 || item.UniqueID > 0 {
		i.IsUnique = true
		i.ImageId = fmt.Sprintf("u%d", i.D2sItem.UniqueID)
	}

	if len(item.RareName) > 0 || len(item.RareName2) > 0 {
		i.isRare = true
	}

	if len(item.SetName) > 0 {
		i.IsSetItem = true
		i.ImageId = fmt.Sprintf("s%d", i.D2sItem.SetID)
	}

	if len(item.RunewordName) > 0 {
		i.isRuneword = true
	}

	i.Location = data.EquipmentPositionMap[byte(i.D2sItem.LocationID)]
	if i.Location == domain.LocationInventory {
		if stashType == domain.StashTypeShared {
			i.Location = domain.LocationSharedStash
		} else {
			i.Location = domain.LocationPersonalStash
			value, exists := data.PanelPositionMap[byte(i.D2sItem.AltPositionID)]
			if exists {
				i.Location = value
			}

		}

	}

	if mercenary {
		i.Location = domain.LocationMercenary
	}

	if i.D2sItem.MultiplePictures == 1 {
		// If it's annihilus or hellfire torch, we'll display their images instead.
		if i.IsUnique && (i.D2sItem.Type == "cm1" || i.D2sItem.Type == "cm2") {
			i.ImageId = fmt.Sprintf("u%d", i.D2sItem.UniqueID)
		} else {
			i.ImageId = fmt.Sprintf("%s_%d", i.D2sItem.Type, i.D2sItem.PictureID)
		}
	}

	i.Name = i.GetName()

	if i.IsUnique {
		i.ID = fmt.Sprintf("u%d", i.D2sItem.UniqueID)
	} else if i.IsSetItem {
		i.ID = fmt.Sprintf("s%d", i.D2sItem.SetID)
	} else if i.AdditionalProps.ItemCategory == data.Runes {
		i.ID = i.D2sItem.Type
	} else {
		i.ID = fmt.Sprintf("%d", utils.Hash(string(i.D2sItem.OriginalData)))
	}

}

func (i *SortableItem) SetItemFromSharedStash(item d2s.Item) {
	i.SetItem(item, domain.StashTypeShared, false)
}

func (i *SortableItem) SetItemFromCharacter(item d2s.Item) {
	i.SetItem(item, domain.StashTypeCharacter, false)
}

func (i *SortableItem) SetItemFromMercenary(item d2s.Item) {
	i.SetItem(item, domain.StashTypeCharacter, true)
}

func (i *SortableItem) setPosition(x int, y int) {
	i.D2sItem.PositionX = uint64(x)
	i.D2sItem.PositionY = uint64(y)
	//fmt.Printf("set position X to %d\n", i.D2sItem.PositionX)
	i.D2sItem.OriginalData = utils.WriteBits(i.D2sItem.OriginalData, 65, 4, int(i.D2sItem.PositionX))
	//fmt.Printf("set position Y to %d\n", i.D2sItem.PositionY)
	i.D2sItem.OriginalData = utils.WriteBits(i.D2sItem.OriginalData, 69, 4, int(i.D2sItem.PositionY))
}
