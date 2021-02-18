package api

import (
	"d2-item-sorter/pkg/config"
	"d2-item-sorter/pkg/data"
	"d2-item-sorter/pkg/domain"
	"d2-item-sorter/pkg/internal/reader"
	"d2-item-sorter/pkg/internal/runRecorder"
	"d2-item-sorter/pkg/internal/sorter"
	"d2-item-sorter/pkg/internal/utils"
	"d2-item-sorter/pkg/internal/writer"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/nokka/d2s"
)

var statistics = map[string]int{}

// GetGroups returns the groups parsed from the porovided config file
func GetGroups(groupConfigFile string) []domain.ItemGroup {
	groups := []domain.ItemGroup{}
	if len(groups) <= 0 {
		groups = reader.LoadGroupConfigFile(groupConfigFile)
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

	sharedStash := ReadSharedStash(saveDir)
	characters, _ := reader.ReadAllCharactersFromPath(saveDir)

	for idx, value := range data.UniqueItems {
		if string(value.ItemType) != "" {
			itemMap[fmt.Sprintf("u%d", idx)] = &domain.GrailStatusItem{
				Name:           value.Name,
				Order:          int(idx),
				ItemQuality:    string(data.ItemTypeMap[value.ItemType].ItemQuality),
				Category:       string(data.ItemTypeMap[value.ItemType].ItemCategory),
				SubCategory:    string(data.ItemTypeMap[value.ItemType].ItemSubSubCategory),
				SubSubCategory: string(data.ItemTypeMap[value.ItemType].ItemSubSubCategory),
				Type:           string(value.ItemType),
				TypeName:       utils.GetTypeName(value.ItemType),
				IsUnique:       true,
				Count:          0,
				ImageID:        fmt.Sprintf("u%d", idx),
				StoredAt:       []string{}}
		}
	}

	for idx, value := range data.SetItems {
		if string(value.ItemType) != "" {
			itemMap[fmt.Sprintf("s%d", idx)] = &domain.GrailStatusItem{
				Name:           value.Name,
				Order:          int(idx),
				ItemQuality:    string(data.ItemTypeMap[value.ItemType].ItemQuality),
				Category:       string(data.ItemTypeMap[value.ItemType].ItemCategory),
				SubCategory:    string(data.ItemTypeMap[value.ItemType].ItemSubSubCategory),
				SubSubCategory: string(data.ItemTypeMap[value.ItemType].ItemSubSubCategory),
				Type:           string(value.ItemType),
				TypeName:       utils.GetTypeName(value.ItemType),
				IsSet:          true,
				SetName:        value.SetName,
				Count:          0,
				ImageID:        fmt.Sprintf("s%d", idx),
				StoredAt:       []string{}}
		}
	}

	for i := 1; i <= 33; i++ {
		runeCode := fmt.Sprintf("r%02d", i)
		itemMap[runeCode] = &domain.GrailStatusItem{
			Name:           d2s.MiscCodes[d2s.ItemType(runeCode)],
			Order:          int(i),
			Type:           runeCode,
			ItemQuality:    string(data.ItemTypeMap[d2s.ItemType(runeCode)].ItemQuality),
			Category:       string(data.ItemTypeMap[d2s.ItemType(runeCode)].ItemCategory),
			SubCategory:    string(data.ItemTypeMap[d2s.ItemType(runeCode)].ItemSubSubCategory),
			SubSubCategory: string(data.ItemTypeMap[d2s.ItemType(runeCode)].ItemSubSubCategory),
			TypeName:       utils.GetTypeName(d2s.ItemType(runeCode)),
			IsRune:         true,
			Count:          0,
			ImageID:        runeCode,
			StoredAt:       []string{}}
	}

	for _, char := range characters {
		//fmt.Printf("[%s] %s (Level %d): %d Items\n", char.D2sCharacter.Header.Class, char.D2sCharacter.Header.Name, char.D2sCharacter.Header.Level, len(char.GetAllItems()))

		//fmt.Printf("\t%000d items in personal stash\n", len(char.PersonalStash.GetAllItems()))
		for _, item := range char.PersonalStash.GetAllItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			if _, ok := itemMap[item.ID]; ok == true {
				//itemMap[item.ID].Count++
				addToGrailStatus(itemMap, item, fmt.Sprintf("%s: %s", char.Name, "Personal Stash"))
			}
		}

		fmt.Printf("\t%000d items in inventory\n", len(char.GetInventoryItems()))
		for _, item := range char.GetInventoryItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			if _, ok := itemMap[item.ID]; ok == true {
				//itemMap[item.ID].Count++
				addToGrailStatus(itemMap, item, fmt.Sprintf("%s: %s", char.Name, "Inventory"))
			}
		}
		fmt.Printf("\t%000d items equipped on character\n", len(char.GetEquippedItems()))
		for _, item := range char.GetEquippedItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			if _, ok := itemMap[item.ID]; ok == true {
				//itemMap[item.ID].Count++
				addToGrailStatus(itemMap, item, fmt.Sprintf("%s: %s", char.Name, "Equipped"))
			}
		}
		fmt.Printf("\t%000d items equipped on merc\n", len(char.GetMercItems()))
		for _, item := range char.GetMercItems() {
			//fmt.Printf("\t\t %s\n", item.GetName())
			if _, ok := itemMap[item.ID]; ok == true {
				//itemMap[item.ID].Count++
				addToGrailStatus(itemMap, item, fmt.Sprintf("%s: %s", char.Name, "Mercenary"))
			}
		}
	}

	fmt.Printf("%d items in shared stash\n", len(sharedStash.GetAllItems()))
	for _, item := range sharedStash.GetAllItems() {
		//fmt.Printf("\t\t %s\n", item.GetName())

		if _, ok := itemMap[item.ID]; ok == true {
			//itemMap[item.ID].Count++
			addToGrailStatus(itemMap, item, fmt.Sprintf("%s", "Shared Stash"))
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

func addToGrailStatus(m map[string]*domain.GrailStatusItem, item domain.Item, source string) {
	if _, ok := m[item.ID]; ok == true {
		m[item.ID].Count++
		if !Contains(m[item.ID].StoredAt, source) {
			m[item.ID].StoredAt = append(m[item.ID].StoredAt, source)
		}
	}
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func SortItems(name string, groups []domain.ItemGroup) {
	saveDir, err := GetSaveDir()

	if err != nil {
		panic("can not find Diablo2 Directory")
	}

	if name == "_shared" {
		fmt.Printf("sorting items for Shared Stash")

		sharedStash := ReadSharedStash(saveDir)
		//sharedStash.Reload()
		sharedStash.StashPages = sorter.SortItems(sharedStash, groups)
		writer.WriteStashToFile(sharedStash, "")

		//TOdO: Enable Loot Filter Again
		//internal.WriteBHConfig(utils.GetD2InstallPath(), sharedStash.GetAllItems())

	} else {

		characters, _ := reader.ReadAllCharactersFromPath(saveDir)

		if _, okFrom := characters[name]; okFrom == false && name != "_shared" {
			fmt.Printf("%s is not a valid character name or does not exist\n", name)
			return
		}

		fmt.Printf("sorting items for %s\n", name)

		characters[name] = reader.ReloadAndReturnCharacter(characters[name])
		characters[name].PersonalStash.StashPages = sorter.SortItems(*characters[name].PersonalStash, groups)
		writer.WriteCharacterToFile(characters[name], "")
	}

	return
}

func TransferItems(nameFrom string, nameTo string, groups []domain.ItemGroup) {
	saveDir, err := GetSaveDir()

	if err != nil {
		panic("can not find Diablo2 Directory")
	}

	sharedStash := ReadSharedStash(saveDir)
	characters, _ := reader.ReadAllCharactersFromPath(saveDir)

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
		sharedStash = reader.ReloadAndReturnStash(sharedStash)
		characters[nameTo] = reader.ReloadAndReturnCharacter(characters[nameTo])

		sorter.TransferPagesTo(&sharedStash, characters[nameTo].PersonalStash)

		sharedStash.StashPages = sorter.SortItems(sharedStash, groups)
		characters[nameTo].PersonalStash.StashPages = sorter.SortItems(*characters[nameTo].PersonalStash, groups)

		writer.WriteStashToFile(sharedStash, "")
		writer.WriteCharacterToFile(characters[nameTo], "")

	} else if nameTo == "_shared" {
		fmt.Printf("transfering items from %s to Shared Stash\n", nameFrom)
		sharedStash = reader.ReloadAndReturnStash(sharedStash)
		characters[nameFrom] = reader.ReloadAndReturnCharacter(characters[nameFrom])

		sorter.TransferPagesTo(characters[nameFrom].PersonalStash, &sharedStash)

		sharedStash.StashPages = sorter.SortItems(sharedStash, groups)
		characters[nameFrom].PersonalStash.StashPages = sorter.SortItems(*characters[nameFrom].PersonalStash, groups)

		writer.WriteStashToFile(sharedStash, "")
		writer.WriteCharacterToFile(characters[nameFrom], "")

	} else {
		fmt.Printf("transfering items from %s to %s\n", nameFrom, nameTo)
		characters[nameFrom] = reader.ReloadAndReturnCharacter(characters[nameFrom])
		characters[nameTo] = reader.ReloadAndReturnCharacter(characters[nameTo])

		sorter.TransferPagesTo(characters[nameFrom].PersonalStash, characters[nameTo].PersonalStash)

		characters[nameTo].PersonalStash.StashPages = sorter.SortItems(*characters[nameTo].PersonalStash, groups)
		characters[nameFrom].PersonalStash.StashPages = sorter.SortItems(*characters[nameFrom].PersonalStash, groups)

		writer.WriteCharacterToFile(characters[nameTo], "")
		writer.WriteCharacterToFile(characters[nameFrom], "")
	}

	return
}

func addStatistic(item domain.Item) {
	if item.Quality == domain.ItemQualityUnique || item.Quality == domain.ItemQualitySet {
		statistics[item.ID]++
	}
}

func PrintItemInfo() {
	saveDir, err := GetSaveDir()

	statistics = map[string]int{}

	if err != nil {
		panic("can not find Diablo2 Directory")
	}

	sharedStash := ReadSharedStash(saveDir)
	characters, _ := reader.ReadAllCharactersFromPath(saveDir)
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

func StartWatcher(groups []domain.ItemGroup) {
	saveDir, err := GetSaveDir()

	if err != nil {
		panic("can not find Diablo2 Directory")
	}

	fmt.Printf("Starting Run Recorder\n")
	sharedStash := ReadSharedStash(saveDir)
	runRecorder.StartRunRecorder(&sharedStash, groups)

}

func ReadSharedStash(saveDir string) domain.Stash {
	stashFilename := filepath.Join(saveDir, config.GetConfig().SharedStash.Filename)
	stash, _ := reader.ReadStashFromFile(stashFilename)
	return *stash
}
