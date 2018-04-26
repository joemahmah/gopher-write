package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"encoding/json"
	"github.com/joemahmah/gopher-write/common"
)

func EditSectionAddCharHandler(w http.ResponseWriter, r *http.Request){

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

func EditSectionSetNoteHandler(w http.ResponseWriter, r *http.Request){

	//Get the uids
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["chapteruid"])
	seuidRel, _ := strconv.Atoi(mux.Vars(r)["sectionuid"])

	//Get the section
	section, err := ActiveProject.GetSection(suid, cuidRel, seuidRel)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	inputData := &DataTransferText{}
	
	err = json.NewDecoder(r.Body).Decode(inputData)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogWarning.Println(err)
		return
	}
	
	//Send ok
	w.WriteHeader(http.StatusOK)
	
	//Set the note
	section.Note = inputData.Data
}

func EditSectionSetTextHandler(w http.ResponseWriter, r *http.Request){

	//Get the uids
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["chapteruid"])
	seuidRel, _ := strconv.Atoi(mux.Vars(r)["sectionuid"])

	//Get the section
	section, err := ActiveProject.GetSection(suid, cuidRel, seuidRel)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	inputData := &DataTransferText{}
	
	err = json.NewDecoder(r.Body).Decode(inputData)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogWarning.Println(err)
		return
	}
	
	//Send ok
	w.WriteHeader(http.StatusOK)
	
	//Set the text
	section.Text = inputData.Data
}

func EditSectionSetNameHandler(w http.ResponseWriter, r *http.Request){

	//Get the uids
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["chapteruid"])
	seuidRel, _ := strconv.Atoi(mux.Vars(r)["sectionuid"])

	//Get the section
	section, err := ActiveProject.GetSection(suid, cuidRel, seuidRel)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	inputData := &common.Name{}
	
	err = json.NewDecoder(r.Body).Decode(inputData)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogWarning.Println(err)
		return
	}
	
	//Send ok
	w.WriteHeader(http.StatusOK)
	
	//Set the name (name gets dereferenced and copied)
	section.Name = *inputData
}

func EditChapterSetNameHandler(w http.ResponseWriter, r *http.Request){

	//Get the uids
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["chapteruid"])

	//Get the chapter
	chapter, err := ActiveProject.GetChapter(suid, cuidRel)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	inputData := &common.Name{}
	
	err = json.NewDecoder(r.Body).Decode(inputData)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogWarning.Println(err)
		return
	}
	
	//Send ok
	w.WriteHeader(http.StatusOK)
	
	//Set the name
	chapter.Name = *inputData
}

func EditChapterSetStatusHandler(w http.ResponseWriter, r *http.Request){

	//Get the uids
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["chapteruid"])

	//Get the status number
	status, _ := strconv.Atoi(mux.Vars(r)["status"])

	//Get the chapter
	chapter, err := ActiveProject.GetChapter(suid, cuidRel)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogWarning.Println(err)
		return
	}
	
	//Send ok
	w.WriteHeader(http.StatusOK)
	
	//Set the status
	chapter.Status = status
}

func EditStorySetNameHandler(w http.ResponseWriter, r *http.Request){

	//Get the uids
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])

	//Get the story
	story, err := ActiveProject.GetStory(suid)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	inputData := &common.Name{}
	
	err = json.NewDecoder(r.Body).Decode(inputData)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogWarning.Println(err)
		return
	}
	
	//Send ok
	w.WriteHeader(http.StatusOK)
	
	//Set the name
	story.Name = *inputData
}

func EditStorySetStatusHandler(w http.ResponseWriter, r *http.Request){

	//Get the uids
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])

	//Get the status number
	status, _ := strconv.Atoi(mux.Vars(r)["status"])

	//Get the story
	story, err := ActiveProject.GetStory(suid)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogWarning.Println(err)
		return
	}
	
	//Send ok
	w.WriteHeader(http.StatusOK)
	
	//Set the status
	story.Status = status
}
