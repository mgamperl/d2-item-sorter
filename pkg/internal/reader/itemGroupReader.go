package reader

import (
	"d2-item-sorter/pkg/domain"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadGroupConfigFile(filepath string) []domain.ItemGroup {
	groups := []domain.ItemGroup{}
	//test := Test{A: 1, B: Int(1), c: Int(1)}
	configFileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic("can not open config file: " + filepath)
	}

	err = yaml.Unmarshal(configFileData, &groups)

	if err != nil {
		panic("Error parsing group config file: " + err.Error())
	}

	return groups
}
