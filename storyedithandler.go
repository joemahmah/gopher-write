package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func EditSectionAddCharHandler(w http.ResponseWriter, r *http.Request){

	LogNet.Println("Access " + r.URL.Path + " by " + r.RemoteAddr)

	//Get the uids
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["chapteruid"])
	seuidRel, _ := strconv.Atoi(mux.Vars(r)["sectionuid"])

	//calc uids from relative uids
	cuid := ActiveProject.Stories[suid].Chapters[cuidRel]
	seuid := ActiveProject.Chapters[cuid].Sections[seuidRel]

	newChar, _ := strconv.Atoi(mux.Vars(r)["charid"])

	if _, exists := ActiveProject.Characters[newChar]; exists {

		if selectedSection, exists := ActiveProject.Sections[seuid]; exists {
			w.WriteHeader(http.StatusOK)

			selectedSection.Characters = append(selectedSection.Characters, newChar)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println("Character with uid " + strconv.Itoa(newChar) + " does not exist.")
	}
}
