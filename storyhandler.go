package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/joemahmah/gopher-write/story"
	"encoding/json"
	"html/template"
)

	/*
	TODO:
	-----
	
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}", testHandler) //Chapter overview
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/delete", testHandler) //delete chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/edit", testHandler) //edit chapter
	
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}", testHandler) //Section Overview
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}", testHandler) //delete section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}", testHandler) //edit section
	*/

func NewStoryHandler(w http.ResponseWriter, r *http.Request) {
	
	//Print log message
	LogNet.Println("Access " + r.URL.Path + " by "+ r.RemoteAddr)
	
	//Make new story
	newStory := story.Story{}
	
	//Decode the request
	err := json.NewDecoder(r.Body).Decode(&newStory)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		ActiveProject.AddStory(newStory)
		LogInfo.Println("Story " + newStory.Name.PrimaryName + " added to project " + ActiveProject.Name + ".")
	}
	
}


func NewChapterHandler(w http.ResponseWriter, r *http.Request) {
	
	//Print log message
	LogNet.Println("Access " + r.URL.Path + " by "+ r.RemoteAddr)
	
	//Make new char
	newChapter := story.Chapter{}
	
	//Decode the request
	err := json.NewDecoder(r.Body).Decode(&newChapter)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		ActiveProject.AddChapter(newChapter)
		LogInfo.Println("Chapter " + newChapter.Name.PrimaryName + " added to project " + ActiveProject.Name + ".")
	}
	
}


func NewSectionHandler(w http.ResponseWriter, r *http.Request) {
	
	//Print log message
	LogNet.Println("Access " + r.URL.Path + " by "+ r.RemoteAddr)
	
	//Make new char
	newSection := story.Section{}
	
	//Decode the request
	err := json.NewDecoder(r.Body).Decode(&newSection)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		ActiveProject.AddSection(newSection)
		LogInfo.Println("Section " + newSection.Name.PrimaryName + " added to project " + ActiveProject.Name + ".")
	}
	
}

func ViewStoryHandler(w http.ResponseWriter, r *http.Request) {
	
	//Print log message
	LogNet.Println("Access " + r.URL.Path + " by "+ r.RemoteAddr)
	
	//Get the selectedStory	UID
	suid, err := strconv.Atoi(mux.Vars(r)["storyuid"])
	
	//If unable to convert string to int
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
		
		return
	}
	
	//check if exists
	if _, exists := ActiveProject.Stories[suid]; exists {
		//Parse template
		tmpl, err := template.ParseFiles("data/templates/viewStory.tmpl")
		
		//if error parsing template
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			LogError.Println(err)
			
			return
		} 
		
		//serve template
		err = tmpl.Execute(w, nil)
		
		//If error, return code 500
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			LogError.Println(err)
		}
		
	} else {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println("Story with uid " + strconv.Itoa(suid) + " does not exist in project " + ActiveProject.Name + ".")
	}
	
}

func GetJSONStoryHandler(w http.ResponseWriter, r *http.Request) {
	
	//Print log message
	LogNet.Println("Access " + r.URL.Path + " by "+ r.RemoteAddr)

	//Get the char UID
	suid, err := strconv.Atoi(mux.Vars(r)["storyuid"])
	
	//Check if error convering into int
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} 
	
	//Encode and send off
	r.Header.Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(ActiveProject.Stories[suid])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} 
}

func EditStoryHandler(w http.ResponseWriter, r *http.Request) {
	
	ActiveProject.AddStory(story.Story{})
	
	//Print log message
	LogNet.Println("Access " + r.URL.Path + " by "+ r.RemoteAddr)

}

func ListJSONStoryHandler(w http.ResponseWriter, r *http.Request) {
	
	//Print log message
	LogNet.Println("Access " + r.URL.Path + " by "+ r.RemoteAddr)

	//Slices to store data
	var names []string
	var uids []int
	
	//Fill slices
	for _,elem := range ActiveProject.Stories{
		names = append(names, elem.Name.PrimaryName)
		uids = append(uids, elem.UID)
	}
	
	//Encode and send off
	err := json.NewEncoder(w).Encode(struct{Names []string; UIDS []int}{Names: names, UIDS: uids})
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} 
}

func OverviewStoryHandler(w http.ResponseWriter, r *http.Request) {
	
	//Print log message
	LogNet.Println("Access " + r.URL.Path + " by "+ r.RemoteAddr)
	
	//Parse template
	tmpl, err := template.ParseFiles("data/templates/overviewStories.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl")
	
	//if error parsing template
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
		
		return
	} 
	
	//serve template
	err = tmpl.Execute(w, nil)
	
	//IF error executing template
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}