package domain

type GrailStatusItem struct {
	Name           string   `json:"name"`
	Type           string   `json:"type"`
	ImageID        string   `json:"imageId"`
	IsUnique       bool     `json:"isUnique"`
	IsSet          bool     `json:"isSet"`
	IsRune         bool     `json:"isRune"`
	SetName        string   `json:"setName"`
	Count          int      `json:"count"`
	Order          int      `json:"order"`
	ItemQuality    string   `json:"quality"`
	Category       string   `json:"category"`
	SubCategory    string   `json:"subCategory"`
	SubSubCategory string   `json:"subSubCategory"`
	TypeName       string   `json:"typeName"`
	StoredAt       []string `json:"storedAt"`
}
