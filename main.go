package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"html/template"
)

func main(){

	//Init logs
	InitLogs()

	//Load projects
	err := LoadProjectList("./data/projects/projectList.json")

	//If unable to load projects
	if err != nil {
		LogError.Println("Unable to load projects, quitting.")
		LogError.Println(err)
		//return
	}

	//create router
	router := mux.NewRouter()

	/////////////////////////
	//       overview      //
	/////////////////////////
	router.HandleFunc("/", LandingHandler) //Landing page, routes if project open
	router.HandleFunc("/settings", testHandler) //Program settings

	/////////////////////////
	//      project io     //
	/////////////////////////
	router.HandleFunc("/project/load/{project:[0-9]{14}.[0-9]{6}}", LoadProjectHandler) //Load project
	router.HandleFunc("/project/save", SaveProjectHandler) //Save project
	router.HandleFunc("/project/new", NewProjectHandler) //Create new project
	router.HandleFunc("/project/list", ListJSONProjectHandler) //List projects
	router.HandleFunc("/project", OverviewProjectHandler) //Overview project
	router.HandleFunc("/project/import", ImportProjectHandler) //import project
	router.HandleFunc("/project/export", ExportProjectHandler) //export current project
	
	/////////////////////////
	//      char manip     //
	/////////////////////////
	router.HandleFunc("/char/new", NewCharHandler) //new char
	router.HandleFunc("/char/{cid:[0-9]{1,9}}", ViewCharHandler) //view char
	router.HandleFunc("/char/{cid:[0-9]{1,9}}/json", GetJSONCharHandler) //get char json
	router.HandleFunc("/char/{cid:[0-9]{1,9}}/edit", EditCharHandler) //edit char
	router.HandleFunc("/char/{cid:[0-9]{1,9}}/delete", DeleteCharHandler) //delete char
	router.HandleFunc("/char/{cid:[0-9]{1,9}}/edit/description", EditCharSetDescriptionHandler) //edit char description
	router.HandleFunc("/char/{cid:[0-9]{1,9}}/edit/motivation", EditCharSetMotivationHandler) //edit char motivation
	router.HandleFunc("/char/{cid:[0-9]{1,9}}/edit/goal", EditCharSetGoalHandler) //edit char goal
	router.HandleFunc("/char/{cid:[0-9]{1,9}}/edit/role", EditCharSetRoleHandler) //edit char role
	router.HandleFunc("/char/{cid:[0-9]{1,9}}/edit/name", EditCharSetNameHandler) //edit char name
	router.HandleFunc("/char/{cid:[0-9]{1,9}}/edit/age", EditCharSetAgeHandler) //edit char age
	router.HandleFunc("/char/{cid:[0-9]{1,9}}/edit/addalias", EditCharAddAliasHandler) //add alias
	router.HandleFunc("/char/{cid:[0-9]{1,9}}/edit/removealias", EditCharRemoveAliasHandler) //remove alias
	router.HandleFunc("/char/{cid:[0-9]{1,9}}/aliaslist", AliasListJSONCharHandler) //edit char age
	router.HandleFunc("/char/list",ListJSONCharHandler) //Char list json
	router.HandleFunc("/char", OverviewCharHandler) //Char overview

	/////////////////////////
	//    location manip   //
	/////////////////////////
	router.HandleFunc("/loc/new", NewLocationHandler) //new location
	router.HandleFunc("/loc/{lid:[0-9]{1,9}}", ViewLocationHandler) //view location
	router.HandleFunc("/loc/{lid:[0-9]{1,9}}/json", GetJSONLocationHandler) //get location json
	router.HandleFunc("/loc/{lid:[0-9]{1,9}}/edit", EditLocationHandler) //edit location
	router.HandleFunc("/loc/{lid:[0-9]{1,9}}/delete", DeleteLocationHandler) //delete location
	router.HandleFunc("/loc/{lid:[0-9]{1,9}}/edit/description", EditLocationSetDescriptionHandler) //edit location description
	router.HandleFunc("/loc/{lid:[0-9]{1,9}}/edit/name", EditLocationSetNameHandler) //edit location name
	router.HandleFunc("/loc/{lid:[0-9]{1,9}}/edit/parent", EditLocationSetParentHandler) //add sublocation
	router.HandleFunc("/loc/{lid:[0-9]{1,9}}/subloclist", SublocationListJSONLocationHandler) //serve json list of sublocations
	router.HandleFunc("/loc/{lid:[0-9]{1,9}}/edit/addalias", EditLocationAddAliasHandler) //add alias
	router.HandleFunc("/loc/{lid:[0-9]{1,9}}/edit/removealias", EditLocationRemoveAliasHandler) //remove alias
	router.HandleFunc("/loc/{lid:[0-9]{1,9}}/aliaslist", AliasListJSONLocationHandler) //serve json list of aliases
	router.HandleFunc("/loc/list",ListJSONLocationHandler) //Location list json
	router.HandleFunc("/loc", OverviewLocationHandler) //Location overview

	
	/////////////////////////
	//     story manip     //
	///////////////////////// 
		//Story
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}", ViewStoryHandler) //Story overview
	router.HandleFunc("/story/new", NewStoryHandler) //new story
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/delete", DeleteStoryHandler) //delete story
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/export", ExportStoryHandler) //export story
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/edit", EditStoryHandler) //edit story
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/json", GetJSONStoryHandler) //get story json
	router.HandleFunc("/story/list", ListJSONStoryHandler) //Story list json

		//chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}", ViewChapterHandler) //Chapter overview
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/new", NewChapterHandler) //new chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/import/{bindchar:true|false}/{bindloc:true|false}", ImportChapterHandler) //import chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/delete", DeleteChapterHandler) //delete chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/export", ExportChapterHandler) //export chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/edit", EditChapterHandler) //edit chapter
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/json", GetJSONChapterHandler) //get chapter json
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/list", ListJSONChapterHandler) //Chapter list json

		//section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}", ViewSectionHandler) //Section Overview
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/new", NewSectionHandler) //new section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/import/{bindchar:true|false}/{bindloc:true|false}", ImportSectionHandler) //import section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/delete", DeleteSectionHandler) //delete section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/export", ExportSectionHandler) //export section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/edit", EditSectionHandler) //edit section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/json", GetJSONSectionHandler) //get section json
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/list", ListJSONSectionHandler) //Chapter list json

		//editing (section)
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/edit/addchar/{charid:[0-9]{1,9}}", EditSectionAddCharHandler) //add character to section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/edit/removechar/{charindex:[0-9]{1,9}}", EditSectionRemoveCharHandler) //add character to section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/edit/addloc/{locid:[0-9]{1,9}}", EditSectionAddLocationHandler) //add location to section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/edit/status/{status:[0-4]}", EditSectionSetStatusHandler) //set status for section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/edit/note", EditSectionSetNoteHandler) //set note for section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/edit/text", EditSectionSetTextHandler) //set text for section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/{sectionuid:[0-9]{1,9}}/edit/name", EditSectionSetNameHandler) //set name for section

		//editing (chapter)
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/edit/status/{status:[0-4]}", EditChapterSetStatusHandler) //set status for section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/edit/note", EditChapterSetNoteHandler) //set name for section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/{chapteruid:[0-9]{1,9}}/edit/name", EditChapterSetNameHandler) //set name for section

		//editing (story)
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/edit/status/{status:[0-4]}", EditStorySetStatusHandler) //set status for section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/edit/note", EditStorySetNoteHandler) //set name for section
	router.HandleFunc("/story/{storyuid:[0-9]{1,9}}/edit/name", EditStorySetNameHandler) //set name for section

		//overview
	router.HandleFunc("/story", OverviewStoryHandler) //Story list
	router.HandleFunc("/story/move", SwapStoryHandler) //Story list

	/////////////////////////
	//   move operations   //
	/////////////////////////
	
	router.HandleFunc("/move/story/{first:[0-9]{1,9}}/{second:[0-9]{1,9}}", StoryMoveHandler) //move story positions
	router.HandleFunc("/move/chapter/intra/{fsuid:[0-9]{1,9}}/{first:[0-9]{1,9}}/{second:[0-9]{1,9}}", IntraChapterMoveHandler) //move chapter positions within a story (first in front of second)
	router.HandleFunc("/move/chapter/inter/{fsuid:[0-9]{1,9}}/{first:[0-9]{1,9}}/{ssuid:[0-9]{1,9}}/{second:[0-9]{1,9}}", InterChapterMoveHandler) //move chapter positions between stories (first in front of second)
	router.HandleFunc("/move/section/intra/{fsuid:[0-9]{1,9}}/{fcuid:[0-9]{1,9}}/{first:[0-9]{1,9}}/{second:[0-9]{1,9}}", IntraSectionMoveHandler) //move section positions within a chapter (first in front of second)

	/////////////////////////
	//      resources      //
	/////////////////////////

	router.HandleFunc("/res/note", testHandler)
	router.HandleFunc("/res/note/list", testHandler)
	router.HandleFunc("/res/note/add", testHandler)
	router.HandleFunc("/res/note/{nid:[0-9]{1,9}}", testHandler)
	router.HandleFunc("/res/note/{nid:[0-9]{1,9}}/json", testHandler)
	router.HandleFunc("/res/note/{nid:[0-9]{1,9}}/delete", testHandler)
	router.HandleFunc("/res/note/{nid:[0-9]{1,9}}/edit", testHandler)
	router.HandleFunc("/res/link", testHandler)
	router.HandleFunc("/res/link/list", testHandler)
	router.HandleFunc("/res/link/add", testHandler)
	router.HandleFunc("/res/link/{lid:[0-9]{1,9}}", testHandler)
	router.HandleFunc("/res/link/{lid:[0-9]{1,9}}/json", testHandler)
	router.HandleFunc("/res/link/{lid:[0-9]{1,9}}/delete", testHandler)
	router.HandleFunc("/res/link/{lid:[0-9]{1,9}}/edit", testHandler)

	/////////////////////////
	//        server       //
	/////////////////////////
	
	//Have the router add access log middleware
	router.Use(LogMiddleware)
	
	//Create the server
	server := &http.Server{
		Addr: ":8080",
		Handler: router,
	}
	
	//use a goroutine to handle the server
	go func() {
		//Listen and serve
		err := server.ListenAndServe()
	
		//Report errors
		if err != nil && err != http.ErrServerClosed {
			LogError.Println(err)
			SaveAndExit(server)
		} 
	}()
	
	//A channel to flag done (primarily used to block main thread)
	done := make(chan bool, 1)
	
	//use a goroutine to handle signals
	sig := make(chan os.Signal, 1)
	
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) //bind signals to sig channel
	
	go func() {
		<-sig //wait for signal
		LogInfo.Println("Signal received. Shutting down...") //inform of shutdown
		SaveAndExit(server) //shutdown
		done <- true //Shouldn't reach here, but if so, let program quit
	}()
	
	LogInfo.Println("Waiting for interrupt signal or net command...")
	LogInfo.Println("To quit, either send a request to '/settings/quit' or send a signal to terminate.")
	LogInfo.Println("Both methods should ensure proper shutdown.")
	LogInfo.Println("NOTE: ctrl-c in a cygwin terminal may not have proper shutdown.")
	<- done //Intended to block
}

func SaveAndExit(server *http.Server){
	//Shutdown server
	err := server.Shutdown(nil)
	
	if err != nil {
		LogError.Println(err)
	}

	//If a valid project is not loaded
	if !ValidProjectLoaded {
		LogWarning.Println("A valid project is not loaded. Aborting save.")
		os.Exit(0)
	}
	
	//Save Project
	projectPath := "./data/projects/" + ActiveProject.SaveName + ".json"
	SaveProject(ActiveProject, projectPath)
	
	//Save Project List
	SaveProjectList("./data/projects/projectList.json")

	//Exit Program
	os.Exit(0)
}

//Landing Handler
func LandingHandler(w http.ResponseWriter, r *http.Request){
	//Parse the templates
	tmpl, err := template.ParseFiles("data/templates/landing.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl", "data/templates/footer.tmpl")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
		return
	}

	err = tmpl.Execute(w, ValidProjectLoaded)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

//Dummy handler
func testHandler(w http.ResponseWriter, r *http.Request) {
	//DO NOTHING
}

//Logging middleware
func LogMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		LogNet.Println("Access " + r.URL.Path + " by " + r.RemoteAddr)
		h.ServeHTTP(w, r)
	})
}

