package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/joemahmah/gopher-write/location"
	"github.com/joemahmah/gopher-write/common"
	"encoding/json"
	"html/template"
	"sort"
)

func NewLocationHandler(w http.ResponseWriter, r *http.Request) {
	
	//Make new char
	newLoc := &location.Location{}
	
	//Decode the request
	err := json.NewDecoder(r.Body).Decode(newLoc)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		//Make sublocation map
		newLoc.Sublocations = make(map[int]int)
	
		w.WriteHeader(http.StatusOK)
		ActiveProject.AddLocation(newLoc)
	}
}

func ViewLocationHandler(w http.ResponseWriter, r *http.Request) {
	
	//Get the location UID
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])
	
	loc, err := ActiveProject.GetLocation(lid)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogError.Println(err)
		return
	}
	
	//Parse template
	tmpl, err := template.ParseFiles("data/templates/viewLocation.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl")
	
	//if error parsing template
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
		return
	} 
	
	//serve template
	err = tmpl.Execute(w, loc)
	
	//If error, return code 500
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func GetJSONLocationHandler(w http.ResponseWriter, r *http.Request) {
	
	//Get the char UID
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])
	
	loc, err := ActiveProject.GetLocation(lid)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} 
	
	//Encode and send off
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(loc)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} 
}

func EditLocationHandler(w http.ResponseWriter, r *http.Request) {
	
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])
	
	loc, err := ActiveProject.GetLocation(lid)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogError.Println(err)
		return
	}
	
	newLoc := &location.Location{}
	
	//Decode the request
	err = json.NewDecoder(r.Body).Decode(newLoc)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		
		loc.Name = newLoc.Name;
		loc.Description = newLoc.Description;
	}
}

func ListJSONLocationHandler(w http.ResponseWriter, r *http.Request) {
	
	var locList DataTransferMonoStringDualIntSlice
	
	//Fill slices
	for _,elem := range ActiveProject.Locations{
		locList.Data = append(locList.Data, MonoStringDualInt{S: elem.Name.PrimaryName, I1: elem.UID, I2: elem.Parent})
	}
	
	//Sort the data (to ensure it will always be the same)
	sort.Slice(locList.Data, func(i int, j int) bool {
		return locList.Data[i].I1 < locList.Data[j].I1
	})
	
	//Encode and send off
	w.Header().Set("Content-Type","application/json")
	err := json.NewEncoder(w).Encode(locList)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} 
}

func OverviewLocationHandler(w http.ResponseWriter, r *http.Request) {
	
	//Parse template
	tmpl, err := template.ParseFiles("data/templates/overviewLocation.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl")
	
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

//////////////////////
// Editing Handlers //
//////////////////////

func EditLocationSetNameHandler(w http.ResponseWriter, r *http.Request){
	//Get the uid
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])

	//Get the location
	loc, err := ActiveProject.GetLocation(lid)
	
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
	
	//Set the note
	loc.Name = *inputData
	
}

func EditLocationSetDescriptionHandler(w http.ResponseWriter, r *http.Request){

	//Get the uid
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])
	
	//Get the location
	loc, err := ActiveProject.GetLocation(lid)
	
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
	loc.Description = inputData.Data
}

func EditLocationAddAliasHandler(w http.ResponseWriter, r *http.Request){
	//Get the uid
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])

	//Get the location
	loc, err := ActiveProject.GetLocation(lid)
	
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
	
	//Add the alias
	loc.Aliases = append(loc.Aliases, *inputData)
}

func EditLocationRemoveAliasHandler(w http.ResponseWriter, r *http.Request){
	//Get the uid
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])

	//Get the location
	loc, err := ActiveProject.GetLocation(lid)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	inputData := &DataTransferInt{}
	
	err = json.NewDecoder(r.Body).Decode(inputData)
	
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogWarning.Println(err)
		return
	}
	
	//Send ok
	w.WriteHeader(http.StatusOK)
	
	//Get the uid needed to be removed
	uidToRemove := inputData.Data
	
	//Set the note
	loc.Aliases = append(loc.Aliases[:uidToRemove], loc.Aliases[uidToRemove+1:]...)
	
}


func EditLocationSetParentHandler(w http.ResponseWriter, r *http.Request){
	//Get the uid
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])

	//Get the location
	loc, err := ActiveProject.GetLocation(lid)
	
	//Ensure location exists
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	newParentUID := &DataTransferInt{}
	
	err = json.NewDecoder(r.Body).Decode(newParentUID)
	
	//Check for errors
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogWarning.Println(err)
		return
	}
	
	//Check if the new parent is valid
	newParent, err := ActiveProject.GetLocation(newParentUID.Data);
	if err != nil && newParentUID.Data != -1{
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	//Purge from old parent's sublocation list if needed
	oldParent, err := ActiveProject.GetLocation(loc.Parent)
	if err == nil {
		oldParent.RemoveSublocation(loc.UID)
	}
	
	//Set new parent
	loc.Parent = newParentUID.Data
	
	//Add to new parent sublocation list if needed
	if newParentUID.Data != -1 {
		newParent.AddSublocation(loc.UID)
	}
	
	//Send ok
	w.WriteHeader(http.StatusOK)
}

func SublocationListJSONLocationHandler(w http.ResponseWriter, r *http.Request) {
	//Get the uid
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])

	//Get the character
	loc, err := ActiveProject.GetLocation(lid)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	var data DataTransferMonoStringMonoIntSlice
	
	//Assign to data
	for _, uid := range loc.Sublocations {
		subloc, err := ActiveProject.GetLocation(uid)
		
		//If bad loc reference, skip
		if err != nil {
			continue
		}
		
		data.Data = append(data.Data, MonoStringMonoInt{I: uid, S: subloc.Name.PrimaryName})
	}

	//Encode
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func AliasListJSONLocationHandler(w http.ResponseWriter, r *http.Request) {
	//Get the uid
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])

	//Get the character
	loc, err := ActiveProject.GetLocation(lid)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	var data DataTransferMonoIntMonoNameSlice
	
	//Assign to data
	for index, alias := range loc.Aliases {
		data.Data = append(data.Data, MonoIntMonoName{I: index, Name: alias})
	}

	//Encode
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func DeleteLocationHandler(w http.ResponseWriter, r *http.Request){
	
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])

	ActiveProject.RemoveLocation(lid)

	w.WriteHeader(http.StatusOK)
}
