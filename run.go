package main

import (
	"fmt"
	"os"
	"strconv"
)

type Datapeserta struct {
	Name        string `json:"name"`
	KodePeserta string `json:"kode_peserta"`
}

func main() {
	argsWithProg := os.Args
	var data = []Datapeserta{
		{
			Name:        "Corry",
			KodePeserta: "GLNG023ONL001",
		},
		{
			Name:        "Hamdan",
			KodePeserta: "GLNG023ONL002",
		},
		{
			Name:        "Ken",
			KodePeserta: "GLNG023ONL003",
		},
		{
			Name:        "Prasetya",
			KodePeserta: "GLNG023ONL004",
		},
		{
			Name:        "Yoga",
			KodePeserta: "GLNG023ONL005",
		},
		{
			Name:        "Ajat",
			KodePeserta: "GLNG023ONL006",
		},
	}
	intVar, _ := strconv.Atoi(argsWithProg[1])
	if intVar-1 < len(data) {
		fmt.Println(data[intVar-1])
	} else {
		fmt.Println("data tidak ditemukan")
	}
}
