package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

func LoadProjectHandler(w http.ResponseWriter, r *http.Request) {
	projectName := mux.Vars(r)["project"]
	projectPath := "./data/projects/" + projectName + ".json"
	
	//Print log message
	LogNet.Println("Access " + r.URL.Path + " by "+ r.RemoteAddr)
	
	//load project
	err := LoadProject(ActiveProject, projectPath)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		LogInfo.Println("Loaded project " + ActiveProject.Name + " (" + projectPath + ")")
	}
	
}

func SaveProjectHandler(w http.ResponseWriter, r *http.Request) {
	projectPath := "./data/projects/" + ActiveProject.SaveName + ".json"
	
	//Print log message
	LogNet.Println("Access " + r.URL.Path + " by "+ r.RemoteAddr)
	
	//load project
	err := SaveProject(ActiveProject, projectPath)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		LogInfo.Println("Saved project " + ActiveProject.Name + " (" + projectPath + ")")
	}
}

func NewProjectHandler(w http.ResponseWriter, r *http.Request) {
	projectName := mux.Vars(r)["project"]
	projectSaveName := time.Now().Format("20060102150405.000000")
	projectPath := "./data/projects/" + projectSaveName + ".json"
	
	ActiveProject = MakeProject(projectName, projectSaveName)
	
	//Print log message
	LogNet.Println("Access " + r.URL.Path + " by "+ r.RemoteAddr)
	
	//load project
	err := SaveProject(ActiveProject, projectPath)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		LogInfo.Println("Created project " + ActiveProject.Name + " (" + projectPath + ")")
	}
	
}