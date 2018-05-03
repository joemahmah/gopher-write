package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/joemahmah/gopher-write/story"
	"encoding/json"
	"html/template"
)

func NewStoryHandler(w http.ResponseWriter, r *http.Request) {
	
	//Make new story
	newStory := &story.Story{}
	
	//Decode the request
	err := json.NewDecoder(r.Body).Decode(newStory)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)
		ActiveProject.AddStory(newStory)
		LogInfo.Println("Story " + newStory.Name.PrimaryName + " added to project " + ActiveProject.Name + ".")
	}
	
}


func NewChapterHandler(w http.ResponseWriter, r *http.Request) {

	//Make new chapter
	newChapter := &story.Chapter{}

	//Get the story UID
	suid, err := strconv.Atoi(mux.Vars(r)["storyuid"])

	selectedStory, err := ActiveProject.GetStory(suid)

	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}

	//Decode the request
	err = json.NewDecoder(r.Body).Decode(newChapter)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)

		//Add chapter to project to assign uid
		ActiveProject.AddChapter(newChapter)

		//Add chapter to story
		selectedStory.Chapters = append(selectedStory.Chapters, newChapter.UID)

		//Log
		LogInfo.Println("Chapter " + newChapter.Name.PrimaryName + " added to project " + ActiveProject.Name + ".")
	}

}


func NewSectionHandler(w http.ResponseWriter, r *http.Request) {

	//Make new section
	newSection := &story.Section{}

	//Get the selectedChapter UID
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["chapteruid"])

	selectedChapter, err := ActiveProject.GetChapter(suid, cuidRel)

	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}

	//Decode the request
	err = json.NewDecoder(r.Body).Decode(newSection)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)

		//Add section to project to assign uid
		ActiveProject.AddSection(newSection)

		//Add chapter to story
		selectedChapter.Sections = append(selectedChapter.Sections, newSection.UID)

		//Log
		LogInfo.Println("Section " + newSection.Name.PrimaryName + " added to project " + ActiveProject.Name + ".")
	}

}

func ViewStoryHandler(w http.ResponseWriter, r *http.Request) {

	//Get the selectedStory	UID
	suid, err := strconv.Atoi(mux.Vars(r)["storyuid"])

	selectedStory, err := ActiveProject.GetStory(suid)

	//checks if exists
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}

	//Parse template
	tmpl, err := template.ParseFiles("data/templates/viewStory.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl")

	//if error parsing template
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)

		return
	}

	//serve template
	err = tmpl.Execute(w, selectedStory)

	//If error, return code 500
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}


}

func ViewChapterHandler(w http.ResponseWriter, r *http.Request) {

	//Get the selected chapter UID
	suid, err := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, err := strconv.Atoi(mux.Vars(r)["chapteruid"])

	selectedChapter, err := ActiveProject.GetChapter(suid, cuidRel)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogError.Println(err)
		return
	}

	//Parse template
	tmpl, err := template.ParseFiles("data/templates/viewChapter.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl")

	//if error parsing template
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
		return
	}

	//serve template
	err = tmpl.Execute(w, struct{UIDS [2]int; Chapter *story.Chapter}{UIDS: [2]int{suid,cuidRel}, Chapter: selectedChapter})

	//If error, return code 500
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func ViewSectionHandler(w http.ResponseWriter, r *http.Request) {

	//Get the selected chapter UID
	suid, err := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, err := strconv.Atoi(mux.Vars(r)["chapteruid"])
	seuidRel, err := strconv.Atoi(mux.Vars(r)["sectionuid"])

	selectedSection, err := ActiveProject.GetSection(suid, cuidRel, seuidRel)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogError.Println(err)
		return
	}

	//Parse template
	tmpl, err := template.ParseFiles("data/templates/viewSection.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl")

	//if error parsing template
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)

		return
	}

	//serve template
	err = tmpl.Execute(w, struct{UIDS [3]int; Section *story.Section}{UIDS: [3]int{suid,cuidRel,seuidRel}, Section: selectedSection})

	//If error, return code 500
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func GetJSONStoryHandler(w http.ResponseWriter, r *http.Request) {
	
	//Get the char UID
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])

	selectedStory, err := ActiveProject.GetStory(suid)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}

	//Encode and send off
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(selectedStory)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func GetJSONChapterHandler(w http.ResponseWriter, r *http.Request) {

	//Get the char UID
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["chapteruid"])

	selectedChapter, err := ActiveProject.GetChapter(suid, cuidRel)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}

	//Encode and send off
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(selectedChapter)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func GetJSONSectionHandler(w http.ResponseWriter, r *http.Request) {

	//Get the char UID
	suid, err := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, err := strconv.Atoi(mux.Vars(r)["chapteruid"])
	seuidRel, err := strconv.Atoi(mux.Vars(r)["sectionuid"])

	selectedSection, err := ActiveProject.GetSection(suid, cuidRel, seuidRel)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}

	//Encode and send off
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(selectedSection)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func EditStoryHandler(w http.ResponseWriter, r *http.Request) {

	//Make new story
	newStory := &story.Story{}

	//Get the story UID
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])

	selectedStory, err := ActiveProject.GetStory(suid)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}

	//Decode the request
	err = json.NewDecoder(r.Body).Decode(newStory)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)

		selectedStory.Name = newStory.Name;
		selectedStory.Status = newStory.Status;
	}
}

func EditChapterHandler(w http.ResponseWriter, r *http.Request) {

	//Make new chapter
	newChapter := &story.Chapter{}

	//Get the selectedChapter UID
	suid, err := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, err := strconv.Atoi(mux.Vars(r)["chapteruid"])

	selectedChapter, err := ActiveProject.GetChapter(suid, cuidRel)

	//Decode the request
	err = json.NewDecoder(r.Body).Decode(newChapter)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)

		selectedChapter.Name = newChapter.Name;
		selectedChapter.Status = newChapter.Status;

		//Log
		LogInfo.Println("Chapter " + selectedChapter.Name.PrimaryName + " of project " + ActiveProject.Name + " was updated.")
	}
}

func EditSectionHandler(w http.ResponseWriter, r *http.Request) {

	//Make new section
	newSection := &story.Section{}
	
	//Get the selectedChapter UID
	suid, err := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, err := strconv.Atoi(mux.Vars(r)["chapteruid"])
	seuidRel, err := strconv.Atoi(mux.Vars(r)["sectionuid"])
	
	//Calc project CUID from relative CUID
	cuid := ActiveProject.Stories[suid].Chapters[cuidRel]
	seuid := ActiveProject.Chapters[cuid].Sections[seuidRel]
	
	//check if exists
	if selectedSection, exists := ActiveProject.Sections[seuid]; exists {
	
		//Decode the request
		err = json.NewDecoder(r.Body).Decode(newSection)
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			LogError.Println(err)
		} else {
			w.WriteHeader(http.StatusOK)
			
			selectedSection.Name = newSection.Name
			selectedSection.Status = newSection.Status
			selectedSection.Text = newSection.Text
			selectedSection.Characters = newSection.Characters
			selectedSection.Locations = newSection.Locations
			selectedSection.Note = newSection.Note

			//Log
			LogInfo.Println("Section " + selectedSection.Name.PrimaryName + " of project " + ActiveProject.Name + " was updated.")
		}
		
	} else {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println("Section with uid " + strconv.Itoa(seuid) + " does not exist in project " + ActiveProject.Name + ".")
	}
}

func ListJSONStoryHandler(w http.ResponseWriter, r *http.Request) {
	
	//Slices to store data
	var names []string
	var uids []int
	
	//Fill slices
	for _,elem := range ActiveProject.Stories{
		names = append(names, elem.Name.PrimaryName)
		uids = append(uids, elem.UID)
	}
	
	//Encode and send off
	w.Header().Set("Content-Type","application/json")
	err := json.NewEncoder(w).Encode(struct{Names []string; UIDS []int}{Names: names, UIDS: uids})
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} 
}

func ListJSONChapterHandler(w http.ResponseWriter, r *http.Request) {
	
	//Get the selectedStory	UID
	suid, err := strconv.Atoi(mux.Vars(r)["storyuid"])
	
	//Slices to store data
	var names []string
	var uids []int
	
	//Fill slices
	for index,_ := range ActiveProject.Stories[suid].Chapters{
		cuid := ActiveProject.Stories[suid].Chapters[index] //get the project CUID of the chapter
		names = append(names, ActiveProject.Chapters[cuid].Name.PrimaryName)
		uids = append(uids, index)
	}
	
	//Encode and send off
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(struct{Names []string; UIDS []int}{Names: names, UIDS: uids})
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} 
}

func ListJSONSectionHandler(w http.ResponseWriter, r *http.Request) {
	
	//Get the selected Chapter UID
	suid, err := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, err := strconv.Atoi(mux.Vars(r)["chapteruid"])

	//Calc project CUID from relative CUID
	cuid := ActiveProject.Stories[suid].Chapters[cuidRel]
	
	//Slices to store data
	var names []string
	var uids []int
	
	//Fill slices
	for index,_ := range ActiveProject.Chapters[cuid].Sections{
		secuid := ActiveProject.Chapters[cuid].Sections[index] //get the project CUID of the chapter
		names = append(names, ActiveProject.Sections[secuid].Name.PrimaryName)
		uids = append(uids, index)
	}
	
	//Encode and send off
	w.Header().Set("Content-Type","application/json")
	err = json.NewEncoder(w).Encode(struct{Names []string; UIDS []int}{Names: names, UIDS: uids})
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
	
}

func OverviewStoryHandler(w http.ResponseWriter, r *http.Request) {
	
	//Parse template
	tmpl, err := template.ParseFiles("data/templates/overviewStories.tmpl", "data/templates/style.tmpl", "data/templates/header.tmpl", "data/templates/js.tmpl")
	
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

func DeleteStoryHandler(w http.ResponseWriter, r *http.Request) {
	
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])

	ActiveProject.RemoveStory(suid)

	w.WriteHeader(http.StatusOK)
}

func DeleteChapterHandler(w http.ResponseWriter, r *http.Request){
	
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["chapteruid"])

	ActiveProject.RemoveChapterRel(suid, cuidRel)

	w.WriteHeader(http.StatusOK)
}

func DeleteSectionHandler(w http.ResponseWriter, r *http.Request){
	
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["chapteruid"])
	seuidRel, _ := strconv.Atoi(mux.Vars(r)["sectionuid"])

	ActiveProject.RemoveSectionRel(suid, cuidRel, seuidRel)

	w.WriteHeader(http.StatusOK)
}


func ExportStoryHandler(w http.ResponseWriter, r *http.Request) {
	
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	
	storyToExport, err := ActiveProject.GetStory(suid);

	//If error, return 404
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	//Actual export var
	var storyData ExportStory
	
	//Assign story
	storyData.Story = *storyToExport
	
	//Assign chapters and sections
	for _, chapterUID := range storyToExport.Chapters {
		//TODO: add getter to get section by absolute uid
		//TODO: add error checking
		storyData.Chapters = append(storyData.Chapters, *ActiveProject.Chapters[chapterUID])
		
		//Assign sections
		var sects []story.Section
		for _, sectionUID := range ActiveProject.Chapters[chapterUID].Sections{
			sects = append(sects, *ActiveProject.Sections[sectionUID])
		}
		storyData.Sections = append(storyData.Sections, sects)
	}
	
	//Set header
	w.Header().Set("Content-Disposition", "attachment; filename=" + storyToExport.Name.PrimaryName + ".json")
	w.Header().Set("Content-Type","application/json")
	
	//Encode
	err = json.NewEncoder(w).Encode(storyData)

	//Report errors
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func ExportChapterHandler(w http.ResponseWriter, r *http.Request) {
	
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["chapteruid"])
	
	chapter, err := ActiveProject.GetChapter(suid, cuidRel);

	//If error, return 404
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	//Actual export var
	var chapterData ExportChapter
	
	//Assign chapter
	chapterData.Chapter = *chapter
	
	//Assign sections
	for _, sectionUID := range chapter.Sections {
		//TODO: add getter to get section by absolute uid
		//TODO: add error checking
		chapterData.Sections = append(chapterData.Sections, *ActiveProject.Sections[sectionUID])
	}
	
	//Set header
	w.Header().Set("Content-Disposition", "attachment; filename=" + chapter.Name.PrimaryName + ".json")
	w.Header().Set("Content-Type","application/json")
	
	//Encode
	err = json.NewEncoder(w).Encode(chapterData)

	//Report errors
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func ExportSectionHandler(w http.ResponseWriter, r *http.Request) {
	
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["chapteruid"])
	seuidRel, _ := strconv.Atoi(mux.Vars(r)["sectionuid"])
	
	section, err := ActiveProject.GetSection(suid, cuidRel, seuidRel);

	//If error, return 404
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println(err)
		return
	}
	
	//Set header
	w.Header().Set("Content-Disposition", "attachment; filename=" + section.Name.PrimaryName + ".json")
	w.Header().Set("Content-Type","application/json")
	
	//Encode
	err = json.NewEncoder(w).Encode(section)

	//Report errors
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	}
}

func ImportChapterHandler(w http.ResponseWriter, r *http.Request) {

	//Get the uids of the story and chapter
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	bindchar, _ := strconv.ParseBool(mux.Vars(r)["bindchar"])
	bindloc, _ := strconv.ParseBool(mux.Vars(r)["bindloc"])

	selectedStory, err := ActiveProject.GetStory(suid)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogError.Println(err)
		return
	}
	
	//Get the JSON sent
	importedData := &ExportChapter{}
	
	err = json.NewDecoder(r.Body).Decode(importedData)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)

		//Add Chapter to project to assigun uid
		ActiveProject.AddChapter(&importedData.Chapter)
		
		//Add chapter to story
		selectedStory.Chapters = append(selectedStory.Chapters, importedData.Chapter.UID)
		
		//Remove section references (UIDS not valid)
		importedData.Chapter.Sections = make([]int, 0)
		
		//Add sections
		for _, section := range importedData.Sections {
			//Assign section to new section otherwise the 
			//previous sections get overwritten.
			newSect := section
			
			//unbind chars and locations if needed
			if !bindchar {
				newSect.Characters = make([]int, 0)
			}
			if !bindloc {
				newSect.Locations = make([]int, 0)
			}
			
			//Add section to project to assign uid
			ActiveProject.AddSection(&newSect)
			LogInfo.Println(ActiveProject.Sections[4])
	
			//Add section to chapter
			importedData.Chapter.Sections = append(importedData.Chapter.Sections, newSect.UID)
		}
		
		
		//Log
		LogInfo.Println("Chapter " + importedData.Chapter.Name.PrimaryName + " imported to project " + ActiveProject.Name + ".")
	}
	
}

func ImportSectionHandler(w http.ResponseWriter, r *http.Request) {

	//Get the uids of the story and chapter
	suid, _ := strconv.Atoi(mux.Vars(r)["storyuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["chapteruid"])
	bindchar, _ := strconv.ParseBool(mux.Vars(r)["bindchar"])
	bindloc, _ := strconv.ParseBool(mux.Vars(r)["bindloc"])

	selectedChapter, err := ActiveProject.GetChapter(suid, cuidRel)
	
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		LogError.Println(err)
		return
	}
	
	//Get the JSON sent
	importedSection := &story.Section{}
	
	err = json.NewDecoder(r.Body).Decode(importedSection)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
	} else {
		w.WriteHeader(http.StatusOK)

		//unbind chars and locations if needed
		if !bindchar {
			importedSection.Characters = make([]int, 0)
		}
		if !bindloc {
			importedSection.Locations = make([]int, 0)
		}
		
		//Add section to project to assign uid
		ActiveProject.AddSection(importedSection)

		//Add section to chapter
		selectedChapter.Sections = append(selectedChapter.Sections, importedSection.UID)

		//Log
		LogInfo.Println("Section " + importedSection.Name.PrimaryName + " imported to project " + ActiveProject.Name + ".")
	}
	
}