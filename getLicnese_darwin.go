package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func fetchLicense(softwareName string) string {
	var licenses []string

	re := regexp.MustCompile("[0-9]+") // Find all digits for plist file names

	files, err := ioutil.ReadDir(homeDir + "/Library/Preferences") // gets list of all plist file names
	if err != nil {
		log.Fatal(err)
	}

	// returns all license year numbers found in plist file names from the files variable
	for _, f := range files {
		file := strings.Contains(f.Name(), strings.ToLower(softwareName)+".license.")

		if file {
			appYear := re.FindAllString(f.Name(), -1)
			licenses = append(licenses, appYear[0])
		}
	}
	return chooseLicense(softwareName, licenses)
}
