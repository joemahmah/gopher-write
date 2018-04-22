package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joemahmah/gopher-write/story"
	"strconv"
)

/*
	router.HandleFunc("/move/story/{first:[0-9]{1,9}}/{second:[0-9]{1,9}}", testHandler) //Swap story positions
	router.HandleFunc("/move/chapter/intra/suid:[0-9]{1,9}}/{first:[0-9]{1,9}}/{second:[0-9]{1,9}}", testHandler) //swap chapter positions within a story (first in front of second)
	router.HandleFunc("/move/chapter/inter/fsuid:[0-9]{1,9}}/{first:[0-9]{1,9}}/ssuid:[0-9]{1,9}}/{second:[0-9]{1,9}}", testHandler) //move chapter positions between stories (first in front of second)
*/

func StoryMoveHandler(w http.ResponseWriter, r *http.Request) {
	
	//Print log message
	LogNet.Println("Access " + r.URL.Path + " by "+ r.RemoteAddr)
	
	//Get the story uids
	firstStoryIndex, _ := strconv.Atoi(mux.Vars(r)["first"])
	secondStoryIndex, _ := strconv.Atoi(mux.Vars(r)["second"])
	
	//If story indices are out of bounds
	if secondStoryIndex > len(ActiveProject.Stories) || firstStoryIndex >= len(ActiveProject.Stories) || firstStoryIndex == secondStoryIndex{
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println("Story index out of bounds.")
		return
	}
	
	//If moving story to end
	if len(ActiveProject.Stories) == secondStoryIndex {
		left := ActiveProject.Stories[:firstStoryIndex]
		var right []*story.Story
		
		//edge case, first is at right
		if(firstStoryIndex == len(ActiveProject.Stories) - 1){
			//Do noting since nothing is there
		} else {
			right = ActiveProject.Stories[firstStoryIndex+1:]
		}
		
		//create new slice (eliminates slice BS)
		var newStorySlice []*story.Story;
		newStorySlice = append(newStorySlice, left...)
		newStorySlice = append(newStorySlice, right...)
		newStorySlice = append(newStorySlice, ActiveProject.Stories[firstStoryIndex])
		
		//set stories to new slice
		ActiveProject.Stories = newStorySlice
	} else if (firstStoryIndex < secondStoryIndex){
		left := ActiveProject.Stories[:firstStoryIndex]
		right := ActiveProject.Stories[secondStoryIndex:]
		between := ActiveProject.Stories[firstStoryIndex+1:secondStoryIndex]
		
		//create new slice (eliminates slice BS)
		var newStorySlice []*story.Story;
		newStorySlice = append(newStorySlice, left...)
		newStorySlice = append(newStorySlice, between...)
		newStorySlice = append(newStorySlice, ActiveProject.Stories[firstStoryIndex])
		newStorySlice = append(newStorySlice, right...)
		
		//set stories to new slice
		ActiveProject.Stories = newStorySlice
	} else {
		left := ActiveProject.Stories[:secondStoryIndex]
		right := ActiveProject.Stories[firstStoryIndex:]
		between := ActiveProject.Stories[secondStoryIndex:firstStoryIndex]
		
		//create new slice (eliminates slice BS)
		var newStorySlice []*story.Story;
		newStorySlice = append(newStorySlice, left...)
		newStorySlice = append(newStorySlice, ActiveProject.Stories[firstStoryIndex])
		newStorySlice = append(newStorySlice, between...)
		newStorySlice = append(newStorySlice, right...)
		
		//set stories to new slice
		ActiveProject.Stories = newStorySlice
	}
	
	//Recalculate story UID
	//Note: this is fairly slow, but realistically, there
	//won't be that many stories that the process is too slow...
	for index, story := range ActiveProject.Stories {
		story.UID = index
	}
	
	//Write response
	w.WriteHeader(http.StatusOK)

	//Log Action
	LogInfo.Println("Moved story " + ActiveProject.Stories[secondStoryIndex - 1].Name.PrimaryName + " from project " + ActiveProject.Name + ".")
}

func IntraChapterMoveHandler(w http.ResponseWriter, r *http.Request) {
		
}

func InterChapterMoveHandler(w http.ResponseWriter, r *http.Request) {
		
}
