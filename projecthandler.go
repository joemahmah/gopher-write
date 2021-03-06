package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"sort"
	"time"
)

//Vars to make sure the project list json doesn't need to be regenerated all the time
var projectListHasChanged bool = false
var projectListDataTransferSlice []DualStringMonoBool = nil

func LoadProjectHandler(w http.ResponseWriter, r *http.Request) {
	projectName := mux.Vars(r)["project"]
	projectPath := "./data/projects/" + projectName + ".json"

	//load project
	err := LoadProject(ActiveProject, projectPath)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		LogInfo.Println("Loaded project " + ActiveProject.Name + " (" + projectPath + ")")
		projectListHasChanged = true

		//Flag valid project loaded
		ValidProjectLoaded = true
	}

}

func SaveProjectHandler(w http.ResponseWriter, r *http.Request) {

	//If the default project is being saved
	if !ValidProjectLoaded {
		ActiveProject.SaveName = time.Now().Format("20060102150405.000000")
		ActiveProject.Name = "Untitled Project"
		ValidProjectLoaded = true

		//Add to project list
		ProjectList[ActiveProject.SaveName] = ActiveProject.Name //Add to project list
		projectListHasChanged = true                             //Flag list as having changed
		SaveProjectList("./data/projects/projectList.json")
	}

	projectPath := "./data/projects/" + ActiveProject.SaveName + ".json"

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
	//Get the JSON sent
	inputData := &DataTransferText{}

	err := json.NewDecoder(r.Body).Decode(inputData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
		return
	}

	//Extract the project name
	projectName := inputData.Data

	//Create save name and save path
	projectSaveName := time.Now().Format("20060102150405.000000")
	projectPath := "./data/projects/" + projectSaveName + ".json"

	ActiveProject = MakeProject(projectName, projectSaveName)

	//save the new project
	err = SaveProject(ActiveProject, projectPath)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		LogInfo.Println("Created project " + ActiveProject.Name + " (" + projectPath + ")")
		ProjectList[projectSaveName] = projectName //Add to project list
		projectListHasChanged = true               //Flag list as having changed
		SaveProjectList("./data/projects/projectList.json")

		//Flag valid project loaded
		ValidProjectLoaded = true
	}

}

func OverviewProjectHandler(w http.ResponseWriter, r *http.Request) {
	//Parse template
	tmpl, err := template.ParseFiles("data/templates/overviewProject.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl", "data/templates/footer.tmpl")

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

	if projectListHasChanged || (projectListDataTransferSlice == nil) { //If need to update cached data
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

		projectListHasChanged = false
		projectListDataTransferSlice = data.Data
	} else {
		data.Data = projectListDataTransferSlice
	}

	//Encode
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func ImportProjectHandler(w http.ResponseWriter, r *http.Request) {
	//Get the JSON sent
	importedProject := &Project{}

	err := json.NewDecoder(r.Body).Decode(importedProject)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
		return
	}

	//Generate new save name
	projectSaveName := time.Now().Format("20060102150405.000000")
	importedProject.SaveName = projectSaveName

	//Save the project
	projectPath := "./data/projects/" + importedProject.SaveName + ".json"

	ActiveProject = importedProject

	//save the new project
	err = SaveProject(ActiveProject, projectPath)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		LogInfo.Println("Created project " + ActiveProject.Name + " (" + projectPath + ")")
		ProjectList[importedProject.SaveName] = importedProject.Name //Add to project list
		projectListHasChanged = true                                 //Flag list as having changed
		SaveProjectList("./data/projects/projectList.json")
	}

}

func ExportProjectHandler(w http.ResponseWriter, r *http.Request) {
	//Set header
	w.Header().Set("Content-Disposition", "attachment; filename="+ActiveProject.Name+".json")
	w.Header().Set("Content-Type", "application/json")

	//Encode
	err := json.NewEncoder(w).Encode(*ActiveProject)

	//Report errors
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}
