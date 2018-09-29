package main

import (
	"github.com/joemahmah/gopher-write/common"
	"github.com/joemahmah/gopher-write/story"
)

type DataTransferText struct {
	Data string
}

type DataTransferInt struct {
	Data int
}

type DataTransferTextSlice struct {
	Data []string
}

type DataTransferIntSlice struct {
	Data []int
}

type DualString struct {
	S1 string
	S2 string
}

type DataTransferDualString struct {
	Data DualString
}

type DataTransferDualStringSlice struct {
	Data []DualString
}

type DualInt struct {
	I1 int
	I2 int
}

type DataTransferDualInt struct {
	Data DualInt
}

type DataTransferDualIntSlice struct {
	Data []DualInt
}

type DualStringMonoBool struct {
	S1 string
	S2 string
	B  bool
}

type DataTransferDualStringMonoBool struct {
	Data DualStringMonoBool
}

type DataTransferDualStringMonoBoolSlice struct {
	Data []DualStringMonoBool
}

type DualStringMonoIntMonoBool struct {
	S1 string
	S2 string
	I  int
	B  bool
}

type DataTransferDualStringMonoIntMonoBool struct {
	Data DualStringMonoIntMonoBool
}

type DataTransferDualStringMonoIntMonoBoolSlice struct {
	Data []DualStringMonoIntMonoBool
}

type MonoStringMonoInt struct {
	S string
	I int
}

type DataTransferMonoStringMonoInt struct {
	Data MonoStringMonoInt
}

type DataTransferMonoStringMonoIntSlice struct {
	Data []MonoStringMonoInt
}

type MonoStringDualInt struct {
	S  string
	I1 int
	I2 int
}

type DataTransferMonoStringDualInt struct {
	Data MonoStringDualInt
}

type DataTransferMonoStringDualIntSlice struct {
	Data []MonoStringDualInt
}

type MonoIntMonoName struct {
	I    int
	Name common.Name
}

type DataTransferMonoIntMonoName struct {
	Data MonoIntMonoName
}

type DataTransferMonoIntMonoNameSlice struct {
	Data []MonoIntMonoName
}

////////////////////
// Export Structs //
////////////////////

type ExportChapter struct {
	Chapter  story.Chapter
	Sections []story.Section
}

type ExportStory struct {
	Story    story.Story
	Chapters []story.Chapter
	Sections [][]story.Section
}
