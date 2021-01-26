package main

import (
	"fmt"
	"github.com/gen2brain/dlgs"
	"io/ioutil"
	"os"
	"runtime"
)

func main() {
	softwareSelect, closeDiag, err := dlgs.List("Vectorworks, Inc. - App Cleaner", "What software package are you attempting to edit?", []string{"Vectorworks","Vision"})
	if err != nil {
		panic(err)
	}
	if !closeDiag{
		fmt.Println("Closed by user...")
	} else {
		getLicenses(softwareSelect)
	}
}

func getLicenses(softwareSelect string){
	licenses := []string

	if runtime.GOOS == "Darwin" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		files, err := ioutil.ReadDir(homeDir + "/Library/Preferences")

	} else {
		fmt.Println("Running On Windows.")
	}
	runSoftware(softwareSelect)
}

func runSoftware(softwareSelect string){
	if softwareSelect == "Vectorworks" {
		plist := []string{
			"net.nemetschek.vectorworks.license.<appyear>.plist",
			"net.nemetschek.vectorworks.",
			"net.nemetschek.vectorworks.",
			"net.nemetschek.vectorworks.",
			"net.nemetschek.vectorworks.",
			"net.nemetschek.vectorworks.",
			"net.nemetschek.vectorworks.",
		}
		fmt.Println(plist[0])
	} else {
		fmt.Println("Vision Picked")
	}
}