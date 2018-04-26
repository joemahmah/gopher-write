package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/joemahmah/gopher-write/character"
	"encoding/json"
	"html/template"
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
	r.Header.Set("Content-Type","application/json")
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
	
	//Slices to store data
	var names []string
	var uids []int
	
	//Fill slices
	for _,elem := range ActiveProject.Characters{
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