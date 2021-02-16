package data

type ItemQuality string

const (
	ItemQualityNormal      ItemQuality = "Normal"
	ItemQualityExceptional ItemQuality = "Exceptional"
	ItemQualityElite       ItemQuality = "Elite"
)

type SortMode string

const (
	SortModeName       SortMode = "name"
	SortModeItemLevel  SortMode = "itemLevel"
	SortModeRuneNumber SortMode = "runeNumber"
	SortModeGemQuality SortMode = "gemQuality"
)
