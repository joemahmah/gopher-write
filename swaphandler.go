package main

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/joemahmah/gopher-write/story"
	"net/http"
	"strconv"
)

func moveItemInSlice(slice []int, targetIndex int, moveBeforeIndex int) ([]int, error) {
	if targetIndex == moveBeforeIndex {
		//Don't do anything, but don't error because
		//the operation is already done...
		return slice, nil
	}

	if targetIndex >= len(slice) {
		return nil, errors.New("targetIndex out of bounds.")
	}

	if moveBeforeIndex > len(slice) {
		return nil, errors.New("moveBeforeIndex out of bounds.")
	}

	if len(slice) == moveBeforeIndex {
		left := slice[:targetIndex]
		var right []int

		//edge case, targetIndex is at right
		if targetIndex == len(slice)-1 {
			//Do noting since nothing is there
		} else {
			right = slice[targetIndex+1:]
		}

		//create new slice (eliminates slice BS)
		var newSlice []int
		newSlice = append(newSlice, left...)
		newSlice = append(newSlice, right...)
		newSlice = append(newSlice, slice[targetIndex])

		//set stories to new slice
		return newSlice, nil
	} else if targetIndex < moveBeforeIndex {
		left := slice[:targetIndex]
		right := slice[moveBeforeIndex:]
		between := slice[targetIndex+1 : moveBeforeIndex]

		//create new slice (eliminates slice BS)
		var newSlice []int
		newSlice = append(newSlice, left...)
		newSlice = append(newSlice, between...)
		newSlice = append(newSlice, slice[targetIndex])
		newSlice = append(newSlice, right...)

		//set stories to new slice
		return newSlice, nil
	} else {
		left := slice[:moveBeforeIndex]
		right := slice[targetIndex+1:]
		between := slice[moveBeforeIndex:targetIndex]

		//create new slice (eliminates slice BS)
		var newSlice []int
		newSlice = append(newSlice, left...)
		newSlice = append(newSlice, slice[targetIndex])
		newSlice = append(newSlice, between...)
		newSlice = append(newSlice, right...)

		//set stories to new slice
		return newSlice, nil
	}
}

func moveItemFromSlice(slice []int, targetIndex int) (int, []int, error) {
	targetAtEnd, err := moveItemInSlice(slice, targetIndex, len(slice))

	if err != nil {
		return 0, nil, err
	}

	return targetAtEnd[len(targetAtEnd)-1], targetAtEnd[:len(targetAtEnd)-1], nil
}

func insertItemIntoSlice(slice []int, item int, moveBeforeIndex int) ([]int, error) {
	appendedSlice := append(slice, item)
	newSlice, err := moveItemInSlice(appendedSlice, len(appendedSlice)-1, moveBeforeIndex)

	if err != nil {
		return nil, err
	}

	return newSlice, nil
}

func StoryMoveHandler(w http.ResponseWriter, r *http.Request) {

	//Get the story uids
	firstStoryIndex, _ := strconv.Atoi(mux.Vars(r)["first"])
	secondStoryIndex, _ := strconv.Atoi(mux.Vars(r)["second"])

	//If story indices are out of bounds
	if secondStoryIndex > len(ActiveProject.Stories) || firstStoryIndex >= len(ActiveProject.Stories) || firstStoryIndex == secondStoryIndex {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println("Story index out of bounds.")
		return
	}

	//If moving story to end
	if len(ActiveProject.Stories) == secondStoryIndex {
		left := ActiveProject.Stories[:firstStoryIndex]
		var right []*story.Story

		//edge case, first is at right
		if firstStoryIndex == len(ActiveProject.Stories)-1 {
			//Do noting since nothing is there
		} else {
			right = ActiveProject.Stories[firstStoryIndex+1:]
		}

		//create new slice (eliminates slice BS)
		var newStorySlice []*story.Story
		newStorySlice = append(newStorySlice, left...)
		newStorySlice = append(newStorySlice, right...)
		newStorySlice = append(newStorySlice, ActiveProject.Stories[firstStoryIndex])

		//set stories to new slice
		ActiveProject.Stories = newStorySlice
	} else if firstStoryIndex < secondStoryIndex {
		left := ActiveProject.Stories[:firstStoryIndex]
		right := ActiveProject.Stories[secondStoryIndex:]
		between := ActiveProject.Stories[firstStoryIndex+1 : secondStoryIndex]

		//create new slice (eliminates slice BS)
		var newStorySlice []*story.Story
		newStorySlice = append(newStorySlice, left...)
		newStorySlice = append(newStorySlice, between...)
		newStorySlice = append(newStorySlice, ActiveProject.Stories[firstStoryIndex])
		newStorySlice = append(newStorySlice, right...)

		//set stories to new slice
		ActiveProject.Stories = newStorySlice
	} else {
		left := ActiveProject.Stories[:secondStoryIndex]
		right := ActiveProject.Stories[firstStoryIndex+1:]
		between := ActiveProject.Stories[secondStoryIndex:firstStoryIndex]

		//create new slice (eliminates slice BS)
		var newStorySlice []*story.Story
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
}

func IntraChapterMoveHandler(w http.ResponseWriter, r *http.Request) {

	//Get the story uids
	firstChapterIndex, _ := strconv.Atoi(mux.Vars(r)["first"])
	secondChapterIndex, _ := strconv.Atoi(mux.Vars(r)["second"])
	suid, _ := strconv.Atoi(mux.Vars(r)["fsuid"])

	if len(ActiveProject.Stories) <= suid {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println("Story index out of bounds.")
		return
	}

	newChapterSlice, err := moveItemInSlice(ActiveProject.Stories[suid].Chapters, firstChapterIndex, secondChapterIndex)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
		return
	}

	ActiveProject.Stories[suid].Chapters = newChapterSlice

	w.WriteHeader(http.StatusOK)
}

func InterChapterMoveHandler(w http.ResponseWriter, r *http.Request) {

	//Get the story uids
	firstChapterIndex, _ := strconv.Atoi(mux.Vars(r)["first"])
	secondChapterIndex, _ := strconv.Atoi(mux.Vars(r)["second"])
	fsuid, _ := strconv.Atoi(mux.Vars(r)["fsuid"])
	ssuid, _ := strconv.Atoi(mux.Vars(r)["ssuid"])

	//Check if this is actually an intra story move
	if fsuid == ssuid {
		IntraChapterMoveHandler(w, r)
		return
	}

	if len(ActiveProject.Stories) <= fsuid || len(ActiveProject.Stories) <= ssuid {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println("Story index out of bounds.")
		return
	}

	//Make new first chapter slice and get value (remove and get)
	value, newFirstChapterSlice, err := moveItemFromSlice(ActiveProject.Stories[fsuid].Chapters, firstChapterIndex)

	//Check for errors
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err)
		return
	}

	//Make new second chapter slice (insert and move)
	newSecondChapterSlice, err2 := insertItemIntoSlice(ActiveProject.Stories[ssuid].Chapters, value, secondChapterIndex)

	//Check for errors
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println(err2)
		return
	}

	ActiveProject.Stories[fsuid].Chapters = newFirstChapterSlice
	ActiveProject.Stories[ssuid].Chapters = newSecondChapterSlice

	w.WriteHeader(http.StatusOK)
}

func IntraSectionMoveHandler(w http.ResponseWriter, r *http.Request) {

	//Get the story uids
	firstSectionIndex, _ := strconv.Atoi(mux.Vars(r)["first"])
	secondSectionIndex, _ := strconv.Atoi(mux.Vars(r)["second"])
	suid, _ := strconv.Atoi(mux.Vars(r)["fsuid"])
	cuidRel, _ := strconv.Atoi(mux.Vars(r)["fcuid"])

	//Process relative uids
	cuid := ActiveProject.Stories[suid].Chapters[cuidRel]

	if len(ActiveProject.Stories) <= suid {
		w.WriteHeader(http.StatusInternalServerError)
		LogError.Println("Story index out of bounds.")
		return
	}

	if selectedChapter, exists := ActiveProject.Chapters[cuid]; exists {

		newChapterSlice, err := moveItemInSlice(selectedChapter.Sections, firstSectionIndex, secondSectionIndex)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			LogError.Println(err)
			return
		}

		selectedChapter.Sections = newChapterSlice

		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
		LogWarning.Println("Chapter with uid " + strconv.Itoa(cuid) + " does not exist in project " + ActiveProject.Name + ".")
	}
}
