package data

var StashPageByteMap = map[byte]StashPageType{
	0x07: MainIndexPage,
	0x06: MainIndexPage,
	0x03: IndexPage,
	0x02: IndexPage,
	0x01: NormalPage,
	0x00: NormalPage,
}

type StashPageType string

const (
	MainIndexPage StashPageType = "MainIndex"
	IndexPage     StashPageType = "Index"
	NormalPage    StashPageType = "Normal"
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

var EquipmentPositionMap = map[byte]EquipmentLocation{
	0x00: LocationInventory,
	0x01: LocationEquipped,
	0x02: LocationBelt,
	0x06: LocationSocketed,
	0x04: LocationCursor,
}

var PanelPositionMap = map[byte]EquipmentLocation{
	0x01: LocationInventory,
	0x05: LocationPersonalStash,
}

type StashType string

const (
	StashTypeShared    StashType = "SSS"
	StashTypeCharacter StashType = "CST"
	StashTypeUnknown   StashType = ""
)

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
