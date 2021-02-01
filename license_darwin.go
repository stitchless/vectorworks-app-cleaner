package main

import (
	"bytes"
	"fmt"
	"howett.net/plist"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type LicenseOpts struct {
	license map[string]string `plist:"NNA User License"`
}

func fetchAppInfo(softwareName string) []Version {
	var appYears []string
	var versions []Version
	configs := map[string]softwareConfig{}
	re := regexp.MustCompile("[0-9]+")

	files, _ := ioutil.ReadDir(homeDir + "/Library/Preferences") // gets list of all plist file names

	// returns all license year numbers found in plist file names from the files variable
	for _, f := range files {
		file := strings.Contains(f.Name(), strings.ToLower(softwareName)+".license.")
		if file {
			appYear := re.FindString(f.Name())
			if appYear != "" {
				appYears = append(appYears, appYear)
			}
		}
	}

	if len(appYears) == 0 {
		return nil
	}

	for _, year := range appYears {
		configs[year] = generateConfig(softwareName, year)
		version := Version{Year: year, Serial: getSerial(configs[year]), Config: configs[year]}
		versions = append(versions, version)
	}

	return versions
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

//func inputNewSerial(serial string) string {
//	newSerial, _, err := dlgs.Entry("Input New Serial", "Please input a new Series E Serial:", serial)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return newSerial
//}
//
//func replaceOldSerial(newSerial string, config softwareConfig) {
//	plistFile, err := os.Open(config.license)
//	check(err)
//	err = plistFile.Truncate(0)
//
//	plistData := &LicenseOpts{
//		license: map[string]string{"NNA User License": newSerial},
//	}
//
//	fmt.Println(plistData.license)
//	buffer := &bytes.Buffer{}
//	encoder := plist.NewEncoder(buffer)
//
//	err = encoder.Encode(plistData.license)
//	check(err)
//
//	err = ioutil.WriteFile(config.license, buffer.Bytes(), 0644)
//	check(err)
//
//	w := bufio.NewWriter(buffer)
//	n4, err := w.WriteString("buffered\n")
//	check(err)
//	fmt.Printf("wrote %d bytes\n", n4)
//
//	err = w.Flush()
//	check(err)
//
//	refreshPList()
//}

func refreshPList() {
	fmt.Println("Refreshing plist files...")
	// osascript -e 'do shell script "sudo killall -u $USER cfprefsd" with administrator privileges'
	cmd := exec.Command(`osascript`, "-s", "h", "-e", `do shell script "sudo killall -u $USER cfprefsd" with administrator privileges`)
	stderr, err := cmd.StderrPipe()
	log.SetOutput(os.Stderr)
	check(err)

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}

	cmdErrOutput, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", cmdErrOutput)

	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
