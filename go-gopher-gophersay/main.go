package main

import (
	"embed"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//go:embed gophers
var embedGopherFiles embed.FS

func main() {
	// display usage/help message
	if len(os.Args) == 1 || len(os.Args) == 2 && os.Args[1] == "-h" || len(os.Args) == 2 && os.Args[1] == "--help" {
		uasge := "Gophersay is inspired by Cowsay program.\nGopherSay allow you to display a message said by a cute ranom Gopher.\n\n Usage:\n gophersasy Message\n\n Example:\n gophersay hello Gopher lovers"
		fmt.Println(uasge)
		return
	} else if len(os.Args) > 1 {
		message := strings.Join(os.Args[1:], " ")
		nbChar := len(message)

		line := " "
		for i := 0; i < nbChar; i++ {
			line += "-"
		}

		fmt.Println(line)
		fmt.Println("< " + message + " >")
		fmt.Println("        \\")
		fmt.Println("         \\")
		// Generate a random integer depending on get the number of ascii files
		rand.NewSource(time.Now().UnixNano())
		randInt := rand.Intn(getNbOfGopherFiles() - 1)
		// Display random gopher ASCII embed files
		fileData, err := embedGopherFiles.ReadFile("gophers/gopher" + strconv.Itoa(randInt) + ".txt")
		if err != nil {
			log.Fatal("Error during read gopher ascii file", err)
		}
		fmt.Println(string(fileData))
	}
}

func getNbOfGopherFiles() int {

	files, err := embedGopherFiles.ReadDir("gophers")
	if err != nil {
		log.Fatal("Error during reading gophers folder", err)
	}

	nbOfFiles := 0
	for _, _ = range files {
		nbOfFiles++
	}

	return nbOfFiles
}
