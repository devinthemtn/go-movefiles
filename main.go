package main

import (
	"fmt"
	"log"
	"os"

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
		{name: "Tit", dirName: "Tit"},
		{name: "Nse", dirName: "Nse"},
	}
	// var sourceDir string = "mypath"
	// var destinationDir string = "path"
	//  sourcePath := filepath.Join("source_dir", "my_file.txt")
	// destinationPath := filepath.Join("destination_dir", "my_file.txt")
	var (
		clientName string
	)
	var options []huh.Option[string]

	for _, val := range clientList {
		options = append(options, huh.Option[string]{Value: val.name, Key: val.name})
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
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(cwd)
}
