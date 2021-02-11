package sorter

import (
	"strconv"

	"github.com/nokka/d2s"
)

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
	ItemQuality        ItemQuality
	ItemCategory       ItemCategoryName
	ItemSubCategory    ItemSubCategoryName
	ItemSubSubCategory ItemSubSubCategoryName
	SizeX              int
	SizeY              int
}

type ItemQuality string

const (
	ItemQualityNormal      ItemQuality = "Normal"
	ItemQualityExceptional ItemQuality = "Exceptional"
	ItemQualityElite       ItemQuality = "Elite"
)

var itemPropertyMap = map[d2s.ItemType]*ItemPropertyData{
	d2s.Cap:                           {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                    // Cap
	d2s.SkullCap:                      {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                    // Skull Cap
	d2s.Helm:                          {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                    // Helm
	d2s.FullHelm:                      {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                    // Full Helm
	d2s.GreatHelm:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                    // Great Helm
	d2s.Crown:                         {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                    // Crown
	d2s.Mask:                          {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                    // Mask
	d2s.BoneHelm:                      {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                    // Bone Helm
	d2s.WarHat:                        {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // War Hat
	d2s.Sallet:                        {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Sallet
	d2s.Casque:                        {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Casque
	d2s.Basinet:                       {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Basinet
	d2s.WingedHelm:                    {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Winged Helm
	d2s.GrandCrown:                    {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Grand Crown
	d2s.DeathMask:                     {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Death Mask
	d2s.GrimHelm:                      {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},               // Grim Helm
	d2s.Shako:                         {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                     // Shako
	d2s.Hydraskull:                    {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                     // Hydraskull
	d2s.Armet:                         {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                     // Armet
	d2s.GiantConch:                    {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                     // Giant Conch
	d2s.SpiredHelm:                    {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                     // Spired Helm
	d2s.Corona:                        {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                     // Corona
	d2s.Demonhead:                     {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                     // Demonhead
	d2s.BoneVisage:                    {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: NormalHelmet},                     // Bone Visage
	d2s.QuiltedArmor:                  {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},              // Quilted Armor
	d2s.LeatherArmor:                  {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},              // Leather Armor
	d2s.HardLeatherArmor:              {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},              // Hard Leather
	d2s.StuddedLeather:                {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},              // Studded Leather
	d2s.RingMail:                      {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},              // Ring Mail
	d2s.ScaleMail:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},              // Scale Mail
	d2s.ChainMail:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},              // Chain Mail
	d2s.BreastPlate:                   {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},              // Breast Plate
	d2s.SplintMail:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},              // Splint Mail
	d2s.PlateMail:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},              // Plate Mail
	d2s.FieldPlate:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},              // Field Plate
	d2s.GothicPlate:                   {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},              // Gothic Plate
	d2s.FullPlateMail:                 {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},              // Full Plate Mail
	d2s.AncientArmor:                  {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},              // Ancient Armor
	d2s.LightPlate:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},              // Light Plate
	d2s.GhostArmor:                    {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Ghost Armor
	d2s.SerpentskinArmor:              {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Serpentskin
	d2s.DemonhideArmor:                {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Demonhide Armor
	d2s.TrellisedArmor:                {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Trellised Armor
	d2s.LinkedMail:                    {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Linked Mail
	d2s.TigulatedMail:                 {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Tigulated Mail
	d2s.MeshArmor:                     {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Mesh Armor
	d2s.Cuirass:                       {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Cuirass
	d2s.RussetArmor:                   {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Russet Armor
	d2s.TemplarCoat:                   {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Templar Coat
	d2s.SharktoothArmor:               {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Sharktooth
	d2s.EmbossedPlate:                 {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Embossed Plate
	d2s.ChaosArmor:                    {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Chaos Armor
	d2s.OrnatePlate:                   {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Ornate Armor
	d2s.MagePlate:                     {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},         // Mage Plate
	d2s.DuskShroud:                    {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},               // Dusk Shroud
	d2s.Wyrmhide:                      {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},               // Wyrmhide
	d2s.ScarabHusk:                    {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},               // Scarab Husk
	d2s.WireFleece:                    {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},               // Wire Fleece
	d2s.DiamondMail:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},               // Diamond Mail
	d2s.LoricatedMail:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},               // Loricated Mail
	d2s.Boneweave:                     {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},               // Boneweave
	d2s.GreatHauberk:                  {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},               // Great Hauberk
	d2s.BalrogSkin:                    {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},               // Balrog Skin
	d2s.HellforgedPlate:               {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},               // Hellforge Plate
	d2s.KrakenShell:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},               // Kraken Shell
	d2s.LacqueredPlate:                {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},               // Lacquered Plate
	d2s.ShadowPlate:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},               // Shadow Plate
	d2s.SacredArmor:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},               // Sacred Armor
	d2s.ArchonPlate:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: BodyArmor, ItemSubSubCategory: NormalBodyArmor},               // Archon Plate
	d2s.Buckler:                       {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                    // Buckler
	d2s.SmallShield:                   {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                    // Small Shield
	d2s.LargeShield:                   {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                    // Large Shield
	d2s.KiteShield:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                    // Kite Shield
	d2s.TowerShield:                   {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                    // Tower Shield
	d2s.GothicShield:                  {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                    // Gothic Shield
	d2s.BoneShield:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                    // Bone Shield
	d2s.SpikedShield:                  {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                    // Spiked Shield
	d2s.Defender:                      {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Defender
	d2s.RoundShield:                   {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Round Shield
	d2s.Scutum:                        {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Scutum
	d2s.DragonShield:                  {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Dragon Shield
	d2s.Pavise:                        {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Pavise
	d2s.AncientShield:                 {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Ancient Shield
	d2s.GrimShield:                    {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Grim Shield
	d2s.BarbedShield:                  {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},               // Barbed Shield
	d2s.Heater:                        {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                     // Heater
	d2s.Luna:                          {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                     // Luna
	d2s.Hyperion:                      {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                     // Hyperion
	d2s.Monarch:                       {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                     // Monarch
	d2s.Aegis:                         {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                     // Aegis
	d2s.Ward:                          {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                     // Ward
	d2s.TrollNest:                     {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                     // Troll Nest
	d2s.BladeBarrier:                  {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NormalShield},                     // Blade Barrier
	d2s.LeatherGloves:                 {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},                    // Leather Gloves
	d2s.HeavyGloves:                   {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},                    // Heavy Gloves
	d2s.Bracers:                       {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},                    // Chain Gloves
	d2s.LightGauntlets:                {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},                    // Light Gauntlets
	d2s.Gauntlets:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},                    // Gauntlets
	d2s.DemonhideGloves:               {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Demonhide Glove
	d2s.SharkskinGloves:               {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Sharkskin Glove
	d2s.HeavyBracers:                  {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Heavy Bracers
	d2s.BattleGauntlets:               {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // Battle Gauntlet
	d2s.WarGauntlets:                  {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},               // War Gauntlets
	d2s.BrambleMitts:                  {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},                     // Bramble Mitts
	d2s.VampireboneGloves:             {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},                     // Vampirebone Gl
	d2s.Vambraces:                     {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},                     // Vambraces
	d2s.CrusaderGauntlets:             {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},                     // Crusader Gaunt
	d2s.OgreGauntlets:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Gloves, ItemSubSubCategory: NormalGloves},                     // Ogre Gauntlets
	d2s.LeatherBoots:                  {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                      // Boots
	d2s.HeavyBoots:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                      // Heavy Boots
	d2s.ChainBoots:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                      // Chain Boots
	d2s.LightPlatedBoots:              {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                      // Light Plate
	d2s.Greaves:                       {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                      // Greaves
	d2s.DemonhideBoots:                {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Demonhide Boots
	d2s.SharkskinBoots:                {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Sharkskin Boots
	d2s.MeshBoots:                     {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Mesh Boots
	d2s.BattleBoots:                   {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // Battle Boots
	d2s.WarBoots:                      {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                 // War Boots
	d2s.WyrmhideBoots:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                       // Wyrmhide Boots
	d2s.ScarabshellBoots:              {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                       // Scarabshell Bts
	d2s.BoneweaveBoots:                {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                       // Boneweave Boots
	d2s.MirroredBoots:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                       // Mirrored Boots
	d2s.MyrmidonGreaves:               {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Boots, ItemSubSubCategory: NormalBoots},                       // Myrmidon Greave
	d2s.Sash:                          {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                       // Sash
	d2s.LightBelt:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                       // Light Belt
	d2s.Belt:                          {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                       // Belt
	d2s.HeavyBelt:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                       // Heavy Belt
	d2s.PlatedBelt:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                       // Plated Belt
	d2s.DemonhideSash:                 {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Demonhide Sash
	d2s.SharkskinBelt:                 {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Sharkskin Belt
	d2s.MeshBelt:                      {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Mesh Belt
	d2s.BattleBelt:                    {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // Battle Belt
	d2s.WarBelt:                       {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                  // War Belt
	d2s.SpiderwebSash:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                        // Spiderweb Sash
	d2s.VampirefangBelt:               {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                        // Vampirefang Blt
	d2s.MithrilCoil:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                        // Mithril Coil
	d2s.TrollBelt:                     {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                        // Troll Belt
	d2s.ColossusGirdle:                {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 1, ItemCategory: Armor, ItemSubCategory: Belt, ItemSubSubCategory: NormalBelts},                        // Colossus Girdle
	d2s.WolfHead:                      {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                           // Wolf Head
	d2s.HawkHelm:                      {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                           // Hawk Helm
	d2s.Antlers:                       {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                           // Antlers
	d2s.FalconMask:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                           // Falcon Mask
	d2s.SpiritMask:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                           // Spirit Mask
	d2s.AlphaHelm:                     {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Alpha Helm
	d2s.GriffonHeadress:               {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Griffon Headress
	d2s.HuntersGuise:                  {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Hunter’s Guise
	d2s.SacredFeathers:                {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Sacred Feathers
	d2s.TotemicMask:                   {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                      // Totemic Mask
	d2s.BloodSpirit:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                            // Blood Spirit
	d2s.SunSpirit:                     {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                            // Sun Spirit
	d2s.EarthSpirit:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                            // Earth Spirit
	d2s.SkySpirit:                     {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                            // Sky Spirit
	d2s.DreamSpirit:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Pelts},                            // Dream Spirit
	d2s.JawboneCap:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},                // Jawbone Cap
	d2s.FangedHelm:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},                // Fanged Helm
	d2s.HornedHelm:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},                // Horned Helm
	d2s.AssaultHelmet:                 {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},                // Assualt Helmet
	d2s.AvengerGuard:                  {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},                // Avenger Guard
	d2s.JawboneVisor:                  {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Jawbone Visor
	d2s.LionHelm:                      {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Lion Helm
	d2s.RageMask:                      {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Rage Mask
	d2s.SavageHelmet:                  {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Savage Helmet
	d2s.SlayerGuard:                   {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},           // Slayer Guard
	d2s.CarnageHelm:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},                 // Carnage Helm
	d2s.FuryVisor:                     {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},                 // Fury Visor
	d2s.DestroyerHelm:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},                 // Destroyer Helm
	d2s.ConquerorCrown:                {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},                 // Conqueror Crown
	d2s.GuardianCrown:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: BarbarianHelmets},                 // Guardian Crown
	d2s.Targe:                         {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},                  // Targe
	d2s.Rondache:                      {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},                  // Rondache
	d2s.HeraldicShield:                {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},                  // Heraldic Shield
	d2s.AerinShield:                   {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},                  // Aerin Shield
	d2s.CrownShield:                   {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},                  // Crown Shield
	d2s.AkaranTarge:                   {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Akaran Targe
	d2s.AkaranRondache:                {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Akaran Rondache
	d2s.ProtectorShield:               {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Protector Shld
	d2s.GildedShield:                  {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Guilded Shield
	d2s.RoyalShield:                   {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},             // Royal Shield
	d2s.SacredTarge:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},                   // Sacred Targe
	d2s.SacredRondache:                {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},                   // Sacred Rondache
	d2s.KurastShield:                  {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},                   // Kurast Shield
	d2s.ZakarumShield:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},                   // Zakarum Shield
	d2s.VortexShield:                  {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: PaladinShields},                   // Vortex Shield
	d2s.PreservedHead:                 {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                      // Preserved Head
	d2s.ZombieHead:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                      // Zombie Head
	d2s.UnravellerHead:                {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                      // Unraveller Head
	d2s.GargoyleHead:                  {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                      // Gargoyle Head
	d2s.DemonHead:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                      // Demon Head
	d2s.MummifiedTrophy:               {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Mummified Trphy
	d2s.FetishTrophy:                  {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Fetish Trophy
	d2s.SextonTrophy:                  {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Sexton Trophy
	d2s.CantorTrophy:                  {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Cantor Trophy
	d2s.HeirophantTrophy:              {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                 // Heirophant Trphy
	d2s.MinionSkull:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                       // Minion Skull
	d2s.HellspawnSkull:                {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                       // Hellspawn Skull
	d2s.OverseerSkull:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                       // Overseer Skull
	d2s.SuccubusSkull:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                       // Succubae Skull
	d2s.BloodlordSkull:                {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Shield, ItemSubSubCategory: NecroHeads},                       // Bloodlord Skull
	d2s.HandAxe:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                             // Hand Axe
	d2s.Axe:                           {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                             // Axe
	d2s.DoubleAxe:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                             // Double Axe
	d2s.MilitaryPick:                  {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                             // Military Pick
	d2s.WarAxe:                        {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                             // War Axe
	d2s.LargeAxe:                      {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                             // Large Axe
	d2s.BroadAxe:                      {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                             // Broad Axe
	d2s.BattleAxe:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                             // Battle Axe
	d2s.GreatAxe:                      {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                             // Great Axe
	d2s.GiantAxe:                      {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                             // Giant Axe
	d2s.Hatchet:                       {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Hatchet
	d2s.Cleaver:                       {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Cleaver
	d2s.TwinAxe:                       {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Twin Axe
	d2s.Crowbill:                      {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Crowbill
	d2s.Naga:                          {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Naga
	d2s.MilitaryAxe:                   {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Military Axe
	d2s.BeardedAxe:                    {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Bearded Axe
	d2s.Tabar:                         {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Tabar
	d2s.GothicAxe:                     {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Gothic Axe
	d2s.AncientAxe:                    {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                        // Ancient Axe
	d2s.Tomahawk:                      {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                              // Tomahawk
	d2s.SmallCrescent:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                              // Small Crescent
	d2s.EttinAxe:                      {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                              // Ettin Axe
	d2s.WarSpike:                      {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                              // War Spike
	d2s.BerserkerAxe:                  {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                              // Berserker Axe
	d2s.FeralAxe:                      {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                              // Feral Axe
	d2s.SilverEdgedAxe:                {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                              // Silver Edged Ax
	d2s.Decapitator:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                              // Decapitator
	d2s.ChampionAxe:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                              // Champion Axe
	d2s.GloriousAxe:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Axes, ItemSubSubCategory: Axe},                              // Glorious Axe
	d2s.Club:                          {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},                // Club
	d2s.SpikedClub:                    {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},                // Spiked Club
	d2s.Mace:                          {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},                // Mace
	d2s.MorningStar:                   {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},                // Morning Star
	d2s.Flail:                         {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Flails},               // Flail
	d2s.WarHammer:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Hammers},              // War Hammer
	d2s.Maul:                          {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Hammers},              // Maul
	d2s.GreatMaul:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Hammers},              // Great Maul
	d2s.Cudgel:                        {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Cudgel
	d2s.BarbedClub:                    {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Barbed Club
	d2s.FlangedMace:                   {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Flanged Mace
	d2s.JaggedStar:                    {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Jagged Star
	d2s.Knout:                         {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Knout
	d2s.BattleHammer:                  {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Battle Hammer
	d2s.WarClub:                       {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // War Club
	d2s.MarteldeFer:                   {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},           // Martel de Fer
	d2s.Truncheon:                     {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},                 // Truncheon
	d2s.TyrantClub:                    {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},                 // Tyrant Club
	d2s.ReinforcedMace:                {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},                 // Reinforced Mace
	d2s.DevilStar:                     {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},                 // Devil Star
	d2s.Scourge:                       {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Maces},                 // Scourge
	d2s.LegendaryMallet:               {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Hammers},               // Legendary Mallet
	d2s.OgreMaul:                      {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Hammers},               // Ogre Maul
	d2s.ThunderMaul:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: HammersAndMaces, ItemSubSubCategory: Hammers},               // Thunder Maul
	d2s.ShortSword:                    {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},              // Short Sword
	d2s.Scimitar:                      {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},              // Scimitar
	d2s.Sabre:                         {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},              // Saber
	d2s.Falchion:                      {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},              // Falchion
	d2s.CrystalSword:                  {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},              // Crystal Sword
	d2s.BroadSword:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},              // Broad Sword
	d2s.LongSword:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},              // Long Sword
	d2s.WarSword:                      {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},              // War Sword
	d2s.TwoHandedSword:                {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},              // Two-handed Swrd
	d2s.Claymore:                      {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},              // Claymore
	d2s.GiantSword:                    {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},              // Giant Sword
	d2s.BastardSword:                  {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},              // Bastard Sword
	d2s.Flamberge:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},              // Flamberge
	d2s.GreatSword:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},              // Great Sword
	d2s.Gladius:                       {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Gladius
	d2s.Cutlass:                       {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Cutlass
	d2s.Shamshir:                      {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Shamshir
	d2s.Tulwar:                        {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Tulwar
	d2s.DimensionalBlade:              {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Dimensional Bld
	d2s.BattleSword:                   {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Battle Sword
	d2s.RuneSword:                     {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Rune Sword
	d2s.AncientSword:                  {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Ancient Sword
	d2s.Espadon:                       {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Espadon
	d2s.DacianFalx:                    {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Dacian Falx
	d2s.TuskSword:                     {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Tusk Sword
	d2s.GothicSword:                   {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Gothic Sword
	d2s.Zweihander:                    {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Zweihander
	d2s.ExecutionerSword:              {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},         // Executioner Swr
	d2s.Falcata:                       {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},               // Falcata
	d2s.Ataghan:                       {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},               // Ataghan
	d2s.ElegantBlade:                  {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},               // Elegant Blade
	d2s.HydraEdge:                     {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},               // Hydra Edge
	d2s.PhaseBlade:                    {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},               // Phase Blade
	d2s.ConquestSword:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},               // Conquest Sword
	d2s.CrypticSword:                  {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},               // Cryptic Sword
	d2s.MythicalSword:                 {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},               // Mythical Sword
	d2s.LegendSword:                   {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},               // Legend Sword
	d2s.HighlandBlade:                 {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},               // Highland Blade
	d2s.BalrogBlade:                   {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},               // Balrog Blade
	d2s.ChampionSword:                 {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},               // Champion Sword
	d2s.ColossalSword:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},               // Colossal Sword
	d2s.ColossalBlade:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Swords},               // Colossus Blade
	d2s.Dagger:                        {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},             // Dagger
	d2s.Dirk:                          {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},             // Dirk
	d2s.Kriss:                         {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},             // Kriss
	d2s.Blade:                         {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},             // Blade
	d2s.Poignard:                      {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Poignard
	d2s.Rondel:                        {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Rondel
	d2s.Cinquedeas:                    {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Cinquedeas
	d2s.Stilleto:                      {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},        // Stilleto
	d2s.BoneKnife:                     {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},              // Bone Knife
	d2s.MithralPoint:                  {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},              // Mithral Point
	d2s.FangedKnife:                   {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},              // Fanged Knife
	d2s.LegendSpike:                   {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: SwordsAndDaggers, ItemSubSubCategory: Daggers},              // Legend Spike
	d2s.ThrowingKnife:                 {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableKnife},            // Throwing Knife
	d2s.ThrowingAxe:                   {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableAxe},              // Throwing Axe
	d2s.BalancedKnife:                 {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableKnife},            // Balanced Knife
	d2s.BalancedAxe:                   {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableAxe},              // Balanced Axe
	d2s.BattleDart:                    {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableKnife},       // Battle Dart
	d2s.Francisca:                     {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableAxe},         // Francisca
	d2s.WarDart:                       {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableKnife},       // War Dart
	d2s.Hurlbat:                       {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableAxe},         // Hurlbat
	d2s.FlyingKnife:                   {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableKnife},             // Flying Knife
	d2s.FlyingAxe:                     {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableAxe},               // Flying Axe
	d2s.WingedKnife:                   {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableKnife},             // Winged Knife
	d2s.WingedAxe:                     {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowableAxe},               // Winged Axe
	d2s.Javelin:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                     // Javelin
	d2s.Pilum:                         {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                     // Pilum
	d2s.ShortSpear:                    {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                     // Short Spear
	d2s.Glaive:                        {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                     // Glaive
	d2s.ThrowingSpear:                 {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                     // Throwing Spear
	d2s.WarJavelin:                    {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // War Javelin
	d2s.GreatPilum:                    {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Great Pilum
	d2s.Simbilan:                      {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Simbilan
	d2s.Spiculum:                      {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Spiculum
	d2s.Harpoon:                       {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                // Harpoon
	d2s.HyperionJavelin:               {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                      // Hyperion Javeln
	d2s.StygianPilum:                  {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                      // Stygian Pilum
	d2s.BalrogSpear:                   {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                      // Balrog Spear
	d2s.GhostGlaive:                   {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                      // Ghost Glaive
	d2s.WingedHarpoon:                 {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: Javelin},                      // Winged Harpoon
	d2s.Spear:                         {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},             // Spear
	d2s.Trident:                       {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},             // Trident
	d2s.Brandistock:                   {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},             // Brandistock
	d2s.Spetum:                        {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},             // Spetum
	d2s.Pike:                          {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},             // Pike
	d2s.WarSpear:                      {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // War Spear
	d2s.Fuscina:                       {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // Fuscina
	d2s.WarFork:                       {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // War Fork
	d2s.Yari:                          {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // Yari
	d2s.Lance:                         {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},        // Lance
	d2s.HyperionSpear:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},              // Hyperion Spear
	d2s.StygianPike:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},              // Stygian Pike
	d2s.Mancatcher:                    {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},              // Mancatcher
	d2s.GhostSpear:                    {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},              // Ghost Spear
	d2s.WarPike:                       {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Spears},              // War Pike
	d2s.Bardiche:                      {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},           // Bardiche
	d2s.Voulge:                        {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},           // Voulge
	d2s.Scythe:                        {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},           // Scythe
	d2s.Poleaxe:                       {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},           // Poleaxe
	d2s.Halberd:                       {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},           // Halberd
	d2s.WarScythe:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},           // War Scythe
	d2s.LochaberAxe:                   {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Lochaber Axe
	d2s.Bill:                          {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Bill
	d2s.BattleScythe:                  {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Battle Scythe
	d2s.Partizan:                      {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Partizan
	d2s.BecDeCorbin:                   {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Bec-de-Corbin
	d2s.GrimScythe:                    {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},      // Grim Scythe
	d2s.OgreAxe:                       {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},            // Ogre Axe
	d2s.ColossusVoulge:                {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},            // Colossus Voulge
	d2s.Thresher:                      {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},            // Thresher
	d2s.CrypticAxe:                    {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},            // Cryptic Axe
	d2s.GreatPoleaxe:                  {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},            // Great Poleaxe
	d2s.GiantThresher:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: Polearms},            // Giant Thresher
	d2s.ShortBow:                      {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                  // Short Bow
	d2s.HuntersBow:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                  // Hunter’s Bow
	d2s.LongBow:                       {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                  // Long Bow
	d2s.CompositeBow:                  {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                  // Composite Bow
	d2s.ShortBattleBow:                {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                  // Shrt Battle Bow
	d2s.LongBattleBow:                 {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                  // Long Battle Bow
	d2s.ShortWarBow:                   {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                  // Short War Bow
	d2s.LongWarBow:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                  // Long War Bow
	d2s.EdgeBow:                       {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Edge Bow
	d2s.RazorBow:                      {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Razor Bow
	d2s.CedarBow:                      {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Cedar Bow
	d2s.DoubleBow:                     {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Double Bow
	d2s.ShortSiegeBow:                 {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Short Siege Bow
	d2s.LongSiegeBow:                  {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Long Siege Bow
	d2s.RuneBow:                       {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Rune Bow
	d2s.GothicBow:                     {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},             // Gothic Bow
	d2s.SpiderBow:                     {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                   // Spider Bow
	d2s.BladeBow:                      {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                   // Blade Bow
	d2s.ShadowBow:                     {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                   // Shadow Bow
	d2s.GreatBow:                      {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                   // Great Bow
	d2s.DiamondBow:                    {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                   // Diamond Bow
	d2s.CrusaderBow:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                   // Crusader Bow
	d2s.WardBow:                       {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                   // Ward Bow
	d2s.HydraBow:                      {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Bows},                   // Hydra Bow
	d2s.LightCrossbow:                 {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},             // Light Crossbow
	d2s.Crossbow:                      {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},             // Crossbow
	d2s.HeavyCrossbow:                 {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},             // Heavy Crossbow
	d2s.RepeatingCrossbow:             {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},             // Repeating X-bow
	d2s.Arbalest:                      {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Arbalest
	d2s.SiegeCrossbow:                 {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Siege Crossbow
	d2s.Balista:                       {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Ballista
	d2s.ChuKoNu:                       {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},        // Chu-Ko-Nu
	d2s.PelletBow:                     {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},              // Pellet Bow
	d2s.GorgonCrossbow:                {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},              // Gorgon Crossbow
	d2s.ColossusCrossbow:              {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},              // Colossus x-bow
	d2s.DemonCrossbow:                 {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: Crossbows},              // Demon Crossbow
	d2s.ShortStaff:                    {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},                // Short Staff
	d2s.LongStaff:                     {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},                // Long Staff
	d2s.GnarledStaff:                  {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},                // Gnarled Staff
	d2s.BattleStaff:                   {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},                // Battle Staff
	d2s.WarStaff:                      {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},                // War Staff
	d2s.JoStaff:                       {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Jo Staff
	d2s.QuarterStaff:                  {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Quarterstaff
	d2s.CedarStaff:                    {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Cedar Staff
	d2s.GothicStaff:                   {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Gothic Staff
	d2s.RuneStaff:                     {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},           // Rune Staff
	d2s.WalkingStick:                  {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},                 // Walking Stick
	d2s.Stalagmite:                    {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},                 // Stalagmite
	d2s.ElderStaff:                    {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},                 // Elder Staff
	d2s.Shillelah:                     {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},                 // Shillelagh
	d2s.ArchonStaff:                   {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Staves},                 // Archon Staff
	d2s.Wand:                          {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},                 // Wand
	d2s.YewWand:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},                 // Yew Wand
	d2s.BoneWand:                      {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},                 // Bone Wand
	d2s.GrimWand:                      {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},                 // Grim Wand
	d2s.BurntWand:                     {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Burnt Wand
	d2s.PetrifiedWand:                 {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Petrified Wand
	d2s.TombWand:                      {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Tomb Wand
	d2s.GraveWand:                     {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},            // Grave Wand
	d2s.PolishedWand:                  {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},                  // Polished Wand
	d2s.GhostWand:                     {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},                  // Ghost Wand
	d2s.LichWand:                      {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},                  // Lich Wand
	d2s.UnearthedWand:                 {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: Wands},                  // Unearthed Wand
	d2s.Scepter:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                     // Sceptre
	d2s.GrandScepter:                  {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                     // Grand Sceptre
	d2s.WarScepter:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                     // War Sceptre
	d2s.RuneScepter:                   {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                // Rune Sceptre
	d2s.HolyWaterSprinkler:            {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                // Holy Water Spri
	d2s.DivineScepter:                 {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                // Divine Sceptre
	d2s.MightyScepter:                 {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                      // Mighty Sceptre
	d2s.SeraphRod:                     {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                      // Seraph Rod
	d2s.Caduceus:                      {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Scepters, ItemSubSubCategory: SCEPTER},                      // Caduceus
	d2s.Katar:                         {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},                // Katar
	d2s.WristBlade:                    {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},                // Wrist Blade
	d2s.HatchetHands:                  {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},                // Hatchet Hands
	d2s.Cestus:                        {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},                // Cestus
	d2s.Claws:                         {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},                // Claws
	d2s.BladeTalons:                   {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},                // Blade Talons
	d2s.ScissorsKatar:                 {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},                // Scissors Katar
	d2s.Quhab:                         {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Quhab
	d2s.WristSpike:                    {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Wrist Spike
	d2s.Fascia:                        {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Fascia
	d2s.HandScythe:                    {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Hand Scythe
	d2s.GreaterClaws:                  {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Greater Claws
	d2s.GreaterTalons:                 {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Greater Talons
	d2s.ScissorsQuhab:                 {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},           // Scissors Quhab
	d2s.Suwayyah:                      {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},                 // Suwayyah
	d2s.WristSword:                    {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},                 // Wrist Sword
	d2s.WarFist:                       {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},                 // War Fist
	d2s.BattleCestus:                  {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},                 // Battle Cestus
	d2s.FeralClaws:                    {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},                 // Feral Claws
	d2s.RunicTalons:                   {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},                 // Runic Talons
	d2s.ScissorsSuwayyah:              {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Claws, ItemSubSubCategory: AssassinWeapons},                 // Scissors Suwayh
	d2s.EagleOrb:                      {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},           // Eagle Orb
	d2s.SacredGlobe:                   {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},           // Sacred Globe
	d2s.SmokedSphere:                  {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},           // Smoked Sphere
	d2s.ClaspedOrb:                    {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},           // Clasped Orb
	d2s.JaredsStone:                   {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},           // Jared's Stone
	d2s.GlowingOrb:                    {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Glowing Orb
	d2s.CrystallineGlobe:              {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Crystalline Glb
	d2s.CloudySphere:                  {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Cloudy Sphere
	d2s.SparklingBall:                 {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Sparkling Ball
	d2s.SwirlingCrystal:               {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},      // Swirling Crystl
	d2s.HeavenlyStone:                 {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},            // Heavenly Stone
	d2s.EldritchOrb:                   {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},            // Eldritch Orb
	d2s.DemonHeart:                    {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},            // Demon Heart
	d2s.VortexOrb:                     {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 2, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},            // Vortex Orb
	d2s.DimensionalShard:              {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: StavesAndWands, ItemSubSubCategory: SorcWeapons},            // Dimensional Shrd
	d2s.StagBow:                       {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: AmazonWeapons},         // Stag Bow
	d2s.ReflexBow:                     {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: AmazonWeapons},         // Reflex Bow
	d2s.MaidenSpear:                   {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: AmazonWeapons},      // Maiden Spear
	d2s.MaidenPike:                    {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: AmazonWeapons},      // Maiden Pike
	d2s.MaidenJavelin:                 {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: AmazonWeapons},               // Maiden Javelin
	d2s.AshwoodBow:                    {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: AmazonWeapons},    // Ashwood Bow
	d2s.CeremonialBow:                 {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: AmazonWeapons},    // Ceremonial Bow
	d2s.CeremonialSpear:               {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: AmazonWeapons}, // Ceremonial Spr
	d2s.CeremonialPike:                {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: AmazonWeapons}, // Ceremonial Pike
	d2s.CeremonialJavelin:             {ItemQuality: ItemQualityExceptional, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: AmazonWeapons},          // Ceremonial Jav
	d2s.MatriarchalBow:                {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: AmazonWeapons},          // Matriarchal Bow
	d2s.GrandMatronBow:                {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: BowAndCrossbow, ItemSubSubCategory: AmazonWeapons},          // Grnd Matron Bow
	d2s.MatriarchalSpear:              {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: AmazonWeapons},       // Matriarchal Spr
	d2s.MatriarchalPike:               {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 4, ItemCategory: Weapons, ItemSubCategory: PolearmsAndSpears, ItemSubSubCategory: AmazonWeapons},       // Matriarchal Pik
	d2s.MatriarchalJavelin:            {ItemQuality: ItemQualityElite, SizeX: 1, SizeY: 3, ItemCategory: Weapons, ItemSubCategory: Javelins, ItemSubSubCategory: AmazonWeapons},                // Matriarchal Jav
	d2s.Circlet:                       {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Circlets},                        // Circlet
	d2s.Coronet:                       {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Circlets},                        // Coronet
	d2s.Tiara:                         {ItemQuality: ItemQualityExceptional, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Circlets},                   // Tiara
	d2s.Diadem:                        {ItemQuality: ItemQualityElite, SizeX: 2, SizeY: 2, ItemCategory: Armor, ItemSubCategory: Helmet, ItemSubSubCategory: Circlets},                         // Diadem
	d2s.RancidGasPotion:               {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowablePots},                // Rancid Gas Pot
	d2s.OilPotion:                     {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowablePots},                // Oil Potion
	d2s.ChokingGasPotion:              {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowablePots},                // Choking Gas Pot
	d2s.ExplodingPotion:               {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowablePots},                // Exploding Pot
	d2s.StranglingGasPotion:           {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowablePots},                // Strangling Gas
	d2s.FulminatingPotion:             {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Throwables, ItemSubSubCategory: ThrowablePots},                // Fulminating Pot
	d2s.WirtsLeg:                      {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Wirt’s Leg
	d2s.HoradricMalus:                 {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Horadric Malus
	d2s.ScrollOfInifuss1:              {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Scroll of Inifuss 1
	d2s.ScrollOfInifuss2:              {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Scroll of Inifuss 2
	d2s.BookofSkill:                   {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Book of Skill
	d2s.HoradricCube:                  {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Horadric Cube
	d2s.HoradricScroll:                {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Horadric Scroll
	d2s.StaffOfTheKings:               {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Staff of Kings
	d2s.ViperAmulet:                   {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Viper Amulet
	d2s.HoradricStaff:                 {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 4, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Horadric Staff
	d2s.PotionofLife:                  {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Potion of Life
	d2s.JadeFigurine:                  {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // A Jade Figurine
	d2s.TheGoldenBird:                 {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // The Golden Bird
	d2s.LamEsensTome:                  {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Lam Esen’s Tome
	d2s.Gidbinn:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Gidbinn
	d2s.KhalimsFlail:                  {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Khalim’s Flail
	d2s.KhalimsWill:                   {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Khalim’s Will
	d2s.KhalimsEye:                    {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Khalim’s Eye
	d2s.KhalimsHeart:                  {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Khalim’s Heart
	d2s.KhalimsBrain:                  {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Khalim’s Brain
	d2s.MephistoSoulStone:             {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Mephisto’s Soulstone
	d2s.HellforgeHammer:               {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 3, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Hellforge Hammr
	d2s.MalahsPotion:                  {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Malah’s Potion
	d2s.ScrollOfResistance:            {ItemQuality: ItemQualityNormal, SizeX: 2, SizeY: 2, ItemCategory: Misc, ItemSubCategory: QuestItems, ItemSubSubCategory: QuestItem},                    // Scroll of Resistance
	d2s.ChippedAmethyst:               {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Chipped, ItemSubSubCategory: Amethyst},                        // Chipped Amethyst
	d2s.FlawedAmethyst:                {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawed, ItemSubSubCategory: Amethyst},                         // Flawed Amethyst
	d2s.Amethyst:                      {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Normal, ItemSubSubCategory: Amethyst},                         // Amethyst
	d2s.FlawlessAmethyst:              {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawless, ItemSubSubCategory: Amethyst},                       // Flawless Amethyst
	d2s.PerfectAmethyst:               {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Perfect, ItemSubSubCategory: Amethyst},                        // Perfect Amethyst
	d2s.ChippedDiamond:                {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Chipped, ItemSubSubCategory: Diamond},                         // Chipped Diamond
	d2s.FlawedDiamond:                 {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawed, ItemSubSubCategory: Diamond},                          // Flawed Diamond
	d2s.Diamond:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Normal, ItemSubSubCategory: Diamond},                          // Diamond
	d2s.FlawlessDiamond:               {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawless, ItemSubSubCategory: Diamond},                        // Flawless Diamond
	d2s.PerfectDiamond:                {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Perfect, ItemSubSubCategory: Diamond},                         // Perfect Diamond
	d2s.ChippedEmerald:                {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Chipped, ItemSubSubCategory: Emerald},                         // Chipped Emerald
	d2s.FlawedEmerald:                 {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawed, ItemSubSubCategory: Emerald},                          // Flawed Emerald
	d2s.Emerald:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Normal, ItemSubSubCategory: Emerald},                          // Emerald
	d2s.FlawlessEmerald:               {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawless, ItemSubSubCategory: Emerald},                        // Flawless Emerald
	d2s.PerfectEmerald:                {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Perfect, ItemSubSubCategory: Emerald},                         // Perfect Emerald
	d2s.ChippedRuby:                   {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Chipped, ItemSubSubCategory: Ruby},                            // Chipped Ruby
	d2s.FlawedRuby:                    {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawed, ItemSubSubCategory: Ruby},                             // Flawed Ruby
	d2s.Ruby:                          {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Normal, ItemSubSubCategory: Ruby},                             // Ruby
	d2s.FlawlessRuby:                  {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawless, ItemSubSubCategory: Ruby},                           // Flawless Ruby
	d2s.PerfectRuby:                   {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Perfect, ItemSubSubCategory: Ruby},                            // Perfect Ruby
	d2s.ChippedSapphire:               {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Chipped, ItemSubSubCategory: Sapphire},                        // Chipped Sapphire
	d2s.FlawedSapphire:                {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawed, ItemSubSubCategory: Sapphire},                         // Flawed Sapphire
	d2s.Sapphire:                      {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Normal, ItemSubSubCategory: Sapphire},                         // Sapphire
	d2s.FlawlessSapphire:              {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawless, ItemSubSubCategory: Sapphire},                       // Flawless Sapphire
	d2s.PerfectSapphire:               {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Perfect, ItemSubSubCategory: Sapphire},                        // Perfect Sapphire
	d2s.ChippedSkull:                  {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Chipped, ItemSubSubCategory: Skulls},                          // Chipped Skull
	d2s.FlawedSkull:                   {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawed, ItemSubSubCategory: Skulls},                           // Flawed Skull
	d2s.Skull:                         {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Normal, ItemSubSubCategory: Skulls},                           // Skull
	d2s.FlawlessSkull:                 {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawless, ItemSubSubCategory: Skulls},                         // Flawless Skull
	d2s.PerfectSkull:                  {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Perfect, ItemSubSubCategory: Skulls},                          // Perfect Skull
	d2s.ChippedTopaz:                  {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Chipped, ItemSubSubCategory: Topaz},                           // Chipped Topaz
	d2s.FlawedTopaz:                   {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawed, ItemSubSubCategory: Topaz},                            // Flawed Topaz
	d2s.Topaz:                         {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Normal, ItemSubSubCategory: Topaz},                            // Topaz
	d2s.FlawlessTopaz:                 {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Flawless, ItemSubSubCategory: Topaz},                          // Flawless Topaz
	d2s.PerfectTopaz:                  {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Gems, ItemSubCategory: Perfect, ItemSubSubCategory: Topaz},                           // Perfect Topaz
	d2s.ElRune:                        {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                        // El Rune
	d2s.EldRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                        // Eld Rune
	d2s.TirRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                        // Tir Rune
	d2s.NefRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                        // Nef Rune
	d2s.EthRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                        // Eth Rune
	d2s.IthRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                        // Ith Rune
	d2s.TalRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                        // Tal Rune
	d2s.RalRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                        // Ral Rune
	d2s.OrtRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                        // Ort Rune
	d2s.ThulRune:                      {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes1to10, ItemSubSubCategory: Rune},                        // Thul Rune
	d2s.AmnRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                       // Amn Rune
	d2s.SolRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                       // Sol Rune
	d2s.ShaelRune:                     {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                       // Shael Rune
	d2s.DolRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                       // Dol Rune
	d2s.HelRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                       // Hel Rune
	d2s.IoRune:                        {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                       // Io Rune
	d2s.LumRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                       // Lum Rune
	d2s.KoRune:                        {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                       // Ko Rune
	d2s.FalRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                       // Fal Rune
	d2s.LemRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes11to20, ItemSubSubCategory: Rune},                       // Lem Rune
	d2s.PulRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                       // Pul Rune
	d2s.UmRune:                        {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                       // Um Rune
	d2s.MalRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                       // Mal Rune
	d2s.IstRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                       // Ist Rune
	d2s.GulRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                       // Gul Rune
	d2s.VexRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                       // Vex Rune
	d2s.OhmRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                       // Ohm Rune
	d2s.LoRune:                        {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                       // Lo Rune
	d2s.SurRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                       // Sur Rune
	d2s.BerRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes20to30, ItemSubSubCategory: Rune},                       // Ber Rune
	d2s.JahRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes31to33, ItemSubSubCategory: Rune},                       // Jah Rune
	d2s.ChamRune:                      {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes31to33, ItemSubSubCategory: Rune},                       // Cham Rune
	d2s.ZodRune:                       {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Runes, ItemSubCategory: Runes31to33, ItemSubSubCategory: Rune},                       // Zod Rune
	d2s.AntidotePotion:                {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: MiscPotions, ItemSubSubCategory: Potion},                   // Antidote Potion
	d2s.StaminaPotion:                 {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: MiscPotions, ItemSubSubCategory: Potion},                   // Stamina Potion
	d2s.ThawingPotion:                 {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: MiscPotions, ItemSubSubCategory: Potion},                   // Thawing Potion
	d2s.MinorHealingPotion:            {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: HealingPotions, ItemSubSubCategory: Potion},                // Minor Healing Potion
	d2s.MinorManaPotion:               {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: ManaPotions, ItemSubSubCategory: Potion},                   // Minor Mana Potion
	d2s.LightHealingPotion:            {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: HealingPotions, ItemSubSubCategory: Potion},                // Light Healing Potion
	d2s.LightManaPotion:               {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: ManaPotions, ItemSubSubCategory: Potion},                   // Light Mana Potion
	d2s.HealingPotion:                 {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: HealingPotions, ItemSubSubCategory: Potion},                // Healing Potion
	d2s.ManaPotion:                    {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: ManaPotions, ItemSubSubCategory: Potion},                   // Mana Potion
	d2s.GreaterHealingPotion:          {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: HealingPotions, ItemSubSubCategory: Potion},                // Greater Healing Potion
	d2s.GreaterManaPotion:             {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: ManaPotions, ItemSubSubCategory: Potion},                   // Greater Mana Potion
	d2s.SuperHealingPotion:            {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: HealingPotions, ItemSubSubCategory: Potion},                // Super Healing Potion
	d2s.SuperManaPotion:               {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: ManaPotions, ItemSubSubCategory: Potion},                   // Super Mana Potion
	d2s.RejuvenationPotion:            {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: RejuvenationPotions, ItemSubSubCategory: Potion},           // Rejuv Potion
	d2s.FullRejuvenationPotion:        {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Potions, ItemSubCategory: RejuvenationPotions, ItemSubSubCategory: Potion},           // Full Rejuv Potion
	d2s.SmallCharm:                    {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Charms, ItemSubCategory: SmallCharms, ItemSubSubCategory: CHARM},                     // Charm Small
	d2s.LargeCharm:                    {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Charms, ItemSubCategory: LargeCharms, ItemSubSubCategory: CHARM},                     // Charm Large
	d2s.GrandCharm:                    {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Charms, ItemSubCategory: GrandCharms, ItemSubSubCategory: CHARM},                     // Charm Grand
	d2s.IdentifyScroll:                {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: ScrollsAndTomes, ItemSubSubCategory: Scroll},                  // Identify Scroll
	d2s.ScrollOfTownPortal:            {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: ScrollsAndTomes, ItemSubSubCategory: Scroll},                  // Town Portal Scroll
	d2s.TomeOfTownPortal:              {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubCategory: ScrollsAndTomes, ItemSubSubCategory: Scroll},                  // Tome of Town Portal
	d2s.TomeOfIdentify:                {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubCategory: ScrollsAndTomes, ItemSubSubCategory: Scroll},                  // Tome of Identify
	d2s.Arrows:                        {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Misc, ItemSubCategory: Consumables, ItemSubSubCategory: Other},                       // Arrows
	d2s.Bolts:                         {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 3, ItemCategory: Misc, ItemSubCategory: Consumables, ItemSubSubCategory: Other},                       // Bolts
	d2s.Key:                           {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Consumables, ItemSubSubCategory: Other},                       // Key
	d2s.Gold:                          {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubCategory: Consumables, ItemSubSubCategory: Other},                       // gold
	d2s.PlayerEar:                     {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubSubCategory: Other},                                                     // Ear
	d2s.StandardofHeroes:              {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubSubCategory: Other},                                                     // Standard of Heroes
	d2s.Jewel:                         {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Jewelery, ItemSubCategory: Jewels, ItemSubSubCategory: Jewel},                        // Jewel
	d2s.Amulet:                        {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Jewelery, ItemSubCategory: RingsAndAmulets, ItemSubSubCategory: Amulets},             // amulet
	d2s.Ring:                          {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Jewelery, ItemSubCategory: RingsAndAmulets, ItemSubSubCategory: Rings},               // ring
	d2s.KeyofTerror:                   {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubSubCategory: UberKeys},                                                  // Key of Terror
	d2s.KeyofHate:                     {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubSubCategory: UberKeys},                                                  // Key of Hate
	d2s.KeyofDestruction:              {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubSubCategory: UberKeys},                                                  // Key of Destruction
	d2s.BaalsEye:                      {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubSubCategory: UberParts},                                                 // Baal's Eye
	d2s.DiablosHorn:                   {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubSubCategory: UberParts},                                                 // Diablo's Horn
	d2s.MephistosBrain:                {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 2, ItemCategory: Misc, ItemSubSubCategory: UberParts},                                                 // Mephisto's Brain
	d2s.BurningEssenceofTerror:        {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubSubCategory: Essences},                                                  // Burning Essence of Terror
	d2s.ChargedEssenceofHatred:        {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubSubCategory: Essences},                                                  // Charged Essence of Hatred
	d2s.FesteringEssenceofDestruction: {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubSubCategory: Essences},                                                  // Festering Essence of Destruction
	d2s.TwistedEssenceofSuffering:     {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubSubCategory: Essences},                                                  // Twisted Essence of Suffering
	d2s.TokenofAbsolution:             {ItemQuality: ItemQualityNormal, SizeX: 1, SizeY: 1, ItemCategory: Misc, ItemSubSubCategory: Essences},                                                  // Token of Absolution
}

func getRuneNumber(id string) int {
	runeNumber, err := strconv.ParseInt(id[1:], 10, 0)
	if err != nil {
		return 0
	}
	return int(runeNumber)
}

type UniqueItem struct {
	Name     string
	ItemType d2s.ItemType
}

type SetItem struct {
	Name     string
	SetName  string
	ItemType d2s.ItemType
}

var UniqueItems = map[uint64]UniqueItem{
	0:   {Name: "The Gnasher"},
	1:   {Name: "Deathspade"},
	2:   {Name: "Bladebone"},
	3:   {Name: "Skull splitter"},
	4:   {Name: "Rakescar"},
	5:   {Name: "Axe of Fechmar"},
	6:   {Name: "Goreshovel"},
	7:   {Name: "The Chiefthan"},
	8:   {Name: "Brainhew"},
	9:   {Name: "Humongous"},
	10:  {Name: "Torch of Iros"},
	11:  {Name: "Maelstorm"},
	12:  {Name: "Gravenspine"},
	13:  {Name: "Umes Lament"},
	14:  {Name: "Felloak"},
	15:  {Name: "Knell Striker"},
	16:  {Name: "Rusthandle"},
	17:  {Name: "Stormeye"},
	18:  {Name: "Stoutnail"},
	19:  {Name: "Crushflange"},
	20:  {Name: "Bloodrise"},
	21:  {Name: "The Generals Tan Do Li Ga"},
	22:  {Name: "Ironstone"},
	23:  {Name: "Bonesnap"},
	24:  {Name: "Steeldriver"},
	25:  {Name: "Rixot's Keen"},
	26:  {Name: "Blood Crescent"},
	27:  {Name: "Skewer of Krintiz"},
	28:  {Name: "Gleamscythe"},
	29:  {Name: "Azurewrath"},
	30:  {Name: "Griswold's Edge"},
	31:  {Name: "Hellplague"},
	32:  {Name: "Culwens Point"},
	33:  {Name: "Shadowfang"},
	34:  {Name: "Soulflay"},
	35:  {Name: "Kinemils Awl"},
	36:  {Name: "Blacktongue"},
	37:  {Name: "Ripsaw"},
	38:  {Name: "The Patriarch"},
	39:  {Name: "Gull"},
	40:  {Name: "The Diggler"},
	41:  {Name: "The Jade Tan Do"},
	42:  {Name: "Spectral Shard"},
	43:  {Name: "The Dragon Chang"},
	44:  {Name: "Razortine"},
	45:  {Name: "Bloodthief"},
	46:  {Name: "Lance of Yaggai"},
	47:  {Name: "The Tannr Gorerod"},
	48:  {Name: "Dimoaks Hew"},
	49:  {Name: "Steelgoad"},
	50:  {Name: "Soul Harvest"},
	51:  {Name: "The Battlebranch"},
	52:  {Name: "Woestave"},
	53:  {Name: "The Grim Reaper"},
	54:  {Name: "Bane Ash"},
	55:  {Name: "Serpent Lord"},
	56:  {Name: "Spire of Lazarus"},
	57:  {Name: "The Salamander"},
	58:  {Name: "The Iron Jang Bong"},
	59:  {Name: "Pluckeye"},
	60:  {Name: "Witherstring"},
	61:  {Name: "Raven Claw"},
	62:  {Name: "Rogue's Bow"},
	63:  {Name: "Stormstrike"},
	64:  {Name: "Wizendraw"},
	65:  {Name: "Hellclap"},
	66:  {Name: "Blastbark"},
	67:  {Name: "Leadcrow"},
	68:  {Name: "Ichorsting"},
	69:  {Name: "Hellcast"},
	70:  {Name: "Doomslinger"},
	71:  {ItemType: d2s.Cap, Name: "Biggin's Bonnet"},
	72:  {ItemType: d2s.SkullCap, Name: "Tarnhelm"},
	73:  {Name: "Coif of Glory"},
	74:  {Name: "Duskdeep"},
	75:  {Name: "Wormskull"},
	76:  {Name: "Howltusk"},
	77:  {Name: "Undead Crown"},
	78:  {Name: "The Face of Horror"},
	79:  {Name: "Greyform"},
	80:  {Name: "Blinkbat's Form"},
	81:  {Name: "The Centurion"},
	82:  {Name: "Twitchthroe"},
	83:  {Name: "Darkglow"},
	84:  {Name: "Hawkmail"},
	85:  {Name: "Sparking Mail"},
	86:  {Name: "Venom Ward"},
	87:  {Name: "Iceblink"},
	88:  {Name: "Boneflesh"},
	89:  {Name: "Rockfleece"},
	90:  {Name: "Rattlecage"},
	91:  {Name: "Goldskin"},
	92:  {Name: "Victors Silk"},
	93:  {Name: "Heavenly Garb"},
	94:  {Name: "Pelta Lunata"},
	95:  {Name: "Umbral Disk"},
	96:  {Name: "Stormguild"},
	97:  {Name: "Wall of the Eyeless"},
	98:  {Name: "Swordback Hold"},
	99:  {Name: "Steelclash"},
	100: {Name: "Bverrit Keep"},
	101: {Name: "The Ward"},
	102: {Name: "The Hand of Broc"},
	103: {Name: "Bloodfist"},
	104: {Name: "Chance Guards"},
	105: {Name: "Magefist"},
	106: {Name: "Frostburn"},
	107: {Name: "Hotspur"},
	108: {Name: "Gorefoot"},
	109: {Name: "Treads of Cthon"},
	110: {Name: "Goblin Toe"},
	111: {Name: "Tearhaunch"},
	112: {Name: "Lenymo"},
	113: {Name: "Snakecord"},
	114: {Name: "Nightsmoke"},
	115: {Name: "Goldwrap"},
	116: {Name: "Bladebuckle"},
	117: {Name: "Nokozan Relic"},
	118: {Name: "The Eye of Etlich"},
	119: {Name: "The Mahim-Oak Curio"},
	120: {Name: "Nagelring"},
	121: {Name: "Manald Heal"},
	122: {Name: "The Stone of Jordan"},
	123: {Name: "Amulet of the Viper"},
	124: {Name: "Staff of Kings"},
	125: {Name: "Horadric Staff"},
	126: {Name: "Hell Forge Hammer"},
	127: {Name: "Khalim's Flail"},
	128: {Name: "Super Khalim's Flail"},
	129: {Name: "Coldkill"},
	130: {Name: "Butcher's Pupil"},
	131: {Name: "Islestrike"},
	132: {Name: "Pompe's Wrath"},
	133: {Name: "Guardian Naga"},
	134: {Name: "Warlord's Trust"},
	135: {Name: "Spellsteel"},
	136: {Name: "Stormrider"},
	137: {Name: "Boneslayer Blade"},
	138: {Name: "The Minataur"},
	139: {Name: "Suicide Branch"},
	140: {Name: "Carin Shard"},
	141: {Name: "Arm of King Leoric"},
	142: {Name: "Blackhand Key"},
	143: {Name: "Dark Clan Crusher"},
	144: {Name: "Zakarum's Hand"},
	145: {Name: "The Fetid Sprinkler"},
	146: {Name: "Hand of Blessed Light"},
	147: {Name: "Fleshrender"},
	148: {Name: "Sureshrill Frost"},
	149: {Name: "Moonfall"},
	150: {Name: "Baezil's Vortex"},
	151: {Name: "Earthshaker"},
	152: {Name: "Bloodtree Stump"},
	153: {Name: "The Gavel of Pain"},
	154: {Name: "Bloodletter"},
	155: {Name: "Coldsteel Eye"},
	156: {Name: "Hexfire"},
	157: {Name: "Blade of Ali Baba"},
	158: {Name: "Ginther's Rift"},
	159: {Name: "Headstriker"},
	160: {Name: "Plague Bearer"},
	161: {Name: "The Atlantian"},
	162: {Name: "Crainte Vomir"},
	163: {Name: "Bing Sz Wang"},
	164: {Name: "The Vile Husk"},
	165: {Name: "Cloudcrack"},
	166: {Name: "Todesfaelle Flamme"},
	167: {Name: "Swordguard"},
	168: {Name: "Spineripper"},
	169: {Name: "Heart Carver"},
	170: {Name: "Blackbog's Sharp"},
	171: {Name: "Stormspike"},
	172: {Name: "The Impaler"},
	173: {Name: "Kelpie Snare"},
	174: {Name: "Soulfeast Tine"},
	175: {Name: "Hone Sundan"},
	176: {Name: "Spire of Honor"},
	177: {Name: "The Meat Scraper"},
	178: {Name: "Blackleach Blade"},
	179: {Name: "Athena's Wrath"},
	180: {Name: "Pierre Tombale Couant"},
	181: {Name: "Husoldal Evo"},
	182: {Name: "Grim's Burning Dead"},
	183: {Name: "Razorswitch"},
	184: {Name: "Ribcracker"},
	185: {Name: "Chromatic Ire"},
	186: {Name: "Warpspear"},
	187: {Name: "Skullcollector"},
	188: {Name: "Skystrike"},
	189: {Name: "Riphook"},
	190: {Name: "Kuko Shakaku"},
	191: {Name: "Endlesshail"},
	192: {Name: "Whichwild String"},
	193: {Name: "Cliffkiller"},
	194: {Name: "Magewrath"},
	195: {Name: "Godstrike Arch"},
	196: {Name: "Langer Briser"},
	197: {Name: "Pus Spiter"},
	198: {Name: "Buriza-Do Kyanon"},
	199: {Name: "Demon Machine"},
	200: {Name: "Armor (Unknown)"},
	201: {ItemType: d2s.WarHat, Name: "Peasant Crown"},
	202: {ItemType: d2s.Sallet, Name: "Rockstopper"},
	203: {Name: "Stealskull"},
	204: {Name: "Darksight Helm"},
	205: {Name: "Valkyrie Wing"},
	206: {Name: "Crown of Thieves"},
	207: {Name: "Blckhorn's Face"},
	208: {Name: "Vampire Gaze"},
	209: {Name: "The Spirit Shroud"},
	210: {Name: "Skin of the Vipermagi"},
	211: {Name: "Skin of the Flayed One"},
	212: {Name: "Ironpelt"},
	213: {Name: "Spiritforge"},
	214: {Name: "Crow Caw"},
	215: {Name: "Shaftstop"},
	216: {Name: "Duriel's Shell"},
	217: {Name: "Skullder's Ire"},
	218: {Name: "Guardian Angel"},
	219: {Name: "Toothrow"},
	220: {Name: "Atma's Wail"},
	221: {Name: "Black Hades"},
	222: {Name: "Corpsemourn"},
	223: {Name: "Que-Hegan's Wisdom"},
	224: {Name: "Visceratuant"},
	225: {Name: "Mosers Blessed Circle"},
	226: {Name: "Stormchaser"},
	227: {Name: "Tiamat's Rebuke"},
	228: {Name: "Gerke's Sanctuary"},
	229: {Name: "Radimant's Sphere"},
	230: {Name: "Lidless Wall"},
	231: {Name: "Lance Guard"},
	232: {Name: "Venom Grip"},
	233: {Name: "Gravepalm"},
	234: {Name: "Ghoulhide"},
	235: {Name: "Lavagout"},
	236: {Name: "Hellmouth"},
	237: {Name: "Infernostride"},
	238: {Name: "Waterwalk"},
	239: {Name: "Silkweave"},
	240: {Name: "Wartraveler"},
	241: {Name: "Gorerider"},
	242: {Name: "String of Ears"},
	243: {Name: "Razortail"},
	244: {Name: "Gloomstrap"},
	245: {Name: "Snowclash"},
	246: {Name: "Thundergod's Vigor"},
	247: {Name: "Elite unique"},
	248: {ItemType: d2s.Shako, Name: "Harlequin Crest"},
	249: {Name: "Veil of Steel"},
	250: {Name: "The Gladiator's Bane"},
	251: {Name: "Arkaine's Valor"},
	252: {Name: "Blackoak Shield"},
	253: {Name: "Stormshield"},
	254: {Name: "Hellslayer"},
	255: {Name: "Messerschmidt's Reaver"},
	256: {Name: "Baranar's Star"},
	257: {Name: "Schaefer's Hammer"},
	258: {Name: "The Cranium Basher"},
	259: {Name: "Lightsabre"},
	260: {Name: "Doombringer"},
	261: {Name: "The Grandfather"},
	262: {Name: "Wizardspike"},
	263: {Name: "Constricting Ring"},
	264: {Name: "Stormspire"},
	265: {Name: "Eaglehorn"},
	266: {Name: "Windforce"},
	267: {Name: "Ring"},
	268: {Name: "Bul Katho's Wedding Band"},
	269: {Name: "The Cat's Eye"},
	270: {Name: "The Rising Sun"},
	271: {Name: "Crescent Moon"},
	272: {Name: "Mara's Kaleidoscope"},
	273: {Name: "Atma's Scarab"},
	274: {Name: "Dwarf Star"},
	275: {Name: "Raven Frost"},
	276: {Name: "Highlord's Wrath"},
	277: {Name: "Saracen's Chance"},
	278: {Name: "Class specific"},
	279: {Name: "Arreat's Face"},
	280: {Name: "Homunculus"},
	281: {Name: "Titan's Revenge"},
	282: {Name: "Lycander's Aim"},
	283: {Name: "Lycander's Flank"},
	284: {Name: "The Oculus"},
	285: {Name: "Herald of Zakarum"},
	286: {Name: "Bartuc's Cut-Throat"},
	287: {Name: "Jalal's Mane"},
	288: {Name: "The Scalper"},
	289: {Name: "Bloodmoon"},
	290: {Name: "Djinnslayer"},
	291: {Name: "Deathbit"},
	292: {Name: "Warshrike"},
	293: {Name: "Gutsiphon"},
	294: {Name: "Razoredge"},
	295: {Name: "Gore Ripper"},
	296: {Name: "Demon Limb"},
	297: {Name: "Steel Shade"},
	298: {Name: "Tomb Reaver"},
	299: {Name: "Death's Web"},
	300: {Name: "Nature's Peace"},
	301: {Name: "Azurewrath"},
	302: {Name: "Seraph's Hymn"},
	303: {Name: "Zakarum's Salvation"},
	304: {Name: "Fleshripper"},
	305: {Name: "Odium"},
	306: {Name: "Horizon's Tornado"},
	307: {Name: "Stone Crusher"},
	308: {Name: "Jade Talon"},
	309: {Name: "Shadow Dancer"},
	310: {Name: "Cerebus' Bite"},
	311: {Name: "Tyrael's Might"},
	312: {Name: "Soul Drainer"},
	313: {Name: "Rune Master"},
	314: {Name: "Death Cleaver"},
	315: {Name: "Executioner's Justice"},
	316: {Name: "Stoneraven"},
	317: {Name: "Leviathan"},
	318: {Name: "Larzuk's Champion"},
	319: {Name: "Wisp Projector"},
	320: {Name: "Gargoyle's Bite"},
	321: {Name: "Lacerator"},
	322: {Name: "Mang Song's Lesson"},
	323: {Name: "Viperfork"},
	324: {Name: "Ethereal Edge"},
	325: {Name: "Demonhorn's Edge"},
	326: {Name: "The Reaper's Toll"},
	327: {Name: "Spiritkeeper"},
	328: {Name: "Hellrack"},
	329: {Name: "Alma Negra"},
	330: {Name: "Darkforge Spawn"},
	331: {Name: "Widowmaker"},
	332: {Name: "Bloodraven's Charge"},
	333: {Name: "Ghostflame"},
	334: {Name: "Shadowkiller"},
	335: {Name: "Gimmershred"},
	336: {Name: "Griffon's Eye"},
	337: {Name: "Windhammer"},
	338: {Name: "Thunderstroke"},
	339: {Name: "Giant Maimer"},
	340: {Name: "Demon's Arch"},
	341: {Name: "Boneflame"},
	342: {Name: "Steelpillar"},
	343: {Name: "Nightwing's Veil"},
	344: {Name: "Crown of Ages"},
	345: {Name: "Andariel's Visage"},
	346: {Name: "Darkfear"},
	347: {Name: "Dragonscale"},
	348: {Name: "Steel Carapice"},
	349: {Name: "Medusa's Gaze"},
	350: {Name: "Ravenlore"},
	351: {Name: "Boneshade"},
	352: {Name: "Nethercrow"},
	353: {Name: "Flamebellow"},
	354: {Name: "Fathom"},
	355: {Name: "Wolfhowl"},
	356: {Name: "Spirit Ward"},
	357: {Name: "Kira's Guardian"},
	358: {Name: "Ormus Robes"},
	359: {Name: "Gheed's Fortune"},
	360: {Name: "Stormlash"},
	361: {Name: "Halaberd's Reign"},
	362: {Name: "Warriv's Warder"},
	363: {Name: "Spike Thorn"},
	364: {Name: "Dracul's Grasp"},
	365: {Name: "Frostwind"},
	366: {Name: "Templar's Might"},
	367: {Name: "Eschuta's Temper"},
	368: {Name: "Firelizard's Talons"},
	369: {Name: "Sandstorm Trek"},
	370: {Name: "Marrowwalk"},
	371: {Name: "Heaven's Light"},
	372: {Name: "Merman's Speed"},
	373: {Name: "Arachnid Mesh"},
	374: {Name: "Nosferatu's Coil"},
	375: {Name: "Metalgrid"},
	376: {Name: "Verdugo's Hearty Cord"},
	377: {Name: "Sigurd's Staunch"},
	378: {Name: "Carrion Wind"},
	379: {Name: "Giantskull"},
	380: {Name: "Ironward"},
	381: {ItemType: d2s.SmallCharm, Name: "Annihilus"},
	382: {Name: "Arioc's Needle"},
	383: {Name: "Cranebeak"},
	384: {Name: "Nord's Tenderizer"},
	385: {Name: "Earthshifter"},
	386: {Name: "Wraithflight"},
	387: {Name: "Bonehew"},
	388: {Name: "Ondal's Wisdom"},
	389: {Name: "The Reedeemer"},
	390: {Name: "Headhunter's Glory"},
	391: {Name: "Steelrend"},
	392: {ItemType: d2s.Jewel, Name: "Rainbow Facet"},
	393: {ItemType: d2s.Jewel, Name: "Rainbow Facet"},
	394: {ItemType: d2s.Jewel, Name: "Rainbow Facet"},
	395: {ItemType: d2s.Jewel, Name: "Rainbow Facet"},
	396: {ItemType: d2s.Jewel, Name: "Rainbow Facet"},
	397: {ItemType: d2s.Jewel, Name: "Rainbow Facet"},
	398: {ItemType: d2s.Jewel, Name: "Rainbow Facet"},
	399: {ItemType: d2s.Jewel, Name: "Rainbow Facet"},
	400: {ItemType: d2s.LargeCharm, Name: "Hellfire Torch"},
}

var SetItems = map[uint64]SetItem{
	0:   {ItemType: d2s.LargeShield, SetName: "Civerb's Vestments", Name: "Civerb's Ward"},
	1:   {ItemType: d2s.Amulet, SetName: "Civerb's Vestments", Name: "Civerb's Icon"},
	2:   {ItemType: d2s.GrandScepter, SetName: "Civerb's Vestments", Name: "Civerb's Cudgel"},
	3:   {ItemType: d2s.ChainBoots, SetName: "Hsarus' Defense", Name: "Hsaru's Iron Heel"},
	4:   {ItemType: d2s.Buckler, SetName: "Hsarus' Defense", Name: "Hsaru's Iron Fist"},
	5:   {ItemType: d2s.Belt, SetName: "Hsarus' Defense", Name: "Hsaru's Iron Stay"},
	6:   {ItemType: d2s.LongSword, SetName: "Cleglaw's Brace", Name: "Cleglaw's Tooth"},
	7:   {ItemType: d2s.SmallShield, SetName: "Cleglaw's Brace", Name: "Cleglaw's Claw"},
	8:   {ItemType: d2s.Bracers, SetName: "Cleglaw's Brace", Name: "Cleglaw's Pincers"},
	9:   {ItemType: d2s.Amulet, SetName: "Iratha's Finery", Name: "Iratha's Collar"},
	10:  {ItemType: d2s.LightGauntlets, SetName: "Iratha's Finery", Name: "Iratha's Cuff"},
	11:  {ItemType: d2s.Crown, SetName: "Iratha's Finery", Name: "Iratha's Coil"},
	12:  {ItemType: d2s.HeavyBelt, SetName: "Iratha's Finery", Name: "Iratha's Cord"},
	13:  {ItemType: d2s.BroadSword, SetName: "Isenhart's Armory", Name: "Isenhart's Lightbrand"},
	14:  {ItemType: d2s.GothicShield, SetName: "Isenhart's Armory", Name: "Isenhart's Parry"},
	15:  {ItemType: d2s.BreastPlate, SetName: "Isenhart's Armory", Name: "Isenhart's Case"},
	16:  {ItemType: d2s.FullHelm, SetName: "Isenhart's Armory", Name: "Isenhart's Horns"},
	17:  {ItemType: d2s.LongBattleBow, SetName: "Vidala's Rig", Name: "Vidala's Barb"},
	18:  {ItemType: d2s.LightPlatedBoots, SetName: "Vidala's Rig", Name: "Vidala's Fetlock"},
	19:  {ItemType: d2s.LeatherArmor, SetName: "Vidala's Rig", Name: "Vidala's Ambush"},
	20:  {ItemType: d2s.Amulet, SetName: "Vidala's Rig", Name: "Vidala's Snare"},
	21:  {ItemType: d2s.KiteShield, SetName: "Milabrega's Regalia", Name: "Milabrega's Orb"},
	22:  {ItemType: d2s.AncientArmor, SetName: "Milabrega's Regalia", Name: "Milabrega's Rod"},
	23:  {ItemType: d2s.Crown, SetName: "Milabrega's Regalia", Name: "Milabrega's Diadem"},
	24:  {ItemType: d2s.WarScepter, SetName: "Milabrega's Regalia", Name: "Mialbrega's Robe"},
	25:  {ItemType: d2s.BattleStaff, SetName: "Cathan's Trap", Name: "Cathan's Rule"},
	26:  {ItemType: d2s.ChainMail, SetName: "Cathan's Trap", Name: "Cathan's Mesh"},
	27:  {ItemType: d2s.Mask, SetName: "Cathan's Trap", Name: "Cathan's Visage"},
	28:  {ItemType: d2s.Amulet, SetName: "Cathan's Trap", Name: "Cathan's Sigil"},
	29:  {ItemType: d2s.Ring, SetName: "Cathan's Trap", Name: "Cathan's Seal"},
	30:  {ItemType: d2s.MilitaryPick, SetName: "Tancred's Battlegear", Name: "Tancred's Crowbill"},
	31:  {ItemType: d2s.FullPlateMail, SetName: "Tancred's Battlegear", Name: "Tancred's Spine"},
	32:  {ItemType: d2s.LeatherBoots, SetName: "Tancred's Battlegear", Name: "Tancred's Hobnails"},
	33:  {ItemType: d2s.Amulet, SetName: "Tancred's Battlegear", Name: "Tancred's Weird"},
	34:  {ItemType: d2s.BoneHelm, SetName: "Tancred's Battlegear", Name: "Tancred's Skull"},
	35:  {ItemType: d2s.Gauntlets, SetName: "Sigon's Complete Steel", Name: "Sigon's Gage"},
	36:  {ItemType: d2s.GreatHelm, SetName: "Sigon's Complete Steel", Name: "Sigon's Visor"},
	37:  {ItemType: d2s.GothicPlate, SetName: "Sigon's Complete Steel", Name: "Sigon's Shelter"},
	38:  {ItemType: d2s.Greaves, SetName: "Sigon's Complete Steel", Name: "Sigon's Sabot"},
	39:  {ItemType: d2s.PlatedBelt, SetName: "Sigon's Complete Steel", Name: "Sigon's Wrap"},
	40:  {ItemType: d2s.TowerShield, SetName: "Sigon's Complete Steel", Name: "Sigon's Guard"},
	41:  {ItemType: d2s.Cap, SetName: "Infernal Tools", Name: "Infernal Cranium"},
	42:  {ItemType: d2s.GrimWand, SetName: "Infernal Tools", Name: "Infernal Torch"},
	43:  {ItemType: d2s.HeavyBelt, SetName: "Infernal Tools", Name: "Infernal Sign"},
	44:  {ItemType: d2s.Helm, SetName: "Berserker's Arsenal", Name: "Berserker's Headgear"},
	45:  {ItemType: d2s.SplintMail, SetName: "Berserker's Arsenal", Name: "Berserker's Hauberk"},
	46:  {ItemType: d2s.DoubleAxe, SetName: "Berserker's Arsenal", Name: "Berserker's Hatchet"},
	47:  {ItemType: d2s.LeatherGloves, SetName: "Death's Disguise", Name: "Death's Hand"},
	48:  {ItemType: d2s.Sash, SetName: "Death's Disguise", Name: "Death's Guard"},
	49:  {ItemType: d2s.WarSword, SetName: "Death's Disguise", Name: "Death's Touch"},
	50:  {ItemType: d2s.Sabre, SetName: "Angelic Raiment", Name: "Angelic Sickle"},
	51:  {ItemType: d2s.RingMail, SetName: "Angelic Raiment", Name: "Angelic Mantle"},
	52:  {ItemType: d2s.Ring, SetName: "Angelic Raiment", Name: "Angelic Halo"},
	53:  {ItemType: d2s.Amulet, SetName: "Angelic Raiment", Name: "Angelic Wings"},
	54:  {ItemType: d2s.ShortWarBow, SetName: "Arctic Gear", Name: "Arctic Horn"},
	55:  {ItemType: d2s.QuiltedArmor, SetName: "Arctic Gear", Name: "Arctic Furs"},
	56:  {ItemType: d2s.LightBelt, SetName: "Arctic Gear", Name: "Arctic Binding"},
	57:  {ItemType: d2s.LightGauntlets, SetName: "Arctic Gear", Name: "Arctic Mitts"},
	58:  {ItemType: d2s.Amulet, SetName: "Arcanna's Tricks", Name: "Arcanna's Sign"},
	59:  {ItemType: d2s.WarStaff, SetName: "Arcanna's Tricks", Name: "Arcanna's Deathwand"},
	60:  {ItemType: d2s.SkullCap, SetName: "Arcanna's Tricks", Name: "Arcanna's Head"},
	61:  {ItemType: d2s.LightPlate, SetName: "Arcanna's Tricks", Name: "Arcanna's Flesh"},
	62:  {ItemType: d2s.GrimHelm, SetName: "Natalya's Odium", Name: "Natalya's Totem"},
	63:  {ItemType: d2s.ScissorsSuwayyah, SetName: "Natalya's Odium", Name: "Natalya's Mark"},
	64:  {ItemType: d2s.LoricatedMail, SetName: "Natalya's Odium", Name: "Natalya's Shadow"},
	65:  {ItemType: d2s.MeshBoots, SetName: "Natalya's Odium", Name: "Natalya's Soul"},
	66:  {ItemType: d2s.HuntersGuise, SetName: "Aldur's Watchtower", Name: "Aldur's Stony Gaze"},
	67:  {ItemType: d2s.ShadowPlate, SetName: "Aldur's Watchtower", Name: "Aldur's Deception"},
	68:  {ItemType: d2s.JaggedStar, SetName: "Aldur's Watchtower", Name: "Aldur's Rhythm"},
	69:  {ItemType: d2s.BattleBoots, SetName: "Aldur's Watchtower", Name: "Aldur's Advance"},
	70:  {ItemType: d2s.AvengerGuard, SetName: "Immortal King", Name: "Immortal King's Will"},
	71:  {ItemType: d2s.SacredArmor, SetName: "Immortal King", Name: "Immortal King's Soul Cage"},
	72:  {ItemType: d2s.WarBelt, SetName: "Immortal King", Name: "Immortal King's Detail"},
	73:  {ItemType: d2s.WarGauntlets, SetName: "Immortal King", Name: "Immortal King's Forge"},
	74:  {ItemType: d2s.WarBoots, SetName: "Immortal King", Name: "Immortal King's Pillar"},
	75:  {ItemType: d2s.OgreMaul, SetName: "Immortal King", Name: "Immortal King's Stone Crusher"},
	76:  {ItemType: d2s.MeshBelt, SetName: "Tal Rasha's Wrappings", Name: "Tal Rasha's Fine-Spun Cloth"},
	77:  {ItemType: d2s.Amulet, SetName: "Tal Rasha's Wrappings", Name: "Tal Rasha's Adjudication"},
	78:  {ItemType: d2s.SwirlingCrystal, SetName: "Tal Rasha's Wrappings", Name: "Tal Rasha's Lidless Eye"},
	79:  {ItemType: d2s.LacqueredPlate, SetName: "Tal Rasha's Wrappings", Name: "Tal Rasha's Guardianship"},
	80:  {ItemType: d2s.DeathMask, SetName: "Tal Rasha's Wrappings", Name: "Tal Rasha's Horadric Crest"},
	81:  {ItemType: d2s.Corona, SetName: "Griswold's Legacy", Name: "Griswold's Valor"},
	82:  {ItemType: d2s.OrnatePlate, SetName: "Griswold's Legacy", Name: "Griswold's Heart"},
	83:  {ItemType: d2s.Caduceus, SetName: "Griswold's Legacy", Name: "Griswold's Redemption"},
	84:  {ItemType: d2s.VortexShield, SetName: "Griswold's Legacy", Name: "Griswold's Honor"},
	85:  {ItemType: d2s.BoneVisage, SetName: "Trang-Oul's Avatar", Name: "Trang-Oul's Guise"},
	86:  {ItemType: d2s.ChaosArmor, SetName: "Trang-Oul's Avatar", Name: "Trang-Oul's Scales"},
	87:  {ItemType: d2s.CantorTrophy, SetName: "Trang-Oul's Avatar", Name: "Trang-Oul's Wing"},
	88:  {ItemType: d2s.HeavyBracers, SetName: "Trang-Oul's Avatar", Name: "Trang-Oul's Claws"},
	89:  {ItemType: d2s.TrollBelt, SetName: "Trang-Oul's Avatar", Name: "Trang-Oul's Girth"},
	90:  {ItemType: d2s.Diadem, SetName: "M'avina's Battle Hymn", Name: "M'avina's True Sight"},
	91:  {ItemType: d2s.KrakenShell, SetName: "M'avina's Battle Hymn", Name: "M'avina's Embrace"},
	92:  {ItemType: d2s.BattleGauntlets, SetName: "M'avina's Battle Hymn", Name: "M'avina's Icy Clutch"},
	93:  {ItemType: d2s.SharkskinBelt, SetName: "M'avina's Battle Hymn", Name: "M'avina's Tenet"},
	94:  {ItemType: d2s.GrandMatronBow, SetName: "M'avina's Battle Hymn", Name: "M'avina's Caster"},
	95:  {ItemType: d2s.Amulet, SetName: "The Disciple", Name: "Telling of Beads"},
	96:  {ItemType: d2s.BrambleMitts, SetName: "The Disciple", Name: "Laying of Hands"},
	97:  {ItemType: d2s.DemonhideBoots, SetName: "The Disciple", Name: "Rite of Passage"},
	98:  {ItemType: d2s.DuskShroud, SetName: "The Disciple", Name: "Dark Adherent"},
	99:  {ItemType: d2s.MithrilCoil, SetName: "The Disciple", Name: "Credendum"},
	100: {ItemType: d2s.ReinforcedMace, SetName: "Heaven's Brethren", Name: "Dangoon's Teaching"},
	101: {ItemType: d2s.Ward, SetName: "Heaven's Brethren", Name: "Taebaek's Glory"},
	102: {ItemType: d2s.Cuirass, SetName: "Heaven's Brethren", Name: "Haemosu's Adament"},
	103: {ItemType: d2s.SpiredHelm, SetName: "Heaven's Brethren", Name: "Ondal's Almighty"},
	104: {ItemType: d2s.WingedHelm, SetName: "Orphan's Call", Name: "Guillaume's Face"},
	105: {ItemType: d2s.BattleBelt, SetName: "Orphan's Call", Name: "Wilhelm's Pride"},
	106: {ItemType: d2s.SharkskinGloves, SetName: "Orphan's Call", Name: "Magnus' Skin"},
	107: {ItemType: d2s.RoundShield, SetName: "Orphan's Call", Name: "Whitstan's Guard"},
	108: {ItemType: d2s.GrandCrown, SetName: "Hwanin's Majesty", Name: "Hwanin's Splendor"},
	109: {ItemType: d2s.TigulatedMail, SetName: "Hwanin's Majesty", Name: "Hwanin's Refuge"},
	110: {ItemType: d2s.Belt, SetName: "Hwanin's Majesty", Name: "Hwanin's Blessing"},
	111: {ItemType: d2s.Bill, SetName: "Hwanin's Majesty", Name: "Hwanin's Justice"},
	112: {ItemType: d2s.CrypticSword, SetName: "Sazabi's Grand Tribute", Name: "Sazabi's Cobalt Redeemer"},
	113: {ItemType: d2s.BalrogSkin, SetName: "Sazabi's Grand Tribute", Name: "Sazabi's Ghost Liberator"},
	114: {ItemType: d2s.Basinet, SetName: "Sazabi's Grand Tribute", Name: "Sazabi's Mental Sheath"},
	115: {ItemType: d2s.ColossalBlade, SetName: "Bul-Kathos' Children", Name: "Bul-Katho's Sacred Charge"},
	116: {ItemType: d2s.MythicalSword, SetName: "Bul-Kathos' Children", Name: "Bul-Katho's Tribal Guardian"},
	117: {ItemType: d2s.WarHat, SetName: "Cow King's Leathers", Name: "Cow King's Horns"},
	118: {ItemType: d2s.StuddedLeather, SetName: "Cow King's Leathers", Name: "Cow King's Hide"},
	119: {ItemType: d2s.HeavyBoots, SetName: "Cow King's Leathers", Name: "Cow King's Hooves"},
	120: {ItemType: d2s.ElderStaff, SetName: "Naj's Ancient Vestige", Name: "Naj's Puzzler"},
	121: {ItemType: d2s.HellforgedPlate, SetName: "Naj's Ancient Vestige", Name: "Naj's Light Plate"},
	122: {ItemType: d2s.Circlet, SetName: "Naj's Ancient Vestige", Name: "Naj's Circlet"},
	123: {ItemType: d2s.Cap, SetName: "Sander's Folly", Name: "Sander's Paragon"},
	124: {ItemType: d2s.HeavyBoots, SetName: "Sander's Folly", Name: "Sander's Riprap"},
	125: {ItemType: d2s.HeavyGloves, SetName: "Sander's Folly", Name: "Sander's Taboo"},
	126: {ItemType: d2s.BoneWand, SetName: "Sander's Folly", Name: "Sander's Superstition"},
}
