package main

type DataTransferText struct {
	Data	string
}

type DataTransferInt struct {
	Data	int
}

type DataTransferTextSlice struct {
	Data	[]string
}

type DataTransferIntSlice struct {
	Data	[]int
}

type DualString struct {
	S1	string
	S2	string
}

type DataTransferDualString struct {
	Data	DualString
}

type DataTransferDualStringSlice struct {
	Data	[]DualString
}

