package main

import (
	"bytes"
	"github.com/gen2brain/dlgs"
	"howett.net/plist"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type softwareOpts struct {
	license map[string]string `plist:"NNA User License"`
}

func FindAndChooseLicense(softwareName string) string {
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

func getSerial(config softwareConfig) string {
	plistFile, err := ioutil.ReadFile(config.license)
	buffer := bytes.NewReader(plistFile)
	check(err)

	var plistData softwareOpts
	decoder := plist.NewDecoder(buffer)
	err = decoder.Decode(&plistData.license)
	check(err)

	return plistData.license[`NNA User License`]
}

func inputNewSerial(serial string) string {
	newSerial, _, err := dlgs.Entry("Input New Serial", "Please input a new Series E Serial:", serial)
	if err != nil {
		log.Fatal(err)
	}
	return newSerial
}

func replaceOldSerial(newSerial string, config softwareConfig) {
	//
	//plistData := &softwareOpts{
	//	//license: `NNA User License`: newSerial,
	//}
	//
	//buffer := &bytes.Buffer{}
	//encoder := plist.NewEncoder(buffer)
	//err := encoder.Encode(plistData)
	//check(err)
	//fmt.Println(buffer.String())
}
