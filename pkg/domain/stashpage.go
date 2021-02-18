package domain

import "d2-item-sorter/pkg/data"

type StashPage struct {
	Nr     int                `json:"number"`
	Name   string             `json:"name"`
	Items  []Item             `json:"items"`
	Type   data.StashPageType `json:"type"`
	spaces [][]int            `json:"-"`
	Width  int                `json:"width"`
	Height int                `json:"height"`
}

func NewStashPage(nr int, name string, stashPageType data.StashPageType, width int, height int) StashPage {
	page := StashPage{Nr: nr, Name: name, Type: stashPageType, Width: width, Height: height}

	page.spaces = make([][]int, width)
	for i := range page.spaces {
		page.spaces[i] = make([]int, height)
	}

	return page
}

func (p *StashPage) isCollision(xPosition int, xSize int, yPosition int, ySize int) bool {
	for x := 0; x < xSize; x++ {
		for y := 0; y < ySize; y++ {
			if xPosition+x > (p.Width-1) || yPosition+y > (p.Height-1) || p.spaces[xPosition+x][yPosition+y] > 0 {
				return true
			}
		}
	}
	return false
}

func (p *StashPage) allocateItem(xPosition int, xSize int, yPosition int, ySize int) {
	for x := 0; x < xSize; x++ {
		for y := 0; y < ySize; y++ {
			p.spaces[xPosition+x][yPosition+y]++
		}
	}
}

func (p *StashPage) InsertItem(item Item) bool {
	allocated := false
	for y := 0; y < p.Height; y++ {
		for x := 0; x < p.Width; x++ {
			if !allocated && !p.isCollision(x, item.Type.SizeX, y, item.Type.SizeY) {
				item.SetPosition(x, y)
				p.Items = append(p.Items, item)
				p.allocateItem(x, item.Type.SizeX, y, item.Type.SizeY)
				allocated = true
			}
		}
	}
	return allocated
}

func (p *StashPage) InsertItemAtColumn(item Item, column int) bool {
	allocated := false
	x := column
	for y := 0; y < p.Height; y++ {
		if !allocated && !p.isCollision(x, item.Type.SizeX, y, item.Type.SizeY) {
			item.SetPosition(x, y)
			p.Items = append(p.Items, item)
			p.allocateItem(x, item.Type.SizeX, y, item.Type.SizeY)
			allocated = true
		}
	}
	return allocated
}
