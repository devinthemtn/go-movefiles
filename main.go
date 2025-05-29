package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"path/filepath"

	"github.com/charmbracelet/huh"
)

type Client struct {
	name    string
	dirName string
	path    string
}

// CopyFile copies a file from src to dst.
func CopyFile(src, dst string) (int64, error) {
	// 1. Open the source file
	sourceFile, err := os.Open(src)
	if err != nil {
		return 0, fmt.Errorf("failed to open source file %s: %w", src, err)
	}
	defer sourceFile.Close() // Ensure the source file is closed when the function exits

	// Get the file info from the source to preserve permissions
	sourceFileInfo, err := sourceFile.Stat()
	if err != nil {
		return 0, fmt.Errorf("failed to get source file info %s: %w", src, err)
	}

	// Create the destination directory if it doesnt exist
	destDir := filepath.Dir(dst)
	if err := os.MkdirAll(destDir, 0755); err != nil { // 0755 provides read/write/execute for owner, read/execute for group and others
		return 0, fmt.Errorf("failed to create destination directory %s: %w", destDir, err)
	}

	// 2. Create the destination file
	destinationFile, err := os.Create(dst)
	if err != nil {
		return 0, fmt.Errorf("failed to create destination file %s: %w", dst, err)
	}
	defer destinationFile.Close() // Ensure the destination file is closed when the function exits

	// Set permissions for the destination file to match the source file
	if err := os.Chmod(dst, sourceFileInfo.Mode()); err != nil {
		return 0, fmt.Errorf("failed to set permissions for destination file %s: %w", dst, err)
	}

	// 3. Copy the content from source to destination
	bytesCopied, err := io.Copy(destinationFile, sourceFile)
	if err != nil {
		return 0, fmt.Errorf("failed to copy file content from %s to %s: %w", src, dst, err)
	}

	return bytesCopied, nil
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
	// destDir := filepath.Join(destinationDir, foundClient.dirName)

	var filePathsSrc [len(fileList)]string
	var filePathsDst [len(fileList)]string
	for idx, val := range fileList {
		filePathsSrc[idx] = filepath.Join(sourceDir, val)
		filePathsDst[idx] = filepath.Join(destinationDir, val)
		// fmt.Println("files: ", val, "idx: ", idx)
	}
	fmt.Println("filepaths: ", filePathsSrc)

	for idx, val := range filePathsSrc {
		// err := os.WriteFile(val, )
		// src := filepath.Join(sourceDir, val)
		// dst := filepath.Join(destDir, val)
		CopyFile(val, filePathsDst[idx])
	}

	// fmt.Println(cwd)
	fmt.Println("sourceDir: ", sourceDir)
	fmt.Println("destinationDir: ", destinationDir)
}
