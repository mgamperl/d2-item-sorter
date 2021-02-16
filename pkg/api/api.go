package api

import (
	"d2-item-sorter/pkg/data"
	"d2-item-sorter/pkg/domain"
	"d2-item-sorter/pkg/internal"
	"d2-item-sorter/pkg/internal/utils"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/nokka/d2s"
)

var statistics = map[string]int{}

// GetGroups returns the groups parsed from the porovided config file
func GetGroups(groupConfigFile string) []internal.ItemGroup {
	groups := []internal.ItemGroup{}
	if len(groups) <= 0 {
		groups = internal.LoadGroupConfigFile(groupConfigFile)
	}
	fmt.Printf("%d groups loaded\n", len(groups))
	return groups
}

func GetSaveDir() (string, error) {
	saveDir := utils.GetD2SavePath()
	fmt.Printf("D2 save directory: %s\n", saveDir)

	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		panic("directory " + saveDir + " does not exist")
		//(fmt.Printf("%s does not exist", saveDir)
		//return _, err
	}

	saveDir, err := filepath.Abs(saveDir)

	return saveDir, err
}

func GetInstallDir() (string, error) {
	installDir := utils.GetD2InstallPath()
	fmt.Printf("D2 install directory: %s\n", installDir)

	if _, err := os.Stat(installDir); os.IsNotExist(err) {
		panic("directory " + installDir + " does not exist")
		//(fmt.Printf("%s does not exist", saveDir)
		//return _, err
	}

	installDir, err := filepath.Abs(installDir)

	return installDir, err
}

func GetGrailStatusList() []domain.GrailStatusItem {
	itemMap := map[string]*domain.GrailStatusItem{}
	items := []domain.GrailStatusItem{}

	saveDir, err := GetSaveDir()

	statistics = map[string]int{}

	if err != nil {
		panic("can not find Diablo2 Directory")
	}

	sharedStash := internal.ParseSharedStash(saveDir)
	characters, _ := internal.ParseCharacters(saveDir)

	for idx, value := range data.UniqueItems {
		if string(value.ItemType) != "" {
			itemMap[fmt.Sprintf("u%d", idx)] = &domain.GrailStatusItem{
				Name:           value.Name,
				Order:          int(idx),
				ItemQuality:    string(data.ItemPropertyMap[value.ItemType].ItemQuality),
				Category:       string(data.ItemPropertyMap[value.ItemType].ItemCategory),
				SubCategory:    string(data.ItemPropertyMap[value.ItemType].ItemSubSubCategory),
				SubSubCategory: string(data.ItemPropertyMap[value.ItemType].ItemSubSubCategory),
				Type:           string(value.ItemType),
				TypeName:       utils.GetTypeName(value.ItemType),
				IsUnique:       true,
				Count:          0,
				ImageID:        fmt.Sprintf("u%d", idx)}
		}
	}

	for idx, value := range data.SetItems {
		if string(value.ItemType) != "" {
			itemMap[fmt.Sprintf("s%d", idx)] = &domain.GrailStatusItem{
				Name:           value.Name,
				Order:          int(idx),
				ItemQuality:    string(data.ItemPropertyMap[value.ItemType].ItemQuality),
				Category:       string(data.ItemPropertyMap[value.ItemType].ItemCategory),
				SubCategory:    string(data.ItemPropertyMap[value.ItemType].ItemSubSubCategory),
				SubSubCategory: string(data.ItemPropertyMap[value.ItemType].ItemSubSubCategory),
				Type:           string(value.ItemType),
				TypeName:       utils.GetTypeName(value.ItemType),
				IsSet:          true,
				SetName:        value.SetName,
				Count:          0,
				ImageID:        fmt.Sprintf("s%d", idx)}
		}
	}

	for i := 1; i <= 33; i++ {
		runeCode := fmt.Sprintf("r%02d", i)
		itemMap[runeCode] = &domain.GrailStatusItem{
			Name:           d2s.MiscCodes[d2s.ItemType(runeCode)],
			Order:          int(i),
			Type:           runeCode,
			ItemQuality:    string(data.ItemPropertyMap[d2s.ItemType(runeCode)].ItemQuality),
			Category:       string(data.ItemPropertyMap[d2s.ItemType(runeCode)].ItemCategory),
			SubCategory:    string(data.ItemPropertyMap[d2s.ItemType(runeCode)].ItemSubSubCategory),
			SubSubCategory: string(data.ItemPropertyMap[d2s.ItemType(runeCode)].ItemSubSubCategory),
			TypeName:       utils.GetTypeName(d2s.ItemType(runeCode)),
			IsRune:         true,
			Count:          0,
			ImageID:        runeCode}
	}

	for _, char := range characters {
		//fmt.Printf("[%s] %s (Level %d): %d Items\n", char.D2sCharacter.Header.Class, char.D2sCharacter.Header.Name, char.D2sCharacter.Header.Level, len(char.GetAllItems()))

		//fmt.Printf("\t%000d items in personal stash\n", len(char.PersonalStash.GetAllItems()))
		for _, item := range char.PersonalStash.GetAllItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			if _, ok := itemMap[item.ID]; ok == true {
				itemMap[item.ID].Count++
			}
		}

		fmt.Printf("\t%000d items in inventory\n", len(char.GetInventoryItems()))
		for _, item := range char.GetInventoryItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			if _, ok := itemMap[item.ID]; ok == true {
				itemMap[item.ID].Count++
			}
		}
		fmt.Printf("\t%000d items equipped on character\n", len(char.GetEquippedItems()))
		for _, item := range char.GetEquippedItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			if _, ok := itemMap[item.ID]; ok == true {
				itemMap[item.ID].Count++
			}
		}
		fmt.Printf("\t%000d items equipped on merc\n", len(char.GetMercItems()))
		for _, item := range char.GetMercItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			if _, ok := itemMap[item.ID]; ok == true {
				itemMap[item.ID].Count++
			}
		}
	}

	fmt.Printf("%d items in shared stash\n", len(sharedStash.GetAllItems()))
	for _, item := range sharedStash.GetAllItems() {
		//fmt.Printf("\t\t %s\n", item.GetName())

		if _, ok := itemMap[item.ID]; ok == true {
			itemMap[item.ID].Count++
		}
	}

	for _, v := range itemMap {
		items = append(items, *v)
	}

	sort.Slice(items, func(i int, j int) bool {
		return items[i].Order < items[j].Order
	})

	return items
}

func SortItems(name string, groups []internal.ItemGroup) {
	saveDir, err := GetSaveDir()

	if err != nil {
		panic("can not find Diablo2 Directory")
	}

	if name == "_shared" {
		fmt.Printf("sorting items for Shared Stash")

		sharedStash := internal.ParseSharedStash(saveDir)
		sharedStash.Reload()
		sharedStash.SortItems(groups)
		sharedStash.Save("")

		//TOdO: Enable Loot Filter Again
		//internal.WriteBHConfig(utils.GetD2InstallPath(), sharedStash.GetAllItems())

	} else {

		characters, _ := internal.ParseCharacters(saveDir)

		if _, okFrom := characters[name]; okFrom == false && name != "_shared" {
			fmt.Printf("%s is not a valid character name or does not exist\n", name)
			return
		}

		fmt.Printf("sorting items for %s\n", name)

		characters[name].Reload()
		characters[name].SortPersonalStash(groups)
		characters[name].Save("")
	}

	return
}

func TransferItems(nameFrom string, nameTo string, groups []internal.ItemGroup) {
	saveDir, err := GetSaveDir()

	if err != nil {
		panic("can not find Diablo2 Directory")
	}

	sharedStash := internal.ParseSharedStash(saveDir)
	characters, _ := internal.ParseCharacters(saveDir)

	if _, okFrom := characters[nameFrom]; okFrom == false && nameFrom != "_shared" {
		fmt.Printf("%s is not a valid character name or does not exist\n", nameFrom)
		return
	}

	if _, okTo := characters[nameTo]; okTo == false && nameTo != "_shared" {
		fmt.Printf("%s is not a valid character name or does not exist\n", nameTo)
		return
	}

	if nameFrom == "_shared" {
		fmt.Printf("transfering items from Shared Stash to %s\n", nameTo)
		sharedStash.Reload()
		characters[nameTo].Reload()

		sharedStash.TransferPagesTo(&characters[nameTo].PersonalStash)

		sharedStash.SortItems(groups)
		characters[nameTo].SortPersonalStash(groups)

		sharedStash.Save("")
		characters[nameTo].Save("")

	} else if nameTo == "_shared" {
		fmt.Printf("transfering items from %s to Shared Stash\n", nameFrom)
		sharedStash.Reload()
		characters[nameFrom].Reload()

		characters[nameFrom].PersonalStash.TransferPagesTo(&sharedStash)

		sharedStash.SortItems(groups)
		characters[nameFrom].SortPersonalStash(groups)

		sharedStash.Save("")
		characters[nameFrom].Save("")

	} else {
		fmt.Printf("transfering items from %s to %s\n", nameFrom, nameTo)
		characters[nameFrom].Reload()
		characters[nameTo].Reload()

		characters[nameFrom].PersonalStash.TransferPagesTo(&characters[nameTo].PersonalStash)

		characters[nameTo].SortPersonalStash(groups)
		characters[nameFrom].SortPersonalStash(groups)

		characters[nameTo].Save("")
		characters[nameFrom].Save("")
	}

	return
}

func addStatistic(item internal.SortableItem) {
	if item.IsUnique || item.IsSetItem {
		statistics[item.GetName()]++
	}
}

func PrintItemInfo() {
	saveDir, err := GetSaveDir()

	statistics = map[string]int{}

	if err != nil {
		panic("can not find Diablo2 Directory")
	}

	sharedStash := internal.ParseSharedStash(saveDir)
	characters, _ := internal.ParseCharacters(saveDir)
	fmt.Printf("%d characters found\n", len(characters))

	for _, name := range d2s.UniqueNames {
		fmt.Printf("%s  \n", name)
	}

	for _, char := range characters {
		//fmt.Printf("[%s] %s (Level %d): %d Items\n", char.D2sCharacter.Header.Class, char.D2sCharacter.Header.Name, char.D2sCharacter.Header.Level, len(char.GetAllItems()))

		//fmt.Printf("\t%000d items in personal stash\n", len(char.PersonalStash.GetAllItems()))
		for _, item := range char.PersonalStash.GetAllItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			addStatistic(item)
		}
		fmt.Printf("\t%000d items in inventory\n", len(char.GetInventoryItems()))
		for _, item := range char.GetInventoryItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			addStatistic(item)
		}
		fmt.Printf("\t%000d items equipped on character\n", len(char.GetEquippedItems()))
		for _, item := range char.GetEquippedItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			addStatistic(item)
		}
		fmt.Printf("\t%000d items equipped on merc\n", len(char.GetMercItems()))
		for _, item := range char.GetMercItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			addStatistic(item)
		}
	}

	fmt.Printf("%d items in shared stash\n", len(sharedStash.GetAllItems()))
	for _, item := range sharedStash.GetAllItems() {
		//fmt.Printf("\t\t %s\n", item.GetName())
		addStatistic(item)
	}

	for key, value := range statistics {
		fmt.Printf("\"%s\";%d\n", key, value)
	}
}

func StartWatcher(groups []internal.ItemGroup) {
	saveDir, err := GetSaveDir()

	if err != nil {
		panic("can not find Diablo2 Directory")
	}

	fmt.Printf("Starting Run Recorder\n")
	sharedStash := internal.ParseSharedStash(saveDir)
	internal.StartRunRecorder(&sharedStash, groups)

}
