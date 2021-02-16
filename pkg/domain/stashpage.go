package domain

type StashPageType string

const (
	MainIndexPage StashPageType = "MainIndex"
	IndexPage     StashPageType = "Index"
	NormalPage    StashPageType = "Normal"
)

type StashPage struct {
	Nr                 uint16        `json:"number"`
	Name               string        `json:"name"`
	Items              []Item        `json:"items"`
	Type               StashPageType `json:"type"`
	spaces             [10][10]int
	isIgnoredPage      bool
	isPickupPage       bool
	isSharedPickupPage bool
}
