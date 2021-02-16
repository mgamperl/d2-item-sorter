package internal

import (
	"bufio"
	"d2-item-sorter/pkg/domain"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nokka/d2s"
)

type Character struct {
	Name              string         `json:"name"`
	CharacterFilename string         `json:"charFilename"`
	StashFilename     string         `json:"stashFilename"`
	D2sCharacter      d2s.Character  `json:"details"`
	Items             []SortableItem `json:"items"`
	MercItems         []SortableItem `json:"mercItems"`
	PersonalStash     Stash          `json:"personalStash"`
}

func (c *Character) GetAllItems() []SortableItem {
	allItems := []SortableItem{}

	allItems = append(allItems, c.PersonalStash.GetAllItems()...)
	allItems = append(allItems, c.Items...)

	return allItems
}

func filter(ss []SortableItem, test func(SortableItem) bool) (ret []SortableItem) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func (c *Character) GetPersonalStashItems() []SortableItem {
	return c.PersonalStash.GetAllItems()
}

func (c *Character) GetBeltItems() []SortableItem {
	return filter(c.Items, func(item SortableItem) bool {
		return item.Location == domain.LocationBelt
	})
}

func (c *Character) GetInventoryItems() []SortableItem {
	return filter(c.Items, func(item SortableItem) bool {
		return item.Location == domain.LocationInventory
	})
}

func (c *Character) GetEquippedItems() []SortableItem {
	return filter(c.Items, func(item SortableItem) bool {
		return item.Location == domain.LocationEquipped
	})
}

func (c *Character) GetMercItems() []SortableItem {
	return c.MercItems
}

func (c *Character) Reload() {
	//TODO: add c.reloadCharacter to reload the inventory items which might change while this is running
	c.PersonalStash.Reload()
}

func (c *Character) SortPersonalStash(groups []ItemGroup) {
	c.PersonalStash.SortItems(groups)
}

func (c *Character) Save(backupSuffix string) error {
	/*err := c.saveCharacterFileWithoutPersonalStash(backupSuffix)
	if err != nil {
		fmt.Printf("error writing character file: %v", err)
		return err
	}*/
	err := c.PersonalStash.Save(backupSuffix)
	return err
}

func (c *Character) saveCharacterFileWithoutPersonalStash(backupSuffix string) error {

	itemCount := 0
	bytesToWrite := []byte{}
	//read characterfile, keep everything the same except for the items
	file, err := os.Open(c.CharacterFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	bfr := bufio.NewReader(file)

	firstPart, err := bfr.ReadBytes(0x4D) // we are looking for the M in JM
	if err != nil {
		fmt.Printf("error reading first part 1: %v\n", firstPart)
		return err
	}

	foundItems := false
	for foundItems == false {

		bytesToWrite = append(bytesToWrite, firstPart...)
		fmt.Printf("firstPart: %v\n", firstPart)
		if firstPart[len(firstPart)-2] == 0x4A { //is previous byte our J?
			fmt.Printf("found items spot\n")
			foundItems = true
			break
		} else {
			fmt.Printf("not right spot: %v\n", firstPart[len(firstPart)-1])
			firstPart, err = bfr.ReadBytes(0x4D)
			if err != nil {
				fmt.Printf("error reading first part 2: %v\n", firstPart)
				return err
			}
		}
	}

	bytesToWrite = append(bytesToWrite, byte(len(c.Items)), 0x0)

	//fmt.Printf("Stash Page (%s) #%d: %d items\n", page.Name, page_idx, len(page.Items))

	for _, item := range c.Items {
		itemCount++
		bytesToWrite = append(bytesToWrite, item.D2sItem.OriginalData...)
	}

	backupAndWrite(c.CharacterFilename, bytesToWrite, backupSuffix)

	fmt.Printf("wrote %d items to: %s\n", itemCount, c.CharacterFilename)

	return nil
}

func makeSortableItem(item d2s.Item, fromMercenary bool) SortableItem {
	sortable := SortableItem{}
	sortable.SetItem(item, domain.StashTypeCharacter, fromMercenary)
	return sortable
}

func makeSortableItems(items []d2s.Item, fromMercenary bool) []SortableItem {
	result := []SortableItem{}
	for _, item := range items {
		result = append(result, makeSortableItem(item, fromMercenary))
	}
	return result
}

func ParseCharacters(saveDir string) (map[string]*Character, error) {
	characters := []d2s.Character{}

	characterMap := map[string]*Character{}

	err := filepath.Walk(saveDir, func(path string, info os.FileInfo, err error) error {

		if strings.HasSuffix(path, "d2s") {

			//TODO: own function for this to make it possible to reload a single character while running
			char, err := parseCharacter(path)
			if err != nil {
				fmt.Printf("error at %s: %v\n", path, err)
				return err
			}

			identifier := char.Header.Name.String()

			if _, exists := characterMap[identifier]; exists {
				return errors.New("character already exists: " + identifier)
			}

			stashFilename := strings.TrimSuffix(path, "d2s") + "d2x"

			characterStash := Stash{}

			err = characterStash.ParseStash(stashFilename)

			if err != nil {
				return err
			}

			fmt.Printf("!found %c items in personal stash\n", characterStash.ItemCount)

			allExceptPersonalStashItems := filter(makeSortableItems(char.Items, false), func(item SortableItem) bool {
				//fmt.Printf("!found %s in character file: %s\n ", item.GetName(), item.Location)
				return true // item.Location != LocationPersonalStash
			})

			characterMap[identifier] = &Character{
				Name:              char.Header.Name.String(),
				CharacterFilename: path,
				StashFilename:     stashFilename,
				D2sCharacter:      *char,
				PersonalStash:     characterStash,
				Items:             allExceptPersonalStashItems, //PersonalStashItems are stored in PersonalStash as we can also change them
				MercItems:         makeSortableItems(char.MercItems, true),
			}

			characters = append(characters, *char)
		} else if strings.HasSuffix(path, "d2x") {
			stash := Stash{}
			err = stash.ParseStash(path)
			//fmt.Printf("%s: %d items\n", path, len(stash.GetAllItems()))
		} else {
			//fmt.Printf("no d2s SUFF in %s (%s)\n", path, saveDir)
		}

		if info.IsDir() && path != saveDir {
			return filepath.SkipDir
		}
		return nil
	})

	return characterMap, err
}

func parseCharacter(path string) (*d2s.Character, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while opening .d2s file", err)
	}

	defer file.Close()

	char, err := d2s.Parse(file)
	if err != nil {
		return nil, err
	}

	return char, nil
}
