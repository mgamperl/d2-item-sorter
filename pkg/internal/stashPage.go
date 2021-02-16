package internal

import "d2-item-sorter/pkg/domain"

type StashPage struct {
	Nr                 uint16               `json:"number"`
	Name               string               `json:"name"`
	Items              []SortableItem       `json:"items"`
	Type               domain.StashPageType `json:"type"`
	spaces             [10][10]int
	isIgnoredPage      bool
	isPickupPage       bool
	isSharedPickupPage bool
}

func (p *StashPage) isCollision(xPosition int, xSize int, yPosition int, ySize int) bool {
	for x := 0; x < xSize; x++ {
		for y := 0; y < ySize; y++ {
			if xPosition+x > 9 || yPosition+y > 9 || p.spaces[xPosition+x][yPosition+y] > 0 {
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

func (p *StashPage) insertItem(item SortableItem) bool {
	allocated := false
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if !allocated && !p.isCollision(x, item.AdditionalProps.SizeX, y, item.AdditionalProps.SizeY) {
				item.setPosition(x, y)
				p.Items = append(p.Items, item)
				p.allocateItem(x, item.AdditionalProps.SizeX, y, item.AdditionalProps.SizeY)
				allocated = true
			}
		}
	}
	return allocated
}

func (p *StashPage) insertItemAtColumn(item SortableItem, column int) bool {
	allocated := false
	x := column
	for y := 0; y < 10; y++ {
		if !allocated && !p.isCollision(x, item.AdditionalProps.SizeX, y, item.AdditionalProps.SizeY) {
			item.setPosition(x, y)
			p.Items = append(p.Items, item)
			p.allocateItem(x, item.AdditionalProps.SizeX, y, item.AdditionalProps.SizeY)
			allocated = true
		}
	}
	return allocated
}
