package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

type Client struct {
	name    string
	dirName string
	path    string
}

func main() {
	clientList := []Client{
		{name: "Dev", dirName: "Dev"},
		{name: "Salish", dirName: "Salish"},
		{name: "Blackfoot", dirName: "Blackfoot"},
	}

	var (
		clientName string
	)
	var options []huh.Option[string]

	for _, val := range clientList {
		options = append(options, huh.Option[string]{Value: val.name, Key: val.name})
		// fmt.Println(val.name)
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("chose your client").
				Options(options...).Value(&clientName),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(clientName)
}
