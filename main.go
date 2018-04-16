package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
	"fmt"
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
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}", testHandler) //Story overview
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/new", testHandler) //new story
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/delete", testHandler) //delete story
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/edit", testHandler) //edit story
	
		//chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}", testHandler) //Chapter overview
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/new", testHandler) //new chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/delete", testHandler) //delete chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/edit", testHandler) //edit chapter
	
		//section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}", testHandler) //Section Overview
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}", testHandler) //new section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}", testHandler) //delete section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}", testHandler) //edit section
	
		//overview
	router.HandleFunc("/story/list", testHandler) //Story list json
	router.HandleFunc("/story", testHandler) //Story list
	
	http.ListenAndServe(":8080", router)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	//t := template.New("testTemplate")
	t, err := template.ParseFiles("tmpl/test.html")

	if err != nil {
		fmt.Println("Error loading template.")
		return
	}
	fmt.Println("Loaded template.")

	//char := character.Character{Name: common.Name{PrimaryName: "test"}}
	err = t.Execute(w, nil)

	if err != nil {
		fmt.Printf("Error executing template: %s\n", err.Error())
	}
}
