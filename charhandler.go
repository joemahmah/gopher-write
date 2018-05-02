package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/joemahmah/gopher-write/character"
	"github.com/joemahmah/gopher-write/common"
	"encoding/json"
	"html/template"
	"sort"
)

func NewCharHandler(w http.ResponseWriter, r *http.Request) {
	
	//Make new char
	newChar := &character.Character{}
	
	//Decode the request
	err := json.NewDecoder(r.Body).Decode(newChar)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		ActiveProject.AddCharacter(newChar)
		LogInfo.Println("Character " + newChar.Name.PrimaryName + " added to project " + ActiveProject.Name + ".")
	}
	
}

func ViewCharHandler(w http.ResponseWriter, r *http.Request) {
	
	//Get the char UID
	cid, err := strconv.Atoi(mux.Vars(r)["cid"])
	
	//If unable to convert string to int
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
		
		return
	}
	
	//check if exists
	if char, exists := ActiveProject.Characters[cid]; exists {
	
		tmpl := template.New("viewChar.tmpl");
	
		//Add function to check age equality
		tmpl.Funcs(template.FuncMap{
			"ageEq": func(c character.Character) bool{
				return c.Age.BioAge == c.Age.ChronoAge
			},
		})
	
		//Parse template
		tmpl, err := tmpl.ParseFiles("data/templates/viewChar.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl")
		
		//if error parsing template
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			LogError.Println(err)
			
			return
		} 
		
		//serve template
		err = tmpl.Execute(w, char)
		
		//If error, return code 500
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			LogError.Println(err)
		}
		
	} else {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println("Character with uid " + strconv.Itoa(cid) + " does not exist in project " + ActiveProject.Name + ".")
	}
	
}

func GetJSONCharHandler(w http.ResponseWriter, r *http.Request) {
	
	//Get the char UID
	cid, err := strconv.Atoi(mux.Vars(r)["cid"])
	
	//Check if error convering into int
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} 
	
	//Encode and send off
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(ActiveProject.Characters[cid])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} 
}

func EditCharHandler(w http.ResponseWriter, r *http.Request) {
	
	//Make new story
	newChar := &character.Character{}
	
	//Get the story UID
	cuid, err := strconv.Atoi(mux.Vars(r)["cid"])
	
	//check if exists
	if selectedChar, exists := ActiveProject.Characters[cuid]; exists {
		//Decode the request
		err = json.NewDecoder(r.Body).Decode(newChar)
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			LogError.Println(err)
		} else {
			w.WriteHeader(http.StatusOK)
			
			selectedChar.Name = newChar.Name;
			selectedChar.Description = newChar.Description;
			selectedChar.Age = newChar.Age;
			selectedChar.Aliases = newChar.Aliases;
			
			//Log
			LogInfo.Println("Character " + selectedChar.Name.PrimaryName + " of project " + ActiveProject.Name + " was updated.")
		}
		
	} else {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println("Character with uid " + strconv.Itoa(cuid) + " does not exist in project " + ActiveProject.Name + ".")
	}
}

func ListJSONCharHandler(w http.ResponseWriter, r *http.Request) {
	
	var charList DataTransferMonoStringMonoIntSlice
	
	//Fill slices
	for _,elem := range ActiveProject.Characters{
		charList.Data = append(charList.Data, MonoStringMonoInt{S: elem.Name.PrimaryName, I: elem.UID})
	}
	
	//Sort the data (to ensure it will always be the same)
	sort.Slice(charList.Data, func(i int, j int) bool {
		return charList.Data[i].I < charList.Data[j].I
	})
	
	//Encode and send off
	w.Header().Set("Content-Type","application/json")
	err := json.NewEncoder(w).Encode(charList)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} 
}

func OverviewCharHandler(w http.ResponseWriter, r *http.Request) {
	
	//Parse template
	tmpl, err := template.ParseFiles("data/templates/overviewChar.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl")
	
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

func EditCharSetAgeHandler(w http.ResponseWriter, r *http.Request){
	//Get the uid
	cuid, _ := strconv.Atoi(mux.Vars(r)["cid"])

	//Get the character
	character, err := ActiveProject.GetCharacter(cuid)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	inputData := &common.Age{}
	
	err = json.NewDecoder(r.Body).Decode(inputData)
	
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogWarning.Println(err)
		return
	}
	
	//Send ok
	w.WriteHeader(http.StatusOK)
	
	//Set the note
	character.Age = *inputData
	
}

func EditCharSetNameHandler(w http.ResponseWriter, r *http.Request){
	//Get the uid
	cuid, _ := strconv.Atoi(mux.Vars(r)["cid"])

	//Get the character
	character, err := ActiveProject.GetCharacter(cuid)
	
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
	character.Name = *inputData
	
}

func EditCharSetDescriptionHandler(w http.ResponseWriter, r *http.Request){

	//Get the uid
	cuid, _ := strconv.Atoi(mux.Vars(r)["cid"])
	
	//Get the character
	character, err := ActiveProject.GetCharacter(cuid)
	
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
	character.Description = inputData.Data
}

func EditCharSetMotivationHandler(w http.ResponseWriter, r *http.Request){

	//Get the uid
	cuid, _ := strconv.Atoi(mux.Vars(r)["cid"])
	
	//Get the character
	character, err := ActiveProject.GetCharacter(cuid)
	
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
	character.Motivation = inputData.Data
}

func EditCharSetGoalHandler(w http.ResponseWriter, r *http.Request){

	//Get the uid
	cuid, _ := strconv.Atoi(mux.Vars(r)["uid"])
	
	//Get the character
	character, err := ActiveProject.GetCharacter(cuid)
	
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
	character.Goal = inputData.Data
}

func EditCharSetRoleHandler(w http.ResponseWriter, r *http.Request){

	//Get the uid
	cuid, _ := strconv.Atoi(mux.Vars(r)["uid"])
	
	//Get the character
	character, err := ActiveProject.GetCharacter(cuid)
	
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
	character.Role = inputData.Data
}

func EditCharAddAliasHandler(w http.ResponseWriter, r *http.Request){
	//Get the uid
	cuid, _ := strconv.Atoi(mux.Vars(r)["cid"])

	//Get the character
	character, err := ActiveProject.GetCharacter(cuid)
	
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
	character.Aliases = append(character.Aliases, *inputData)
	
}

func EditCharRemoveAliasHandler(w http.ResponseWriter, r *http.Request){
	//Get the uid
	cuid, _ := strconv.Atoi(mux.Vars(r)["cid"])

	//Get the character
	character, err := ActiveProject.GetCharacter(cuid)
	
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
	character.Aliases = append(character.Aliases[:uidToRemove], character.Aliases[uidToRemove+1:]...)
	
}

func AliasListJSONCharHandler(w http.ResponseWriter, r *http.Request) {
	//Get the uid
	cuid, _ := strconv.Atoi(mux.Vars(r)["cid"])

	//Get the character
	character, err := ActiveProject.GetCharacter(cuid)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	var data DataTransferMonoIntMonoNameSlice
	
	//Assign to data
	for index, alias := range character.Aliases {
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

func DeleteCharHandler(w http.ResponseWriter, r *http.Request){
	
	cid, _ := strconv.Atoi(mux.Vars(r)["cid"])

	ActiveProject.RemoveCharacter(cid)

	w.WriteHeader(http.StatusOK)
}
