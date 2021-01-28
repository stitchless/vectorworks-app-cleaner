package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/gen2brain/dlgs"
	"howett.net/plist"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type LicenseOpts struct {
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

	var plistData LicenseOpts
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
	plistFile, err := os.Open(config.license)
	check(err)
	err = plistFile.Truncate(0)

	plistData := &LicenseOpts{
		license: map[string]string{"NNA User License": newSerial},
	}

	fmt.Println(plistData.license)
	buffer := &bytes.Buffer{}
	encoder := plist.NewEncoder(buffer)

	err = encoder.Encode(plistData.license)
	check(err)

	err = ioutil.WriteFile(config.license, buffer.Bytes(), 0644)
	check(err)

	w := bufio.NewWriter(buffer)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	err = w.Flush()
	check(err)

	refreshPList()
}

func refreshPList() {
	cmd := exec.Command("echo", "-u $USER cfprefsd")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//if err := cmd.Process.Kill(); err != nil {
	//	log.Fatal("failed to kill process: ", err)
	//}
}
