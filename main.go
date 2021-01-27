package main

import (
	"fmt"
	"github.com/gen2brain/dlgs"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// Data
type workingData struct {
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
	if runtime.GOOS == "darwin" {
		licenses := getMacLicenses(softwareSelect)         // Finds and returns all licenses found for selected software
		license := chooseLicense(softwareSelect, licenses) // Returns a single license version
		data := macData(softwareSelect, license)           // generate proper data for select license version
		fmt.Println(data.plist[1])
	} else {
		//licenses := getWindowsLicenses(softwareSelect)
		// TODO: Follow same workflow as found above and reuse where I can.
		licenses := getWindowsLicenses(softwareSelect)
		license := chooseLicense(softwareSelect, licenses)
		config := winData(softwareSelect, license)
		chooseAction(softwareSelect, config)
		//winGetData(softwareSelect, license)
	}

	// TODO: Should I do something here? ... Perhaps handle above differently so it reaches this point.
}

func getMacLicenses(softwareSelect string) []string {
	var licenses []string

	re := regexp.MustCompile("[0-9]+") // Find all digits for plist file names

	files, err := ioutil.ReadDir(homeDir + "/Library/Preferences") // gets list of all plist file names
	if err != nil {
		log.Fatal(err)
	}

	// returns all license year numbers found in plist file names from the files variable
	for _, f := range files {
		file := strings.Contains(f.Name(), softwareSelect+".license.")

		if file {
			appYear := re.FindAllString(f.Name(), -1)
			licenses = append(licenses, appYear[0])
		}
	}

	return licenses
}

func getWindowsLicenses(softwareSelect string) []string {
	var licenses []string

	re := regexp.MustCompile("[0-9]+")

	folders, err := ioutil.ReadDir(os.Getenv("APPDATA") + "/" + softwareSelect)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range folders {
		appYear := re.FindString(f.Name())
		licenses = append(licenses, appYear)
	}
	return licenses
}

func winGetData(softwareName string) workingData {
	registry := []string{
		"testing",
		"more Testing",
	}

	directories := []string{
		"testing/dir",
		"more testing/dir",
	}

	return workingData{
		registry: registry,
		directories: directories,
	}
}

// Allow user to choose which licence to start working with.
func chooseLicense(softwareName string, licenses []string) string {
	pickedLicense, _, err := dlgs.List("Choose your license", "Please pick from the list of found "+softwareName+" licenses.", licenses)
	if err != nil {
		log.Fatal(err)
	}

	return pickedLicense // return string with 4 digits representing the application license year.
}

func chooseAction(softwareName string, config workingData) {
	items := []string{"Clean Application", "Replace License"}
	choice, _, err := dlgs.List("Chose your action", "What would you like to do?", items)
	if err != nil {
		log.Fatal(err)
	}

	switch choice {
	case "Replace License":
		fmt.Println(config.license)
		currentUser, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		userUIDStr := currentUser.Uid[4:4]
		uid, err := strconv.Atoi(userUIDStr)
		if err != nil {
			log.Fatal(err)
		}
		key, err := registry.OpenKey(registry.CURRENT_USER, config.license, uint32(uid))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(key)
	}
}