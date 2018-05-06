package location

import (
	"github.com/joemahmah/gopher-write/common"
	"errors"
)

type Location struct {
	UID				int
	
	Name			common.Name
	Aliases			[]common.Name
	Description		string
	Parent			int
	Sublocations	map[int]int
}

func (l *Location) GetSublocationUID(index int) (int, error) {
	if len(l.Sublocations) < index {
		return 0, errors.New("Index out of bounds.")
	}

	return l.Sublocations[index], nil
}

func (l *Location) AddSublocation(uid int) {
	l.Sublocations[uid] = uid
}

func (l *Location) RemoveSublocation(uid int) {
		delete(l.Sublocations, uid)
}