package data

type ItemTypeQuality string

const (
	ItemQualityUnknown     ItemTypeQuality = "Unknown"
	ItemQualityNormal      ItemTypeQuality = "Normal"
	ItemQualityExceptional ItemTypeQuality = "Exceptional"
	ItemQualityElite       ItemTypeQuality = "Elite"
)

type SortMode string

const (
	SortModeName       SortMode = "name"
	SortModeItemLevel  SortMode = "itemLevel"
	SortModeRuneNumber SortMode = "runeNumber"
	SortModeGemQuality SortMode = "gemQuality"
)
