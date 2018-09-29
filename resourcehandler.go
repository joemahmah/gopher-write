package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joemahmah/gopher-write/resources"
	"html/template"
	"net/http"
	"strconv"
)

func NewNoteHandler(w http.ResponseWriter, r *http.Request) {

	//Make new char
	newNote := &resources.Note{}

	//Decode the request
	err := json.NewDecoder(r.Body).Decode(newNote)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		ActiveProject.AddNote(newNote)
	}
}

func NewLinkHandler(w http.ResponseWriter, r *http.Request) {

	//Make new char
	newLink := &resources.Link{}

	//Decode the request
	err := json.NewDecoder(r.Body).Decode(newLink)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		ActiveProject.AddLink(newLink)
	}
}

func ViewNoteHandler(w http.ResponseWriter, r *http.Request) {

	//Get the location UID
	nid, _ := strconv.Atoi(mux.Vars(r)["nid"])

	note, err := ActiveProject.GetNote(nid)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogError.Println(err)
		return
	}

	//Parse template
	tmpl, err := template.ParseFiles("data/templates/viewNote.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl", "data/templates/footer.tmpl")

	//if error parsing template
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
		return
	}

	//serve template
	err = tmpl.Execute(w, note)

	//If error, return code 500
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func ViewLinkHandler(w http.ResponseWriter, r *http.Request) {

	//Get the location UID
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])

	link, err := ActiveProject.GetLink(lid)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogError.Println(err)
		return
	}

	//Parse template
	tmpl, err := template.ParseFiles("data/templates/viewLink.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl", "data/templates/footer.tmpl")

	//if error parsing template
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
		return
	}

	//serve template
	err = tmpl.Execute(w, link)

	//If error, return code 500
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func GetJSONNoteHandler(w http.ResponseWriter, r *http.Request) {

	//Get the char UID
	nid, _ := strconv.Atoi(mux.Vars(r)["nid"])

	note, err := ActiveProject.GetNote(nid)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}

	//Encode and send off
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(note)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func GetJSONLinkHandler(w http.ResponseWriter, r *http.Request) {

	//Get the char UID
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])

	link, err := ActiveProject.GetLink(lid)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}

	//Encode and send off
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(link)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func ListJSONNoteHandler(w http.ResponseWriter, r *http.Request) {

	var notes DataTransferDualStringMonoIntMonoBoolSlice

	//Fill slices
	for _, elem := range ActiveProject.ResNotes {
		notes.Data = append(notes.Data, DualStringMonoIntMonoBool{S1: elem.Title, S2: elem.Description, I: elem.UID, B: elem.Featured})
	}

	//Encode and send off
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(notes)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func ListJSONLinkHandler(w http.ResponseWriter, r *http.Request) {

	var links DataTransferDualStringMonoIntMonoBoolSlice

	//Fill slices
	for _, elem := range ActiveProject.ResLinks {
		links.Data = append(links.Data, DualStringMonoIntMonoBool{S1: elem.Title, S2: elem.Description, I: elem.UID, B: elem.Featured})
	}

	//Encode and send off
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(links)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func OverviewNoteHandler(w http.ResponseWriter, r *http.Request) {

	//Parse template
	tmpl, err := template.ParseFiles("data/templates/overviewNote.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl", "data/templates/footer.tmpl")

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

func OverviewLinkHandler(w http.ResponseWriter, r *http.Request) {

	//Parse template
	tmpl, err := template.ParseFiles("data/templates/overviewLink.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl", "data/templates/footer.tmpl")

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

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {

	nid, _ := strconv.Atoi(mux.Vars(r)["nid"])

	ActiveProject.RemoveNote(nid)

	w.WriteHeader(http.StatusOK)
}

func DeleteLinkHandler(w http.ResponseWriter, r *http.Request) {

	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])

	ActiveProject.RemoveLink(lid)

	w.WriteHeader(http.StatusOK)
}

//////////////////////
// Editing Handlers //
//////////////////////

func EditNoteSetTitleHandler(w http.ResponseWriter, r *http.Request) {
	//Get the uid
	nid, _ := strconv.Atoi(mux.Vars(r)["nid"])

	//Get the location
	note, err := ActiveProject.GetNote(nid)

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

	note.Title = inputData.Data

}

func EditLinkSetTitleHandler(w http.ResponseWriter, r *http.Request) {
	//Get the uid
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])

	//Get the location
	link, err := ActiveProject.GetLink(lid)

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

	link.Title = inputData.Data

}

func EditNoteSetDescriptionHandler(w http.ResponseWriter, r *http.Request) {
	//Get the uid
	nid, _ := strconv.Atoi(mux.Vars(r)["nid"])

	//Get the location
	note, err := ActiveProject.GetNote(nid)

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

	note.Description = inputData.Data

}

func EditLinkSetDescriptionHandler(w http.ResponseWriter, r *http.Request) {
	//Get the uid
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])

	//Get the location
	link, err := ActiveProject.GetLink(lid)

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

	link.Description = inputData.Data

}

func EditNoteSetTextHandler(w http.ResponseWriter, r *http.Request) {
	//Get the uid
	nid, _ := strconv.Atoi(mux.Vars(r)["nid"])

	//Get the location
	note, err := ActiveProject.GetNote(nid)

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

	note.Text = inputData.Data

}

func EditLinkSetURLHandler(w http.ResponseWriter, r *http.Request) {
	//Get the uid
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])

	//Get the location
	link, err := ActiveProject.GetLink(lid)

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

	link.URL = inputData.Data

}

func EditNoteToggleFeaturedHandler(w http.ResponseWriter, r *http.Request) {
	//Get the uid
	nid, _ := strconv.Atoi(mux.Vars(r)["nid"])

	//Get the location
	note, err := ActiveProject.GetLink(nid)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}

	//Send ok
	w.WriteHeader(http.StatusOK)

	note.Featured = !note.Featured

}

func EditLinkToggleFeaturedHandler(w http.ResponseWriter, r *http.Request) {
	//Get the uid
	lid, _ := strconv.Atoi(mux.Vars(r)["lid"])

	//Get the location
	link, err := ActiveProject.GetLink(lid)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}

	//Send ok
	w.WriteHeader(http.StatusOK)

	link.Featured = !link.Featured

}
