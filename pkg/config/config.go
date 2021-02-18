package config

type StashSizeConfig struct {
	Width  int
	Height int
}

type SharedStashConfig struct {
	Enabled  bool
	Filename string
}

type ConfigType struct {
	StashSize   StashSizeConfig
	SharedStash SharedStashConfig
}

func GetConfig() ConfigType {

	config := ConfigType{
		StashSize: StashSizeConfig{
			Width:  10,
			Height: 10,
		},
		SharedStash: SharedStashConfig{
			Enabled:  false,
			Filename: "_LOD_SharedStashSave.sss",
		},
	}

	return config
}
