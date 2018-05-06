package story

import (
	"github.com/joemahmah/gopher-write/common"
	"errors"
)

type Chapter struct {
	UID			int

	Name		common.Name
	Sections	[]int
	Status		common.Status
	Note		string
	Summary		string
	Purpose		string
}

func (c *Chapter) SwapSections(first int, second int) error{
	if len(c.Sections) < first || len(c.Sections) < second {
		return errors.New("Section index out of bounds.")
	}

	c.Sections[first], c.Sections[second] = c.Sections[second], c.Sections[first]
	return nil
}

///////////////////////////////////////////////////////////
//                  Getting Operations                   //
///////////////////////////////////////////////////////////

func (c *Chapter) GetSectionId(reluid int) (int, error) {
	if reluid >= len(c.Sections){
		return -1, errors.New("Section uid out of bounds.")
	}
	
	return c.Sections[reluid], nil
}
