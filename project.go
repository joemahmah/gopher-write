package main

import (
	"github.com/joemahmah/gopher-write/story"
	"github.com/joemahmah/gopher-write/location"
	"github.com/joemahmah/gopher-write/character"
	"github.com/joemahmah/gopher-write/resources"
	"encoding/json"
	"os"
	"bufio"
)

var ActiveProject *Project = MakeProject("","")

type Project struct {
		Name			string
		SaveName		string
		Stories			map[int]story.Story
		Chapters		map[int]story.Chapter
		Sections		map[int]story.Section
		Locations		map[int]location.Location
		Characters		map[int]character.Character
		ResLinks		map[int]resources.Link
		ResNotes		map[int]resources.Note
		
		//Next key available
		//Old keys not reused even
		//if old element deleted
		CharacterNext	int
		StoryNext		int
		ChapterNext		int
		SectionNext		int
		LocationNext	int
		ResLinkNext		int
		ResNoteNext		int
}

func MakeProject(name string, savePath string) *Project {
	project := &Project{}
	
	project.Name = name
	project.SaveName = savePath
	
	project.Stories = make(map[int]story.Story)
	project.Chapters = make(map[int]story.Chapter)
	project.Sections = make(map[int]story.Section)
	project.Locations = make(map[int]location.Location)
	project.Characters = make(map[int]character.Character)
	project.ResLinks = make(map[int]resources.Link)
	project.ResNotes = make(map[int]resources.Note)
	
	project.CharacterNext = 0
	project.StoryNext = 0
	project.ChapterNext = 0
	project.SectionNext = 0
	project.LocationNext = 0
	project.ResLinkNext = 0
	project.ResNoteNext = 0
	
	return project
}

func (p *Project) AddCharacter(char character.Character) {
	char.UID = p.CharacterNext //Set UID to next available
	p.CharacterNext++ //Increment next UID available
	p.Characters[char.UID] = char //Add to map 
}

func (p *Project) RemoveCharacter(uid int) {
	delete(p.Characters, uid)
}


func (p *Project) AddStory(story story.Story) {
	story.UID = p.StoryNext //Set UID to next available
	p.StoryNext++ //Increment next UID available
	p.Stories[story.UID] = story //Add to map 
}

func (p *Project) RemoveStory(uid int) {
	delete(p.Stories, uid)
}

func (p *Project) AddChapter(chapter story.Chapter) {
	chapter.UID = p.ChapterNext //Set UID to next available
	p.ChapterNext++ //Increment next UID available
	p.Chapters[chapter.UID] = chapter //Add to map 
}

func (p *Project) RemoveChapter(uid int) {
	delete(p.Chapters, uid)
}

func (p *Project) AddSection(section story.Section) {
	section.UID = p.SectionNext //Set UID to next available
	p.SectionNext++ //Increment next UID available
	p.Sections[section.UID] = section //Add to map 
}

func (p *Project) RemoveSection(uid int) {
	delete(p.Sections, uid)
}

func (p *Project) AddLocation(loc location.Location) {
	loc.UID = p.LocationNext //Set UID to next available
	p.LocationNext++ //Increment next UID available
	p.Locations[loc.UID] = loc //Add to map 
}

func (p *Project) RemoveLocation(uid int) {
	delete(p.Locations, uid)
}

func (p *Project) AddResLink(link resources.Link) {
	link.UID = p.ResLinkNext //Set UID to next available
	p.ResLinkNext++ //Increment next UID available
	p.ResLinks[link.UID] = link //Add to map 
}

func (p *Project) RemoveResLink(uid int) {
	delete(p.ResLinks, uid)
}

func (p *Project) AddResNote(note resources.Note) {
	note.UID = p.ResNoteNext //Set UID to next available
	p.ResNoteNext++ //Increment next UID available
	p.ResNotes[note.UID] = note //Add to map 
}

func (p *Project) RemoveResNote(uid int) {
	delete(p.ResNotes, uid)
}


///////////////////////////////////////////////////////////

func LoadProject(project *Project, path string) error {
	
	//Attempt to open the file at the path given
	projectFile, err := os.Open(path)
	
	//Check for errors
	if err != nil {
		return err
	}
	
	//Close file (defered)
	defer projectFile.Close()
	
	//Buffer the file in case we get a large file
	projectBufferedFile := bufio.NewReader(projectFile)
	
	//Decode the file
	//json.Decoder.Decode returns type error
	return json.NewDecoder(projectBufferedFile).Decode(project)
	
}

func SaveProject(project *Project, path string) error {
	//Attempt to open file at path if exists, otherwise create file
	projectFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0660)
	
	//Check for errors
	if err != nil {
		return err
	}
	
	//Close file (defered)
	defer projectFile.Close()
	
	//Buffer the writing
	projectBufferedWriter := bufio.NewWriter(projectFile)
	
	//Flush the buffer (defered)
	defer projectBufferedWriter.Flush()
	
	//Encode the file
	//returns type error
	return json.NewEncoder(projectBufferedWriter).Encode(project)
	
}
