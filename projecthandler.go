package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"html/template"
	"encoding/json"
	"sort"
)

//Vars to make sure the project list json doesn't need to be regenerated all the time
var projectListHasChanged bool = false
var projectListDataTransferSlice []DualStringMonoBool = nil

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
		projectListHasChanged = true
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
		ProjectList[projectSaveName] = projectName //Add to project list
	}
	
}

func OverviewProjectHandler(w http.ResponseWriter, r *http.Request) {
	//Parse template
	tmpl, err := template.ParseFiles("data/templates/overviewProject.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl")

	//check for errors
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
		return
	}

	//serve template
	err = tmpl.Execute(w, nil)

	//report errors
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

//Provides JSON containing an array of pairs with the pairs
//being (name, location) of projects
func ListJSONProjectHandler(w http.ResponseWriter, r *http.Request) {
	//var to hold names/paths
	var data DataTransferDualStringMonoBoolSlice

	if projectListHasChanged || projectListDataTransferSlice == nil { //If need to update cached data
		//Assign to data
		for path, name := range ProjectList {
			isActiveProj := false
			if ActiveProject.SaveName == path {
				isActiveProj = true
			}
			data.Data = append(data.Data, DualStringMonoBool{S1: name, S2: path, B: isActiveProj})
		}
		
		//Sort the data (to ensure it will always be the same)
		sort.Slice(data.Data, func(i int, j int) bool {
			return data.Data[i].S2 < data.Data[j].S2
		})
	
		projectListHasChanged = false;
		projectListDataTransferSlice = data.Data
	} else {
		data.Data = projectListDataTransferSlice
	}

	//Encode
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}
