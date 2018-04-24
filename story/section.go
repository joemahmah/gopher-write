package story

import (
	"github.com/joemahmah/gopher-write/common"
)

type Section struct {
	UID			int

	Name		common.Name
	Text		string
	Status		common.Status
	Characters	[]int
	Locations	[]int
	Note		string
}

