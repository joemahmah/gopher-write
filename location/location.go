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
	Sublocations	[]uint
}

func (l *Location) GetSublocationUID(index int) (uint, error) {
	if len(l.Sublocations) < index {
		return 0, errors.New("Index out of bounds.")
	}

	return l.Sublocations[index], nil
}
