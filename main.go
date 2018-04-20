package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main(){

	//Init logs
	InitLogs()

	//create router
	router := mux.NewRouter()

	//Add overview pages
	router.HandleFunc("/", testHandler) //Landing page, routes if project open
	router.HandleFunc("/list", testHandler) //List projects
	router.HandleFunc("/settings", testHandler) //Program settings

	//Add project io routes
	router.HandleFunc("/project/load/{project:[0-9]{14}.[0-9]{6}}", LoadProjectHandler)
	router.HandleFunc("/project/save", SaveProjectHandler)
	router.HandleFunc("/project/new/{project}", NewProjectHandler)

	//Add char manip routes
	router.HandleFunc("/char/new", NewCharHandler) //new char
	router.HandleFunc("/char/view/{cid:[0-9]{1,9}}", ViewCharHandler) //view char
	router.HandleFunc("/char/json/{cid:[0-9]{1,9}}", GetJSONCharHandler) //get char json
	router.HandleFunc("/char/edit/{cid:[0-9]{1,9}}", EditCharHandler) //edit char
	router.HandleFunc("/char/list",ListJSONCharHandler) //Char list json
	router.HandleFunc("/char", OverviewCharHandler) //Char overview

	//Add story manip routes 
		//Story
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}", ViewStoryHandler) //Story overview
	router.HandleFunc("/story/new", NewStoryHandler) //new story
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/delete", testHandler) //delete story
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/edit", EditStoryHandler) //edit story
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/json", GetJSONStoryHandler) //get story json
	router.HandleFunc("/story/list", ListJSONStoryHandler) //Story list json

		//chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}", ViewChapterHandler) //Chapter overview
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/new", NewChapterHandler) //new chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/delete", testHandler) //delete chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/edit", testHandler) //edit chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/json", testHandler) //get chapter json
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/list", ListJSONChapterHandler) //Chapter list json

		//section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}", ViewSectionHandler) //Section Overview
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/new", NewSectionHandler) //new section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/delete", testHandler) //delete section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/edit", testHandler) //edit section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/json", testHandler) //get section json
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/list", ListJSONSectionHandler) //Chapter list json

		//overview
	router.HandleFunc("/story", OverviewStoryHandler) //Story list

	http.ListenAndServe(":8080", router)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	//DO NOTHING
}
