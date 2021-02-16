package domain

import (
	"bufio"
)

type SharedGold struct {
	Amount int    `json:"amount"`
	Bytes  []byte `json:"-"`
}

type StashType string

const (
	StashTypeShared    StashType = "SSS"
	StashTypeCharacter StashType = "CST"
	StashTypeUnknown   StashType = ""
)

type Stash struct {
	StashType  StashType     `json:"stashType"`
	SharedGold SharedGold    `json:"sharedGold"`
	StashPages []StashPage   `json:"stashPages"`
	ItemCount  int           `json:"itemCount"`
	Filepath   string        `json:"filePath"`
	extraByte  byte          `json:"-"`
	version    string        `json:"-"`
	pageCount  int           `json:"-"`
	bfr        *bufio.Reader `json:"-"`
}
