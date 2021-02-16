package domain

import "github.com/nokka/d2s"

type Character struct {
	Name                  string `json:"name"`
	Level                 int    `json:"level"`
	Class                 string `json:"class"`
	Items                 []Item `json:"items"`
	SaveFilename          string
	PersonalStashFilename string
	d2Char                d2s.Character
}
