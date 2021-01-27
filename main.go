package main

import (
	"fmt"
	"github.com/gen2brain/dlgs"
	"log"
	"os"
)

// Data
type softwareConfig struct {
	plist       []string
	registry	[]string
	directories []string
	license		string
}

// Home Directory OS dependent
var homeDir, _ = os.UserHomeDir()

func main() {
	// Start by picking between "Vectorworks" and "Vision"
	softwareSelect, closeDiag, err := dlgs.List("Vectorworks, Inc. - App Cleaner", "What software package are you attempting to edit?", []string{"Vectorworks", "Vision"})
	if err != nil {
		panic(err)
	}

	if !closeDiag {
		fmt.Println("Closed by user...")
	}

	getLicenses(softwareSelect)
}

func getLicenses(softwareSelect string) {
	license := fetchLicense(softwareSelect) // Find and Choose Versions of software based on license

	softwareConfig := constructData(softwareSelect, license) // generate proper data for select license version

	chooseAction(softwareConfig)
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
	case "Replace License":
		replaceLicense(config)
	}
}

func replaceLicense(config softwareConfig) {
	getLicense(config)
	//return licnese
}

//func testing(config softwareConfig) string {
//	currentUser, err := user.Current()
//	if err != nil {
//		log.Fatal(err)
//	}
//	UserID := currentUser.Gid
//	fmt.Println(UserID)
//
//	//userUIDStr := currentUser.Uid[4:4]
//	//uid, err := strconv.Atoi(userUIDStr)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//return string(rune(uid))
//}