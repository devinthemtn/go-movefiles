package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"path/filepath"
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
		{name: "Blackfoot", dirName: "BF"},
		{name: "Colville-Tit", dirName: "Colville-tit"},
		{name: "Colville-nse", dirName: "Colville-nse"},
		{name: "Colville-nxa", dirName: "Colville-nxa"},
	}
	fileList := [...]string{"_colors.scss", "clientConfig.js", "clientStyleConfig.scss", "configData.js"}
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error: ", err)
	}
	sourceDir := filepath.Join(cwd, "/src/Config")
	destinationDir := filepath.Join(cwd, "/src/Config")
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

	err1 := form.Run()
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(clientName)
	var foundClient *Client
	for i := range clientList {
		if clientList[i].name == clientName {
			foundClient = &clientList[i]
			break
		}
	}
	if foundClient == nil {
		fmt.Println("err: client not found")
	}
	sourceDir = filepath.Join(sourceDir, foundClient.dirName)

	var filePaths [4]string
	for idx, val := range fileList {
		filePaths[idx] = filepath.Join(sourceDir, val)
		// fmt.Println("files: ", val, "idx: ", idx)
	}
	fmt.Println("filepaths: ", filePaths)

	// fmt.Println(cwd)
	fmt.Println("sourceDir: ", sourceDir)
	fmt.Println("destinationDir: ", destinationDir)
}
