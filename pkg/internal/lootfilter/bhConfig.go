package lootfilter

import (
	"d2-item-sorter/pkg/domain"
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
)

func WriteBHConfig(installDir string, items []domain.Item) {

	fmt.Printf("Writing BH Config File: %s\n", installDir)
	content, _ := ioutil.ReadFile(path.Join(installDir, "BH.cfg"))

	fmt.Printf("Länge BH.cfg: %d\n", len(content))

	regexp, err := regexp.Compile(`(?s)(.*)(//d2-item-sorter START\s)(.*)(//d2-item-sorter END\s)(.*)`)

	if err == nil {
		result := regexp.ReplaceAllString(string(content), "$1${2}"+generateBHConfig(items)+"\n$4$5")
		fmt.Printf("Länge Result: %d\n", len(result))
		err := ioutil.WriteFile(path.Join(installDir, "BH.cfg"), []byte(result), 0777)
		if err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func generateBHConfig(items []domain.Item) string {

	result := ""
	alreadyWroteUniques := map[string]int{"rin": 1, "amu": 1}
	alreadyWroteSets := map[string]int{"rin": 1, "amu": 1, "tgl": 1}

	for _, item := range items {

		if item.Quality == domain.ItemQualityUnique {
			if _, ok := alreadyWroteUniques[item.Type.ID]; ok == false {
				result += fmt.Sprintf("ItemDisplay[!ID UNI %s]: %%GOLD%%%%NAME%% %%GRAY%%(owned)\n", item.Type.ID)
				alreadyWroteUniques[item.Type.ID] = 1
			}
		} else if item.Quality == domain.ItemQualitySet {
			if _, ok := alreadyWroteSets[item.Type.ID]; ok == false {
				result += fmt.Sprintf("ItemDisplay[!ID SET %s]: %%GREEN%%%%NAME%% %%GRAY%%(owned)\n", item.Type.ID)
				alreadyWroteSets[item.Type.ID] = 1
			}
		}

	}

	return result
}
