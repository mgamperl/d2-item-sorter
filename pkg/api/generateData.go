package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/nokka/d2s"
)

type uniqueItemData struct {
	Name string `json:"index"`
	Type string `json:"code"`
}

type uniqueItempMap map[string]uniqueItemData

func GenerateData(folder string) {

	content, err := ioutil.ReadFile(path.Join(folder, "UniqueItems.json"))

	if err != nil {
		log.Fatal(err)
	}

	var data uniqueItempMap

	if err := json.Unmarshal(content, &data); err != nil {
		fmt.Println("failed to unmarshal:", err)
	} else {

		keys := make([]string, 0, len(data))
		for k := range data {
			keys = append(keys, k)
		}

		sort.Slice(keys, func(i int, j int) bool {
			numA, _ := strconv.Atoi(keys[i])
			numB, _ := strconv.Atoi(keys[j])
			return numA < numB
		})

		idx := 0
		for _, k := range keys {

			t := d2s.WeaponCodes[d2s.ItemType(data[k].Type)]
			if t == "" {
				t = d2s.ArmorCodes[d2s.ItemType(data[k].Type)]
				if t == "" {
					t = d2s.ShieldCodes[d2s.ItemType(data[k].Type)]
					if t == "" {
						t = d2s.MiscCodes[d2s.ItemType(data[k].Type)]
					}
				}
			}

			if t == "" {
				t = data[k].Type
			}

			t = strings.ReplaceAll(t, " ", "")

			if data[k].Type != "" {
				fmt.Printf("%d: {ItemType: d2s.%s, Name: \"%s\"},\n", idx, t, data[k].Name)
				idx++
			} else if data[k].Type == "" && data[k].Name == "Expansion" {
			} else {
				fmt.Printf("%d: {Name: \"%s\"},\n", idx, data[k].Name)
				idx++
			}

		}
	}
}
