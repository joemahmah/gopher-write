package common

import (
	"strings"
)

//A representation of a generic name
type Name struct {
	PrimaryName	string //The primary name 
	AlternateNames	[]string //Alternate names (nicknames, partial names)
	IsAlias		bool
}

func (n *Name) IsPrimaryName(name string) bool {
	return strings.EqualFold(n.PrimaryName, name)
}

func (n *Name) IsAlternateName(name string) bool {
	
	for _, altName := range n.AlternateNames {
		if strings.EqualFold(altName, name) {
			return true
		}
	}

	return false
}

func (n *Name) IsName(name string) bool {
	return n.IsPrimaryName(name) || n.IsAlternateName(name)
}