package domain

import "time"

type Run struct {
	Counter   int
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	Character string
	RunType   string
	NewItems  []Item
	Items     []Item
}
