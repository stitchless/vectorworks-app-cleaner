package main

import (
	"github.com/gen2brain/dlgs"
	"log"
	"os"
)

// Data
type softwareConfig struct {
	plist       []string
	registry    []string
	directories []string
	license     string
}

// Home Directory OS dependent
var homeDir, _ = os.UserHomeDir()

func main() {
	// Start by picking between "Vectorworks" and "Vision"
	softwareName, cancelDiag, err := dlgs.List("Vectorworks, Inc. - App Cleaner", "What software package are you attempting to edit?", []string{"Vectorworks", "Vision"})
	if err != nil {
		log.Fatalf("cannot use the dialog as expected: %e", err)
	}

	if !cancelDiag {
		log.Print("Closed by user...")
	}

	license := FindAndChooseLicense(softwareName)   // Find and Choose Versions of software based on license
	config := generateConfig(softwareName, license) // generate proper data for select license version
	chooseAction(config)
}

// Allow user to choose which licence to start working with.
func chooseLicense(softwareName string, licenses []string) string {
	pickedLicense, _, err := dlgs.List("Choose your license", "Please pick from the list of found "+softwareName+" licenses.", licenses)
	if err != nil {
		log.Fatal(err)
	}
	return pickedLicense // return string with 4 digits representing the application license year.
}

func chooseAction(config softwareConfig) {
	items := []string{"Clean Application", "Replace License"}
	choice, _, err := dlgs.List("Chose your action", "What would you like to do?", items)
	if err != nil {
		log.Fatal(err)
	}

	switch choice {
	// Replace old license with new one
	case "Replace License":
		serial := getSerial(config)
		newSerial := inputNewSerial(serial)
		replaceOldSerial(newSerial, config)
	// Removes all properties and files/folders for the given software/version
	case "Clean Application":
		cleanApplication(config)
	}
}
