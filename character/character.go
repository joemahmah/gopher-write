package character

import (
	"github.com/joemahmah/gopher-write/common"
)

//Represents a character
type Character struct {
	UID			int //The UID is intended to be position in project character array

	Name		common.Name //The primary name of the character
	Aliases		[]common.Name //Aliases the character goes by
	Description	string //A brief description of the character
	Motivation	string //What drives the character
	Goal		string //What the character aims to do
	Role		string //The purpose of the character
	Age			common.Age //The age of the character
	
	
	//Character Profile Information
	BirthDate	string
	BirthPlace	string
	
}

//Checks if a string is the character's name
func (c *Character) IsName(name string) bool {
	return c.Name.IsName(name)
}

//Checks if a string is the character's alias
func (c *Character) IsAlias(name string) bool {

	for _, alias := range c.Aliases {
		if alias.IsName(name) {
			return true
		}
	}

	return false

}

func (c *Character) IsNameOrAlias(name string) bool {
	return c.IsName(name) || c.IsAlias(name)
}

