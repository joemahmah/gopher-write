package story

import (
	"github.com/joemahmah/gopher-write/common"
	"errors"
)

type Chapter struct {
	Name		common.Name
	Sections	[]Section
	Status		common.Status
}

func (c *Chapter) SwapSections(first int, second int) error{
	if len(c.Sections) < first || len(c.Sections) < second {
		return errors.New("Section index out of bounds.")
	}

	c.Sections[first], c.Sections[second] = c.Sections[second], c.Sections[first]
	return nil
}
