package data

import "d2-item-sorter/pkg/domain"

var EquipmentPositionMap = map[byte]domain.EquipmentLocation{
	0x00: domain.LocationInventory,
	0x01: domain.LocationEquipped,
	0x02: domain.LocationBelt,
	0x06: domain.LocationSocketed,
	0x04: domain.LocationCursor,
}

var PanelPositionMap = map[byte]domain.EquipmentLocation{
	0x01: domain.LocationInventory,
	0x05: domain.LocationPersonalStash,
}

var StashPageIndexMap = map[domain.StashType]map[domain.StashPageType]byte{
	domain.StashTypeShared: {
		domain.MainIndexPage: 0x07,
		domain.IndexPage:     0x03,
		domain.NormalPage:    0x01,
	},
	domain.StashTypeCharacter: {
		domain.MainIndexPage: 0x06,
		domain.IndexPage:     0x02,
		domain.NormalPage:    0x00,
	},
}

var StashPageByteMap = map[byte]domain.StashPageType{
	0x07: domain.MainIndexPage,
	0x06: domain.MainIndexPage,
	0x03: domain.IndexPage,
	0x02: domain.IndexPage,
	0x01: domain.NormalPage,
	0x00: domain.NormalPage,
}
