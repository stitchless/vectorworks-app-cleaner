package main

import (
	"bytes"
	"howett.net/plist"
	"io/ioutil"
	"regexp"
	"strings"
)

type LicenseOpts struct {
	serial map[string]string `plist:"NNA User License"`
}

func fetchAppInfo(softwareName string) []Version {
	var appYears []string
	var versions []Version
	re := regexp.MustCompile("[0-9]+")

	files, err := ioutil.ReadDir(homeDir + "/Library/Preferences") // gets list of all plist file names
	check(err)

	// returns all license year numbers found in plist file names from the files variable
	for _, f := range files {
		file := strings.Contains(f.Name(), strings.ToLower(softwareName+".license."))
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
		version := Version{Year: year, Serial: getSerial(softwareName, year)}
		versions = append(versions, version)
	}

	return versions
}

func getSerial(softwareName string, appYear string) string {
	// Determine what software is requested.
	var licenseLocation string
	if softwareName == "Vectorworks" {
		licenseLocation = homeDir + "/Library/Preferences/net.nemetschek.vectorworks.license." + appYear + ".plist"
	} else {
		licenseLocation = homeDir + "/Library/Preferences/net.vectorworks.vision.license." + appYear + ".plist"
	}

	// Read in plist
	plistFile, err := ioutil.ReadFile(licenseLocation)
	buffer := bytes.NewReader(plistFile)
	check(err)

	// parse and return plist serial
	var plistData LicenseOpts
	decoder := plist.NewDecoder(buffer)
	err = decoder.Decode(&plistData.serial)
	check(err)

	return plistData.serial[`NNA User License`]
}

//func inputNewSerial(serial string) string {
//	newSerial, _, err := dlgs.Entry("Input New Serial", "Please input a new Series E Serial:", serial)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return newSerial
//}
//
//func replaceOldSerial(newSerial string, config toBeCleaned) {
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

//func refreshPList() {
//	fmt.Println("Refreshing plist files...")
//	// osascript -e 'do shell script "sudo killall -u $USER cfprefsd" with administrator privileges'
//	cmd := exec.Command(`osascript`, "-s", "h", "-e", `do shell script "sudo killall -u $USER cfprefsd" with administrator privileges`)
//	stderr, err := cmd.StderrPipe()
//	log.SetOutput(os.Stderr)
//	check(err)
//
//	if err = cmd.Start(); err != nil {
//		log.Fatal(err)
//	}
//
//	cmdErrOutput, _ := ioutil.ReadAll(stderr)
//	fmt.Printf("%s\n", cmdErrOutput)
//
//	if err = cmd.Wait(); err != nil {
//		log.Fatal(err)
//	}
//}
