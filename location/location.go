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
	Sublocations	[]int
}

func (l *Location) GetSublocationUID(index int) (int, error) {
	if len(l.Sublocations) < index {
		return 0, errors.New("Index out of bounds.")
	}

	return l.Sublocations[index], nil
}
