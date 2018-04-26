package story

import (
	"github.com/joemahmah/gopher-write/common"
	"errors"
)

type Story struct {
	UID			int //Intended to be the same as the index of project slice

	Name		common.Name
	Chapters	[]int
	Status		common.Status
	Note		string
}

func (s *Story) SwapChapters(first int, second int) error {
	if len(s.Chapters) < first || len(s.Chapters) < second {
		return errors.New("Chapter index out of bounds.")
	}

	s.Chapters[first], s.Chapters[second] = s.Chapters[second], s.Chapters[first]
	return nil
}

//TODO: make this happen via project
func (s *Story) SwapSectionBetweenChapters(chap1 int, chap2 int, sec1 int, sec2 int) error {
	if len(s.Chapters) < chap1 || len(s.Chapters) < chap2 {
		return errors.New("Chapter index out of bounds.")
	}
	//if len(s.Chapters[chap1].Sections) < sec1 || len(s.Chapters[chap2].Sections) < sec2 {
	//	return errors.New("Section index out of bounds.")
	//}

	//s.Chapters[chap1].Sections[sec1], s.Chapters[chap2].Sections[sec2] = s.Chapters[chap2].Sections[sec2], s.Chapters[chap1].Sections[sec1]
	return nil
}


///////////////////////////////////////////////////////////
//                  Getting Operations                   //
///////////////////////////////////////////////////////////

func (s *Story) GetChapterId(reluid int) (int, error) {
	if reluid >= len(s.Chapters){
		return -1, errors.New("Chapter uid out of bounds.")
	}
	
	return s.Chapters[reluid], nil
}
