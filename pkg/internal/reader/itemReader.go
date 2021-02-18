package reader

import (
	"d2-item-sorter/pkg/data"
	"d2-item-sorter/pkg/domain"
	"d2-item-sorter/pkg/internal/utils"
	"fmt"

	"github.com/nokka/d2s"
)

func NewItem(item d2s.Item, stashType data.StashType) *domain.Item {
	//i.D2sItem = item
	i := domain.Item{}
	itemType := data.GetItemProperties(item)

	i.ID = fmt.Sprintf("%d", utils.Hash(string(item.OriginalData)))
	i.ImageID = fmt.Sprintf("%s", item.Type)
	i.Type = domain.Type{ID: item.Type, Name: item.TypeName, Quality: data.GetItemTypeQuality(item), SizeX: itemType.SizeX, SizeY: itemType.SizeY}
	i.Quality = domain.ItemQualityNormal
	i.SetOriginalItem(item)
	i.Identified = utils.Uint64ToBool(item.Identified)
	i.Ethereal = utils.Uint64ToBool(item.Ethereal)
	i.Sockets = int(item.TotalNrOfSockets)
	i.SocketedWith = NewItemsFromSockets(item.SocketedItems)
	i.Category = itemType.ItemCategory
	i.SubCategory = itemType.ItemSubCategory
	i.SubSubCategory = itemType.ItemSubSubCategory

	switch {
	case data.IsUnique(item):
		i.ID = fmt.Sprintf("u%d", item.UniqueID)
		i.Quality = domain.ItemQualityUnique
		i.ImageID = fmt.Sprintf("u%d", item.UniqueID)
	case data.IsSet(item):
		i.ID = fmt.Sprintf("s%d", item.SetID)
		i.Quality = domain.ItemQualitySet
		i.ImageID = fmt.Sprintf("s%d", item.SetID)
	case data.IsRare(item):
		i.Quality = domain.ItemQualityRare
	case data.IsMagic(item):
		i.Quality = domain.ItemQualityMagic
	case data.IsRuneword(item):
		i.Quality = domain.ItemQualityRuneword
	case data.IsRune(item):
		i.ID = item.Type
	}

	i.StoredAt = domain.ItemLocation{X: int(item.PositionX), Y: int(item.PositionY), Name: data.EquipmentPositionMap[byte(item.LocationID)]}

	if i.StoredAt.Name == data.LocationInventory {
		if stashType == data.StashTypeShared {
			i.StoredAt.Name = data.LocationSharedStash
		} else {
			i.StoredAt.Name = data.LocationPersonalStash
			value, exists := data.PanelPositionMap[byte(item.AltPositionID)]
			if exists {
				i.StoredAt.Name = value
			}

		}
	}

	if item.MultiplePictures == 1 {
		// If it's annihilus or hellfire torch, we'll display their images instead.
		if i.Quality == domain.ItemQualityUnique && (i.Type.ID == "cm1" || i.Type.ID == "cm2" || i.Type.ID == "jew") {
			i.ImageID = fmt.Sprintf("u%d", item.UniqueID)
		} else {
			i.ImageID = fmt.Sprintf("%s_%d", item.Type, item.PictureID)
		}
	}

	i.Name = data.GetItemName(item)

	i.Hash = fmt.Sprintf("%d", utils.Hash(fmt.Sprintf("%s", i.Name)))

	return &i
}

func NewItemsFromSockets(items []d2s.Item) []domain.Item {
	result := []domain.Item{}
	for _, item := range items {
		newItem := NewItem(item, data.StashTypeCharacter)
		newItem.StoredAt.Name = data.LocationSocketed
		result = append(result, *newItem)
	}
	return result
}

func NewItemFromSharedStash(item d2s.Item) *domain.Item {
	return NewItem(item, data.StashTypeShared)
}

func NewItemFromCharacter(item d2s.Item) *domain.Item {
	return NewItem(item, data.StashTypeCharacter)
}

func NewItemFromMercenary(item d2s.Item) *domain.Item {
	newItem := NewItem(item, data.StashTypeCharacter)
	newItem.StoredAt.Name = data.LocationMercenary
	return newItem
}

func NewItemsFromMercenary(items []d2s.Item) []domain.Item {
	result := []domain.Item{}
	for _, item := range items {
		result = append(result, *NewItemFromMercenary(item))
	}
	return result
}

func NewItemsFromCharacter(items []d2s.Item) []domain.Item {
	result := []domain.Item{}
	for _, item := range items {
		result = append(result, *NewItemFromCharacter(item))
	}
	return result
}
