package data

import "github.com/nokka/d2s"

var GemSortMap = map[ItemSubCategoryName]int{
	Perfect:  1,
	Flawless: 2,
	Normal:   3,
	Flawed:   4,
	Chipped:  5,
}

type ItemPropertyData struct {
	ItemQuality        ItemQuality
	ItemCategory       ItemCategoryName
	ItemSubCategory    ItemSubCategoryName
	ItemSubSubCategory ItemSubSubCategoryName
	SizeX              int
	SizeY              int
}

var ItemPropertyMap = map[d2s.ItemType]*ItemPropertyData{
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
