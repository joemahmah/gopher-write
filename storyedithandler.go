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

	//Get the id of the character to be added
	newChar, _ := strconv.Atoi(mux.Vars(r)["charid"])

	//Get the section
	section, err := ActiveProject.GetSection(suid, cuidRel, seuidRel)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
	} else {
		//Send ok
		w.WriteHeader(http.StatusOK)

		//Add character
		section.Characters = append(section.Characters, newChar)
	
	}
}

func EditSectionAddLocationHandler(w http.ResponseWriter, r *http.Request){

	LogNet.Println("Access " + r.URL.Path + " by " + r.RemoteAddr)

	//Get the uids
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["chapteruid"])
	seuidRel, _ := strconv.Atoi(mux.Vars(r)["sectionuid"])

	//Get the id of the character to be added
	newLoc, _ := strconv.Atoi(mux.Vars(r)["locid"])

	//Get the section
	section, err := ActiveProject.GetSection(suid, cuidRel, seuidRel)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
	} else {
		//Send ok
		w.WriteHeader(http.StatusOK)

		//Add character
		section.Locations = append(section.Locations, newLoc)
	
	}
}

func EditSectionSetStatusHandler(w http.ResponseWriter, r *http.Request){

	LogNet.Println("Access " + r.URL.Path + " by " + r.RemoteAddr)

	//Get the uids
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["chapteruid"])
	seuidRel, _ := strconv.Atoi(mux.Vars(r)["sectionuid"])

	//Get the id of the character to be added
	status, _ := strconv.Atoi(mux.Vars(r)["status"])

	//Get the section
	section, err := ActiveProject.GetSection(suid, cuidRel, seuidRel)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
	} else {
		//Send ok
		w.WriteHeader(http.StatusOK)

		//Add character
		section.Status = status
	
	}
}
