package reader

import (
	"d2-item-sorter/pkg/domain"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nokka/d2s"
)

func ReloadAndReturnCharacter(character domain.Character) domain.Character {
	c, _ := ReadCharacterFromFile(character.SaveFilename)
	return *c
}

func ReloadCharacter(character *domain.Character) error {
	character, err := ReadCharacterFromFile(character.SaveFilename)
	return err
}

func ReadCharacterFromFile(path string) (*domain.Character, error) {

	//TODO: own function for this to make it possible to reload a single character while running
	char, err := readD2sCharacter(path)

	if err != nil {
		fmt.Printf("error at %s: %v\n", path, err)
		return nil, err
	}

	stashFilename := strings.TrimSuffix(path, "d2s") + "d2x"

	characterStash, err := ReadStashFromFile(stashFilename)

	if err != nil {
		return nil, err
	}

	fmt.Printf("!found %c items in personal stash\n", characterStash.ItemCount)

	allExceptPersonalStashItems := domain.ItemFilterFunc(NewItemsFromCharacter(char.Items), func(item domain.Item) bool {
		//fmt.Printf("!found %s in character file: %s\n ", item.GetName(), item.Location)
		return true // item.Location != LocationPersonalStash
	})

	allExceptPersonalStashItems = append(allExceptPersonalStashItems, NewItemsFromMercenary(char.MercItems)...)

	return &domain.Character{
		Name:                  char.Header.Name.String(),
		Level:                 int(char.Header.Level),
		Class:                 char.Header.Class.String(),
		SaveFilename:          path,
		PersonalStashFilename: stashFilename,
		D2Char:                *char,
		Items:                 allExceptPersonalStashItems,
		PersonalStash:         characterStash,
	}, nil

}

func ReadAllCharactersFromPath(saveDir string) (map[string]domain.Character, error) {

	characterMap := map[string]domain.Character{}

	err := filepath.Walk(saveDir, func(path string, info os.FileInfo, err error) error {

		if strings.HasSuffix(path, "d2s") {
			character, err := ReadCharacterFromFile(path)

			if err != nil {
				return err
			}

			identifier := character.Name

			if _, exists := characterMap[identifier]; exists {
				return errors.New("character already exists: " + identifier)
			}

			characterMap[identifier] = *character

		}

		if info.IsDir() && path != saveDir {
			return filepath.SkipDir
		}
		return nil
	})

	return characterMap, err
}

func readD2sCharacter(path string) (*d2s.Character, error) {
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
