package sorter

import "strconv"

type ItemCategoryName string

const (
	Armor    ItemCategoryName = "Armor"
	Weapons  ItemCategoryName = "Weapons"
	Jewelery ItemCategoryName = "Jewelery"
	Gems     ItemCategoryName = "Gems"
	Charms   ItemCategoryName = "Charms"
	Runes    ItemCategoryName = "Runes"
	Potions  ItemCategoryName = "Potions"
	Misc     ItemCategoryName = "Misc"
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

type ItemSubCategoryName string

const (
	Helmet              ItemSubCategoryName = "Helmet"
	BodyArmor           ItemSubCategoryName = "BodyArmor"
	Shield              ItemSubCategoryName = "Shield"
	Gloves              ItemSubCategoryName = "Gloves"
	Boots               ItemSubCategoryName = "Boots"
	Belt                ItemSubCategoryName = "Belt"
	Head                ItemSubCategoryName = "Head"
	SwordsAndDaggers    ItemSubCategoryName = "SwordsAndDaggers"
	Axes                ItemSubCategoryName = "Axes"
	HammersAndMaces     ItemSubCategoryName = "Maces"
	PolearmsAndSpears   ItemSubCategoryName = "PolearmsAndSpears"
	StavesAndWands      ItemSubCategoryName = "StaffAndWands"
	Scepters            ItemSubCategoryName = "Scepters"
	BowAndCrossbow      ItemSubCategoryName = "BowAndCrossbow"
	QuestItems          ItemSubCategoryName = "QuestItems"
	Perfect             ItemSubCategoryName = "Perfect"
	Flawless            ItemSubCategoryName = "Flawless"
	Normal              ItemSubCategoryName = "Normal"
	Flawed              ItemSubCategoryName = "Flawed"
	Chipped             ItemSubCategoryName = "Chipped"
	Claws               ItemSubCategoryName = "Claws"
	HealingPotions      ItemSubCategoryName = "HealingPotions"
	ManaPotions         ItemSubCategoryName = "ManaPotions"
	RejuvenationPotions ItemSubCategoryName = "RejuvenationPotions"
	MiscPotions         ItemSubCategoryName = "MiscPotions"
	GrandCharms         ItemSubCategoryName = "GrandCharms"
	LargeCharms         ItemSubCategoryName = "LargeCharms"
	SmallCharms         ItemSubCategoryName = "SmallCharms"
	Runes1to10          ItemSubCategoryName = "Runes1to10"
	Runes11to20         ItemSubCategoryName = "Runes11to20"
	Runes20to30         ItemSubCategoryName = "Runes20to30"
	Runes31to33         ItemSubCategoryName = "Runes31to33"
	Throwables          ItemSubCategoryName = "Throwables"
	Javelins            ItemSubCategoryName = "Javelins"
	RingsAndAmulets     ItemSubCategoryName = "RingsAndAmulets"
	Jewels              ItemSubCategoryName = "Jewels"
	ScrollsAndTomes     ItemSubCategoryName = "ScrollsAndTomes"
	Consumables         ItemSubCategoryName = "Consumables"
)

var GemSortMap = map[ItemSubCategoryName]int{
	Perfect:  1,
	Flawless: 2,
	Normal:   3,
	Flawed:   4,
	Chipped:  5,
}

type ItemSubSubCategoryName string

const (
	NormalHelmet     ItemSubSubCategoryName = "NormalHelmet"
	NormalBodyArmor  ItemSubSubCategoryName = "NormalBodyArmor"
	NormalShield     ItemSubSubCategoryName = "NormalShield"
	NormalGloves     ItemSubSubCategoryName = "NormalGloves"
	NormalBoots      ItemSubSubCategoryName = "NormalBoots"
	NormalBelts      ItemSubSubCategoryName = "NormalBelts"
	Pelts            ItemSubSubCategoryName = "Pelts"
	BarbarianHelmets ItemSubSubCategoryName = "BarbarianHelmets"
	PaladinShields   ItemSubSubCategoryName = "PaladinShields"
	NecroHeads       ItemSubSubCategoryName = "NecroHeads"
	Axe              ItemSubSubCategoryName = "AXE"
	Maces            ItemSubSubCategoryName = "Maces"
	Swords           ItemSubSubCategoryName = "Swords"
	Daggers          ItemSubSubCategoryName = "Daggers"
	Throwable        ItemSubSubCategoryName = "THROW"
	Javelin          ItemSubSubCategoryName = "Javelin"
	Spears           ItemSubSubCategoryName = "Spears"
	Polearms         ItemSubSubCategoryName = "Polearms"
	Bows             ItemSubSubCategoryName = "Bows"
	Crossbows        ItemSubSubCategoryName = "Crossbows"
	Staves           ItemSubSubCategoryName = "Staves"
	Wands            ItemSubSubCategoryName = "Wands"
	SCEPTER          ItemSubSubCategoryName = "SCEPTER"
	AssassinWeapons  ItemSubSubCategoryName = "AssassinWeapons"
	SorcWeapons      ItemSubSubCategoryName = "SorcWeapons"
	AmazonWeapons    ItemSubSubCategoryName = "AmazonWeapons"
	Circlets         ItemSubSubCategoryName = "Circlets"
	ThrowablePots    ItemSubSubCategoryName = "ThrowablePots"
	QuestItem        ItemSubSubCategoryName = "QuestItem"
	Rune             ItemSubSubCategoryName = "Rune"
	Potion           ItemSubSubCategoryName = "Potion"
	CHARM            ItemSubSubCategoryName = "Charm"
	Scroll           ItemSubSubCategoryName = "Scroll"
	Other            ItemSubSubCategoryName = "Other"
	Jewel            ItemSubSubCategoryName = "Jewel"
	Amulets          ItemSubSubCategoryName = "Amulets"
	Rings            ItemSubSubCategoryName = "Rings"
	UberKeys         ItemSubSubCategoryName = "UberKey"
	UberParts        ItemSubSubCategoryName = "UberPart"
	Essences         ItemSubSubCategoryName = "Essence"
	Amethyst         ItemSubSubCategoryName = "Amethyst"
	Diamond          ItemSubSubCategoryName = "Diamond"
	Emerald          ItemSubSubCategoryName = "Emerald"
	Ruby             ItemSubSubCategoryName = "Ruby"
	Sapphire         ItemSubSubCategoryName = "Sapphire"
	Topaz            ItemSubSubCategoryName = "Topaz"
	Skulls           ItemSubSubCategoryName = "Skulls"
	HEADS            ItemSubSubCategoryName = "HEADS"
	Flails           ItemSubSubCategoryName = "Flails"
	Hammers          ItemSubSubCategoryName = "Hammers"
	ThrowableKnife   ItemSubSubCategoryName = "ThrowableKnife"
	ThrowableAxe     ItemSubSubCategoryName = "ThrowableAxe"
)

type ItemPropertyData struct {
	ItemCategory       ItemCategoryName
	ItemSubCategory    ItemSubCategoryName
	ItemSubSubCategory ItemSubSubCategoryName
	SizeX              int
	SizeY              int
}

var itemPropertyMap = map[string]*ItemPropertyData{

	"cap": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Cap
	"skp": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Skull Cap
	"hlm": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Helm
	"fhl": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Full Helm
	"ghm": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Great Helm
	"crn": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Crown
	"msk": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Mask
	"bhm": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Bone Helm
	"xap": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // War Hat
	"xkp": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Sallet
	"xlm": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Casque
	"xhl": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Basinet
	"xhm": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Winged Helm
	"xrn": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Grand Crown
	"xsk": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Death Mask
	"xh9": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Grim Helm
	"uap": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Shako
	"ukp": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Hydraskull
	"ulm": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Armet
	"uhl": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Giant Conch
	"uhm": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Spired Helm
	"urn": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Corona
	"usk": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Demonhead
	"uh9": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Bone Visage
	"qui": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Quilted Armor
	"lea": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Leather Armor
	"hla": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Hard Leather
	"stu": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Studded Leather
	"rng": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Ring Mail
	"scl": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Scale Mail
	"chn": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Chain Mail
	"brs": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Breast Plate
	"spl": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Splint Mail
	"plt": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Plate Mail
	"fld": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Field Plate
	"gth": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Gothic Plate
	"ful": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Full Plate Mail
	"aar": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Ancient Armor
	"ltp": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Light Plate
	"xui": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Ghost Armor
	"xea": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Serpentskin
	"xla": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Demonhide Armor
	"xtu": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Trellised Armor
	"xng": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Linked Mail
	"xcl": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Tigulated Mail
	"xhn": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Mesh Armor
	"xrs": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Cuirass
	"xpl": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Russet Armor
	"xlt": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Templar Coat
	"xld": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Sharktooth
	"xth": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Embossed Plate
	"xul": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Chaos Armor
	"xar": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Ornate Armor
	"xtp": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Mage Plate
	"uui": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Dusk Shroud
	"uea": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Wyrmhide
	"ula": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Scarab Husk
	"utu": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Wire Fleece
	"ung": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Diamond Mail
	"ucl": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Loricated Mail
	"uhn": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Boneweave
	"urs": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Great Hauberk
	"upl": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Balrog Skin
	"ult": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Hellforge Plate
	"uld": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Kraken Shell
	"uth": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Lacquered Plate
	"uul": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Shadow Plate
	"uar": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Sacred Armor
	"utp": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Archon Plate
	"buc": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Buckler
	"sml": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Small Shield
	"lrg": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Large Shield
	"kit": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Kite Shield
	"tow": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Tower Shield
	"gts": {SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Gothic Shield
	"bsh": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Bone Shield
	"spk": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Spiked Shield
	"xuc": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Defender
	"xml": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Round Shield
	"xrg": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Scutum
	"xit": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Dragon Shield
	"xow": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Pavise
	"xts": {SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Ancient Shield
	"xsh": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Grim Shield
	"xpk": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Barbed Shield
	"uuc": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Heater
	"uml": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Luna
	"urg": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Hyperion
	"uit": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Monarch
	"uow": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Aegis
	"uts": {SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Ward
	"ush": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Troll Nest
	"upk": {SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Blade Barrier
	"lgl": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Leather Gloves
	"vgl": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Heavy Gloves
	"mgl": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Chain Gloves
	"tgl": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Light Gauntlets
	"hgl": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Gauntlets
	"xlg": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Demonhide Glove
	"xvg": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Sharkskin Glove
	"xmg": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Heavy Bracers
	"xtg": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Battle Gauntlet
	"xhg": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // War Gauntlets
	"ulg": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Bramble Mitts
	"uvg": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Vampirebone Gl
	"umg": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Vambraces
	"utg": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Crusader Gaunt
	"uhg": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Ogre Gauntlets
	"lbt": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Boots
	"vbt": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Heavy Boots
	"mbt": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Chain Boots
	"tbt": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Light Plate
	"hbt": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Greaves
	"xlb": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Demonhide Boots
	"xvb": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Sharkskin Boots
	"xmb": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Mesh Boots
	"xtb": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Battle Boots
	"xhb": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // War Boots
	"ulb": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Wyrmhide Boots
	"uvb": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Scarabshell Bts
	"umb": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Boneweave Boots
	"utb": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Mirrored Boots
	"uhb": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Myrmidon Greave
	"lbl": {SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Sash
	"vbl": {SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Light Belt
	"mbl": {SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Belt
	"tbl": {SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Heavy Belt
	"hbl": {SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Plated Belt
	"zlb": {SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Demonhide Sash
	"zvb": {SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Sharkskin Belt
	"zmb": {SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Mesh Belt
	"ztb": {SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Battle Belt
	"zhb": {SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // War Belt
	"ulc": {SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Spiderweb Sash
	"uvc": {SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Vampirefang Blt
	"umc": {SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Mithril Coil
	"utc": {SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Troll Belt
	"uhc": {SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Colossus Girdle
	"dr1": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Wolf Head
	"dr2": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Hawk Helm
	"dr3": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Antlers
	"dr4": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Falcon Mask
	"dr5": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Spirit Mask
	"dr6": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Alpha Helm
	"dr7": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Griffon Headress
	"dr8": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Hunter’s Guise
	"dr9": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Sacred Feathers
	"dra": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Totemic Mask
	"drb": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Blood Spirit
	"drc": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Sun Spirit
	"drd": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Earth Spirit
	"dre": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Sky Spirit
	"drf": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Dream Spirit
	"ba1": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Jawbone Cap
	"ba2": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Fanged Helm
	"ba3": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Horned Helm
	"ba4": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Assualt Helmet
	"ba5": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Avenger Guard
	"ba6": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Jawbone Visor
	"ba7": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Lion Helm
	"ba8": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Rage Mask
	"ba9": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Savage Helmet
	"baa": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Slayer Guard
	"bab": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Carnage Helm
	"bac": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Fury Visor
	"bad": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Destroyer Helm
	"bae": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Conqueror Crown
	"baf": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Guardian Crown
	"pa1": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Targe
	"pa2": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Rondache
	"pa3": {SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Heraldic Shield
	"pa4": {SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Aerin Shield
	"pa5": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Crown Shield
	"pa6": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Akaran Targe
	"pa7": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Akaran Rondache
	"pa8": {SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Protector Shld
	"pa9": {SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Guilded Shield
	"paa": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Royal Shield
	"pab": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Sacred Targe
	"pac": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Sacred Rondache
	"pad": {SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Kurast Shield
	"pae": {SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Zakarum Shield
	"paf": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Vortex Shield
	"ne1": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Preserved Head
	"ne2": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Zombie Head
	"ne3": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Unraveller Head
	"ne4": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Gargoyle Head
	"ne5": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Demon Head
	"ne6": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Mummified Trphy
	"ne7": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Fetish Trophy
	"ne8": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Sexton Trophy
	"ne9": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Cantor Trophy
	"nea": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Heirophant Trphy
	"neb": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Minion Skull
	"neg": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Hellspawn Skull
	"ned": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Overseer Skull
	"nee": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Succubae Skull
	"nef": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Bloodlord Skull
	"hax": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Hand Axe
	"axe": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Axe
	"2ax": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Double Axe
	"mpi": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Military Pick
	"wax": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // War Axe
	"lax": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Large Axe
	"bax": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Broad Axe
	"btx": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Battle Axe
	"gax": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Great Axe
	"gix": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Giant Axe
	"9ha": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Hatchet
	"9ax": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Cleaver
	"92a": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Twin Axe
	"9mp": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Crowbill
	"9wa": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Naga
	"9la": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Military Axe
	"9ba": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Bearded Axe
	"9bt": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Tabar
	"9ga": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Gothic Axe
	"9gi": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Ancient Axe
	"7ha": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Tomahawk
	"7ax": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Small Crescent
	"72a": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Ettin Axe
	"7mp": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // War Spike
	"7wa": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Berserker Axe
	"7la": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Feral Axe
	"7ba": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Silver Edged Ax
	"7bt": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Decapitator
	"7ga": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Champion Axe
	"7gi": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Glorious Axe
	"clb": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Club
	"spc": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Spiked Club
	"mac": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Mace
	"mst": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Morning Star
	"fla": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Flails},          // Flail
	"whm": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Hammers},         // War Hammer
	"mau": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Hammers},         // Maul
	"gma": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Hammers},         // Great Maul
	"9cl": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Cudgel
	"9sp": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Barbed Club
	"9ma": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Flanged Mace
	"9mt": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Jagged Star
	"9fl": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Knout
	"9wh": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Battle Hammer
	"9m9": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // War Club
	"9gm": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Martel de Fer
	"7cl": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Truncheon
	"7sp": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Tyrant Club
	"7ma": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Reinforced Mace
	"7mt": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Devil Star
	"7fl": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Scourge
	"7wh": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Hammers},         // Legendary Mallet
	"7m7": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Hammers},         // Ogre Maul
	"7gm": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Hammers},         // Thunder Maul
	"ssd": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Short Sword
	"scm": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Scimitar
	"sbr": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Saber
	"flc": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Falchion
	"crs": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Crystal Sword
	"bsd": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Broad Sword
	"lsd": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Long Sword
	"wsd": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // War Sword
	"2hs": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Two-handed Swrd
	"clm": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Claymore
	"gis": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Giant Sword
	"bsw": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Bastard Sword
	"flb": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Flamberge
	"gsd": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Great Sword
	"9ss": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Gladius
	"9sm": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Cutlass
	"9sb": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Shamshir
	"9fc": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Tulwar
	"9cr": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Dimensional Bld
	"9bs": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Battle Sword
	"9ls": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Rune Sword
	"9wd": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Ancient Sword
	"92h": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Espadon
	"9cm": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Dacian Falx
	"9gs": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Tusk Sword
	"9b9": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Gothic Sword
	"9fb": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Zweihander
	"9gd": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Executioner Swr
	"7ss": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Falcata
	"7sm": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Ataghan
	"7sb": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Elegant Blade
	"7fc": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Hydra Edge
	"7cr": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Phase Blade
	"7bs": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Conquest Sword
	"7ls": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Cryptic Sword
	"7wd": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Mythical Sword
	"72h": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Legend Sword
	"7cm": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Highland Blade
	"7gs": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Balrog Blade
	"7b7": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Champion Sword
	"7fb": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Colossal Sword
	"7gd": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Colossus Blade
	"dgr": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Dagger
	"dir": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Dirk
	"kri": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Kriss
	"bld": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Blade
	"9dg": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Poignard
	"9di": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Rondel
	"9kr": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Cinquedeas
	"9bl": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Stilleto
	"7dg": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Bone Knife
	"7di": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Mithral Point
	"7kr": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Fanged Knife
	"7bl": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Legend Spike
	"tkf": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableKnife},       // Throwing Knife
	"tax": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableAxe},         // Throwing Axe
	"bkf": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableKnife},       // Balanced Knife
	"bal": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableAxe},         // Balanced Axe
	"9tk": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableKnife},       // Battle Dart
	"9ta": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableAxe},         // Francisca
	"9bk": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableKnife},       // War Dart
	"9b8": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableAxe},         // Hurlbat
	"7tk": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableKnife},       // Flying Knife
	"7ta": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableAxe},         // Flying Axe
	"7bk": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableKnife},       // Winged Knife
	"7b8": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableAxe},         // Winged Axe
	"jav": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Javelin
	"pil": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Pilum
	"ssp": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Short Spear
	"glv": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Glaive
	"tsp": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Throwing Spear
	"9ja": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // War Javelin
	"9pi": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Great Pilum
	"9s9": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Simbilan
	"9gl": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Spiculum
	"9ts": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Harpoon
	"7ja": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Hyperion Javeln
	"7pi": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Stygian Pilum
	"7s7": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Balrog Spear
	"7gl": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Ghost Glaive
	"7ts": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Winged Harpoon
	"spr": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // Spear
	"tri": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // Trident
	"brn": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // Brandistock
	"spt": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // Spetum
	"pik": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // Pike
	"9sr": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // War Spear
	"9tr": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // Fuscina
	"9br": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // War Fork
	"9st": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // Yari
	"9p9": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // Lance
	"7sr": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // Hyperion Spear
	"7tr": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // Stygian Pike
	"7br": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // Mancatcher
	"7st": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // Ghost Spear
	"7p7": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // War Pike
	"bar": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Bardiche
	"vou": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Voulge
	"scy": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Scythe
	"pax": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Poleaxe
	"hal": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Halberd
	"wsc": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // War Scythe
	"9b7": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Lochaber Axe
	"9vo": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Bill
	"9s8": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Battle Scythe
	"9pa": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Partizan
	"9h9": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Bec-de-Corbin
	"9wc": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Grim Scythe
	"7o7": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Ogre Axe
	"7vo": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Colossus Voulge
	"7s8": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Thresher
	"7pa": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Cryptic Axe
	"7h7": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Great Poleaxe
	"7wc": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Giant Thresher
	"sbw": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Short Bow
	"hbw": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Hunter’s Bow
	"lbw": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Long Bow
	"cbw": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Composite Bow
	"sbb": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Shrt Battle Bow
	"lbb": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Long Battle Bow
	"swb": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Short War Bow
	"lwb": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Long War Bow
	"8sb": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Edge Bow
	"8hb": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Razor Bow
	"8lb": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Cedar Bow
	"8cb": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Double Bow
	"8s8": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Short Siege Bow
	"8l8": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Long Siege Bow
	"8sw": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Rune Bow
	"8lw": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Gothic Bow
	"6sb": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Spider Bow
	"6hb": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Blade Bow
	"6lb": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Shadow Bow
	"6cb": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Great Bow
	"6s7": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Diamond Bow
	"6l7": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Crusader Bow
	"6sw": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Ward Bow
	"6lw": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Hydra Bow
	"lxb": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Light Crossbow
	"mxb": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Crossbow
	"hxb": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Heavy Crossbow
	"rxb": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Repeating X-bow
	"8lx": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Arbalest
	"8mx": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Siege Crossbow
	"8hx": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Ballista
	"8rx": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Chu-Ko-Nu
	"6lx": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Pellet Bow
	"6mx": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Gorgon Crossbow
	"6hx": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Colossus x-bow
	"6rx": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Demon Crossbow
	"sst": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Short Staff
	"lst": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Long Staff
	"cst": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Gnarled Staff
	"bst": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Battle Staff
	"wst": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // War Staff
	"8ss": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Jo Staff
	"8ls": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Quarterstaff
	"8cs": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Cedar Staff
	"8bs": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Gothic Staff
	"8ws": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Rune Staff
	"6ss": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Walking Stick
	"6ls": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Stalagmite
	"6cs": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Elder Staff
	"6bs": {SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Shillelagh
	"6ws": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Archon Staff
	"wnd": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Wand
	"ywn": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Yew Wand
	"bwn": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Bone Wand
	"gwn": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Grim Wand
	"9wn": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Burnt Wand
	"9yw": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Petrified Wand
	"9bw": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Tomb Wand
	"9gw": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Grave Wand
	"7wn": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Polished Wand
	"7yw": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Ghost Wand
	"7bw": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Lich Wand
	"7gw": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Unearthed Wand
	"scp": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                // Sceptre
	"gsc": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                // Grand Sceptre
	"wsp": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                // War Sceptre
	"9sc": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                // Rune Sceptre
	"9qs": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                // Holy Water Spri
	"9ws": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                // Divine Sceptre
	"7sc": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                // Mighty Sceptre
	"7qs": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                // Seraph Rod
	"7ws": {SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                // Caduceus
	"ktr": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Katar
	"wrb": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Wrist Blade
	"axf": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Hatchet Hands
	"ces": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Cestus
	"clw": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Claws
	"btl": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Blade Talons
	"skr": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Scissors Katar
	"9ar": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Quhab
	"9wb": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Wrist Spike
	"9xf": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Fascia
	"9cs": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Hand Scythe
	"9lw": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Greater Claws
	"9tw": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Greater Talons
	"9qr": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Scissors Quhab
	"7ar": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Suwayyah
	"7wb": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Wrist Sword
	"7xf": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // War Fist
	"7cs": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Battle Cestus
	"7lw": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Feral Claws
	"7tw": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Runic Talons
	"7qr": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Scissors Suwayh
	"ob1": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Eagle Orb
	"ob2": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Sacred Globe
	"ob3": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Smoked Sphere
	"ob4": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Clasped Orb
	"ob5": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Jared's Stone
	"ob6": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Glowing Orb
	"ob7": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Crystalline Glb
	"ob8": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Cloudy Sphere
	"ob9": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Sparkling Ball
	"oba": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Swirling Crystl
	"obb": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Heavenly Stone
	"obc": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Eldritch Orb
	"obd": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Demon Heart
	"obe": {SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Vortex Orb
	"obf": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Dimensional Shrd
	"am1": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: AmazonWeapons},    // Stag Bow
	"am2": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: AmazonWeapons},    // Reflex Bow
	"am3": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: AmazonWeapons}, // Maiden Spear
	"am4": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: AmazonWeapons}, // Maiden Pike
	"am5": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: AmazonWeapons},          // Maiden Javelin
	"am6": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: AmazonWeapons},    // Ashwood Bow
	"am7": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: AmazonWeapons},    // Ceremonial Bow
	"am8": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: AmazonWeapons}, // Ceremonial Spr
	"am9": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: AmazonWeapons}, // Ceremonial Pike
	"ama": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: AmazonWeapons},          // Ceremonial Jav
	"amb": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: AmazonWeapons},    // Matriarchal Bow
	"amc": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: AmazonWeapons},    // Grnd Matron Bow
	"amd": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: AmazonWeapons}, // Matriarchal Spr
	"ame": {SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: AmazonWeapons}, // Matriarchal Pik
	"amf": {SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: AmazonWeapons},          // Matriarchal Jav
	"ci0": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Circlets},                   // Circlet
	"ci1": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Circlets},                   // Coronet
	"ci2": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Circlets},                   // Tiara
	"ci3": {SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Circlets},                   // Diadem
	"gps": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowablePots},           // Rancid Gas Pot
	"ops": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowablePots},           // Oil Potion
	"gpm": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowablePots},           // Choking Gas Pot
	"opm": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowablePots},           // Exploding Pot
	"gpl": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowablePots},           // Strangling Gas
	"opl": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowablePots},           // Fulminating Pot
	"leg": {SizeX: 1, SizeY: 3, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Wirt’s Leg
	"hdm": {SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Horadric Malus
	"bks": {SizeX: 2, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Scroll of Inifuss 1
	"bkd": {SizeX: 2, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Scroll of Inifuss 2
	"ass": {SizeX: 2, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Book of Skill
	"box": {SizeX: 2, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Horadric Cube
	"tr1": {SizeX: 2, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Horadric Scroll
	"msf": {SizeX: 1, SizeY: 3, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Staff of Kings
	"vip": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Viper Amulet
	"hst": {SizeX: 1, SizeY: 4, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Horadric Staff
	"xyz": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Potion of Life
	"j34": {SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // A Jade Figurine
	"g34": {SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // The Golden Bird
	"bbb": {SizeX: 2, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Lam Esen’s Tome
	"g33": {SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Gidbinn
	"qf1": {SizeX: 2, SizeY: 3, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Khalim’s Flail
	"qf2": {SizeX: 2, SizeY: 3, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Khalim’s Will
	"qey": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Khalim’s Eye
	"qhr": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Khalim’s Heart
	"qbr": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Khalim’s Brain
	"mss": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Mephisto’s Soulstone
	"hfh": {SizeX: 2, SizeY: 3, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Hellforge Hammr
	"ice": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Malah’s Potion
	"tr2": {SizeX: 2, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},               // Scroll of Resistance
	"gcv": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Chipped, ItemSubSubCategory: Amethyst},                   // Chipped Amethyst
	"gfv": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawed, ItemSubSubCategory: Amethyst},                    // Flawed Amethyst
	"gsv": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Normal, ItemSubSubCategory: Amethyst},                    // Amethyst
	"gzv": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawless, ItemSubSubCategory: Amethyst},                  // Flawless Amethyst
	"gpv": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Perfect, ItemSubSubCategory: Amethyst},                   // Perfect Amethyst
	"gcw": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Chipped, ItemSubSubCategory: Diamond},                    // Chipped Diamond
	"gfw": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawed, ItemSubSubCategory: Diamond},                     // Flawed Diamond
	"gsw": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Normal, ItemSubSubCategory: Diamond},                     // Diamond
	"glw": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawless, ItemSubSubCategory: Diamond},                   // Flawless Diamond
	"gpw": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Perfect, ItemSubSubCategory: Diamond},                    // Perfect Diamond
	"gcg": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Chipped, ItemSubSubCategory: Emerald},                    // Chipped Emerald
	"gfg": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawed, ItemSubSubCategory: Emerald},                     // Flawed Emerald
	"gsg": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Normal, ItemSubSubCategory: Emerald},                     // Emerald
	"glg": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawless, ItemSubSubCategory: Emerald},                   // Flawless Emerald
	"gpg": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Perfect, ItemSubSubCategory: Emerald},                    // Perfect Emerald
	"gcr": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Chipped, ItemSubSubCategory: Ruby},                       // Chipped Ruby
	"gfr": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawed, ItemSubSubCategory: Ruby},                        // Flawed Ruby
	"gsr": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Normal, ItemSubSubCategory: Ruby},                        // Ruby
	"glr": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawless, ItemSubSubCategory: Ruby},                      // Flawless Ruby
	"gpr": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Perfect, ItemSubSubCategory: Ruby},                       // Perfect Ruby
	"gcb": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Chipped, ItemSubSubCategory: Sapphire},                   // Chipped Sapphire
	"gfb": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawed, ItemSubSubCategory: Sapphire},                    // Flawed Sapphire
	"gsb": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Normal, ItemSubSubCategory: Sapphire},                    // Sapphire
	"glb": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawless, ItemSubSubCategory: Sapphire},                  // Flawless Sapphire
	"gpb": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Perfect, ItemSubSubCategory: Sapphire},                   // Perfect Sapphire
	"skc": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Chipped, ItemSubSubCategory: Skulls},                     // Chipped Skull
	"skf": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawed, ItemSubSubCategory: Skulls},                      // Flawed Skull
	"sku": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Normal, ItemSubSubCategory: Skulls},                      // Skull
	"skl": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawless, ItemSubSubCategory: Skulls},                    // Flawless Skull
	"skz": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Perfect, ItemSubSubCategory: Skulls},                     // Perfect Skull
	"gcy": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Chipped, ItemSubSubCategory: Topaz},                      // Chipped Topaz
	"gfy": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawed, ItemSubSubCategory: Topaz},                       // Flawed Topaz
	"gsy": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Normal, ItemSubSubCategory: Topaz},                       // Topaz
	"gly": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawless, ItemSubSubCategory: Topaz},                     // Flawless Topaz
	"gpy": {SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Perfect, ItemSubSubCategory: Topaz},                      // Perfect Topaz
	"r01": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                   // El Rune
	"r02": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                   // Eld Rune
	"r03": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                   // Tir Rune
	"r04": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                   // Nef Rune
	"r05": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                   // Eth Rune
	"r06": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                   // Ith Rune
	"r07": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                   // Tal Rune
	"r08": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                   // Ral Rune
	"r09": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                   // Ort Rune
	"r10": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                   // Thul Rune
	"r11": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                  // Amn Rune
	"r12": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                  // Sol Rune
	"r13": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                  // Shael Rune
	"r14": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                  // Dol Rune
	"r15": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                  // Hel Rune
	"r16": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                  // Io Rune
	"r17": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                  // Lum Rune
	"r18": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                  // Ko Rune
	"r19": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                  // Fal Rune
	"r20": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                  // Lem Rune
	"r21": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                  // Pul Rune
	"r22": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                  // Um Rune
	"r23": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                  // Mal Rune
	"r24": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                  // Ist Rune
	"r25": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                  // Gul Rune
	"r26": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                  // Vex Rune
	"r27": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                  // Ohm Rune
	"r28": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                  // Lo Rune
	"r29": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                  // Sur Rune
	"r30": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                  // Ber Rune
	"r31": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes31to33, ItemSubSubCategory: Rune},                  // Jah Rune
	"r32": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes31to33, ItemSubSubCategory: Rune},                  // Cham Rune
	"r33": {SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes31to33, ItemSubSubCategory: Rune},                  // Zod Rune
	"yps": {SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: MiscPotions, ItemSubSubCategory: Potion},              // Antidote Potion
	"vps": {SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: MiscPotions, ItemSubSubCategory: Potion},              // Stamina Potion
	"wms": {SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: MiscPotions, ItemSubSubCategory: Potion},              // Thawing Potion
	"hp1": {SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: HealingPotions, ItemSubSubCategory: Potion},           // Minor Healing Potion
	"mp1": {SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: ManaPotions, ItemSubSubCategory: Potion},              // Minor Mana Potion
	"hp2": {SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: HealingPotions, ItemSubSubCategory: Potion},           // Light Healing Potion
	"mp2": {SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: ManaPotions, ItemSubSubCategory: Potion},              // Light Mana Potion
	"hp3": {SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: HealingPotions, ItemSubSubCategory: Potion},           // Healing Potion
	"mp3": {SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: ManaPotions, ItemSubSubCategory: Potion},              // Mana Potion
	"hp4": {SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: HealingPotions, ItemSubSubCategory: Potion},           // Greater Healing Potion
	"mp4": {SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: ManaPotions, ItemSubSubCategory: Potion},              // Greater Mana Potion
	"hp5": {SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: HealingPotions, ItemSubSubCategory: Potion},           // Super Healing Potion
	"mp5": {SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: ManaPotions, ItemSubSubCategory: Potion},              // Super Mana Potion
	"rvs": {SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: RejuvenationPotions, ItemSubSubCategory: Potion},      // Rejuv Potion
	"rvl": {SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: RejuvenationPotions, ItemSubSubCategory: Potion},      // Full Rejuv Potion
	"cm1": {SizeX: 1, SizeY: 1, ItemCategory: Charms, ItemSubCategory: SmallCharms, ItemSubSubCategory: CHARM},                // Charm Small
	"cm2": {SizeX: 1, SizeY: 2, ItemCategory: Charms, ItemSubCategory: LargeCharms, ItemSubSubCategory: CHARM},                // Charm Large
	"cm3": {SizeX: 1, SizeY: 3, ItemCategory: Charms, ItemSubCategory: GrandCharms, ItemSubSubCategory: CHARM},                // Charm Grand
	"isc": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: ScrollsAndTomes, ItemSubSubCategory: Scroll},             // Identify Scroll
	"tsc": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: ScrollsAndTomes, ItemSubSubCategory: Scroll},             // Town Portal Scroll
	"tbk": {SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubCategory: ScrollsAndTomes, ItemSubSubCategory: Scroll},             // Tome of Town Portal
	"ibk": {SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubCategory: ScrollsAndTomes, ItemSubSubCategory: Scroll},             // Tome of Identify
	"aqv": {SizeX: 1, SizeY: 3, ItemCategory: Misc, ItemSubCategory: Consumables, ItemSubSubCategory: Other},                  // Arrows
	"cqv": {SizeX: 1, SizeY: 3, ItemCategory: Misc, ItemSubCategory: Consumables, ItemSubSubCategory: Other},                  // Bolts
	"key": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Consumables, ItemSubSubCategory: Other},                  // Key
	"gld": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Consumables, ItemSubSubCategory: Other},                  // gold
	"ear": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubSubCategory: Other},                                                // Ear
	"std": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubSubCategory: Other},                                                // Standard of Heroes
	"jew": {SizeX: 1, SizeY: 1, ItemCategory: Jewelery, ItemSubCategory: Jewels, ItemSubSubCategory: Jewel},                   // Jewel
	"amu": {SizeX: 1, SizeY: 1, ItemCategory: Jewelery, ItemSubCategory: RingsAndAmulets, ItemSubSubCategory: Amulets},        // amulet
	"rin": {SizeX: 1, SizeY: 1, ItemCategory: Jewelery, ItemSubCategory: RingsAndAmulets, ItemSubSubCategory: Rings},          // ring
	"pk1": {SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubSubCategory: UberKeys},                                             // Key of Terror
	"pk2": {SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubSubCategory: UberKeys},                                             // Key of Hate
	"pk3": {SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubSubCategory: UberKeys},                                             // Key of Destruction
	"bey": {SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubSubCategory: UberParts},                                            // Baal's Eye
	"dhn": {SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubSubCategory: UberParts},                                            // Diablo's Horn
	"mbr": {SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubSubCategory: UberParts},                                            // Mephisto's Brain
	"bet": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubSubCategory: Essences},                                             // Burning Essence of Terror
	"ceh": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubSubCategory: Essences},                                             // Charged Essence of Hatred
	"fed": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubSubCategory: Essences},                                             // Festering Essence of Destruction
	"tes": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubSubCategory: Essences},                                             // Twisted Essence of Suffering
	"toa": {SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubSubCategory: Essences},                                             // Token of Absolution
}

func getRuneNumber(id string) int {
	runeNumber, err := strconv.ParseInt(id[1:], 10, 0)
	if err != nil {
		return 0
	}
	return int(runeNumber)
}
