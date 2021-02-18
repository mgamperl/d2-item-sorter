package writer

import (
	"d2-item-sorter/pkg/domain"
)

func WriteCharacterToFile(character domain.Character, backupSuffix string) error {
	/*err := c.saveCharacterFileWithoutPersonalStash(backupSuffix)
	if err != nil {
		fmt.Printf("error writing character file: %v", err)
		return err
	}*/

	return WriteStashToFile(*character.PersonalStash, backupSuffix)
}
