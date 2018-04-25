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

	/////////////////////////
	//       overview      //
	/////////////////////////
	router.HandleFunc("/", testHandler) //Landing page, routes if project open
	router.HandleFunc("/settings", testHandler) //Program settings

	/////////////////////////
	//      project io     //
	/////////////////////////
	router.HandleFunc("/project/load/{project:[0-9]{14}.[0-9]{6}}", LoadProjectHandler)
	router.HandleFunc("/project/save", SaveProjectHandler)
	router.HandleFunc("/project/new/{project}", NewProjectHandler)
	router.HandleFunc("/project/list", testHandler) //List projects

	
	/////////////////////////
	//      char manip     //
	/////////////////////////
	router.HandleFunc("/char/new", NewCharHandler) //new char
	router.HandleFunc("/char/{cid:[0-9]{1,9}}/view", ViewCharHandler) //view char
	router.HandleFunc("/char/{cid:[0-9]{1,9}}/json", GetJSONCharHandler) //get char json
	router.HandleFunc("/char/{cid:[0-9]{1,9}}/edit", EditCharHandler) //edit char
	router.HandleFunc("/char/list",ListJSONCharHandler) //Char list json
	router.HandleFunc("/char", OverviewCharHandler) //Char overview

	
	/////////////////////////
	//     story manip     //
	///////////////////////// 
		//Story
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}", ViewStoryHandler) //Story overview
	router.HandleFunc("/story/new", NewStoryHandler) //new story
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/delete", DeleteStoryHandler) //delete story
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/edit", EditStoryHandler) //edit story
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/json", GetJSONStoryHandler) //get story json
	router.HandleFunc("/story/list", ListJSONStoryHandler) //Story list json

		//chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}", ViewChapterHandler) //Chapter overview
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/new", NewChapterHandler) //new chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/delete", DeleteChapterHandler) //delete chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/edit", EditChapterHandler) //edit chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/json", GetJSONChapterHandler) //get chapter json
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/list", ListJSONChapterHandler) //Chapter list json

		//section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}", ViewSectionHandler) //Section Overview
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/new", NewSectionHandler) //new section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/delete", DeleteSectionHandler) //delete section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/edit", EditSectionHandler) //edit section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/json", GetJSONSectionHandler) //get section json
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/list", ListJSONSectionHandler) //Chapter list json

		//editing
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/edit/addchar/{charid:[0-9]{1,9}}", EditSectionAddCharHandler) //add character to section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/edit/addloc/{locid:[0-9]{1,9}}", EditSectionAddLocationHandler) //add character to section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/edit/status/{status:[0-4]}", EditSectionSetStatusHandler) //add character to section

		//overview
	router.HandleFunc("/story", OverviewStoryHandler) //Story list

	/////////////////////////
	//   move operations   //
	/////////////////////////
	
	router.HandleFunc("/move/story/{first:[0-9]{1,9}}/{second:[0-9]{1,9}}", StoryMoveHandler) //move story positions
	router.HandleFunc("/move/chapter/intra/{suid:[0-9]{1,9}}/{first:[0-9]{1,9}}/{second:[0-9]{1,9}}", IntraChapterMoveHandler) //move chapter positions within a story (first in front of second)
	router.HandleFunc("/move/chapter/inter/{fsuid:[0-9]{1,9}}/{first:[0-9]{1,9}}/{ssuid:[0-9]{1,9}}/{second:[0-9]{1,9}}", InterChapterMoveHandler) //move chapter positions between stories (first in front of second)
	router.HandleFunc("/move/section/intra/{suid:[0-9]{1,9}}/{cuid:[0-9]{1,9}}/{first:[0-9]{1,9}}/{second:[0-9]{1,9}}", IntraSectionMoveHandler) //move section positions within a chapter (first in front of second)

	
	/////////////////////////
	//        server       //
	/////////////////////////
	http.ListenAndServe(":8080", router)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	//DO NOTHING
}
