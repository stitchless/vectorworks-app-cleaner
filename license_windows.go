package main

import (
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// fetchAppInfo fetches the application Year[string], License[string] via getSerial, and a config[map]
// This config will hold license location registries and directories
// In case no entries are found a nil will be returned
func fetchAppInfo(softwareName string) []Version {
	var appYears []string
	var versions []Version
	configs := map[string]toBeCleaned{}
	re := regexp.MustCompile("[0-9]+")
	var softwareFolder string

	// Different software has different locations
	if softwareName == "Vectorworks" {
		softwareFolder = os.Getenv("APPDATA") + "/Nemetschek/Vectorworks"
	} else if softwareName == "Vision" {
		softwareFolder = os.Getenv("APPDATA") + "/Vision"
	}
	folders, _ := ioutil.ReadDir(softwareFolder)

	for _, f := range folders {
		appYear := re.FindString(f.Name())
		if appYear != "" {
			appYears = append(appYears, appYear)
		}
	}

	// In case no versions are found, we stop from proceeding and return nil
	if len(appYears) == 0 {
		return nil
	}

	// Attach configs, versions, and app years all into on object then return that object
	for _, year := range appYears {
		configs[year] = generateConfig(softwareName, year)
		version := Version{Year: year, Serial: getSerial(softwareName, year)}
		versions = append(versions, version)
	}

	return versions
}

// getSerial will search the registry for any valid serials.
func getSerial(softwareName string, appYear string) string {
	var licenseLocation string

	appVersion := convertYearToVersion(appYear)
	if softwareName == "Vectorworks" {
		licenseLocation = "SOFTWARE\\Nemetschek\\Vectorworks " + appVersion + "\\Registration"
	} else {
		licenseLocation = "SOFTWARE\\VectorWorks\\Vision " + appYear + "\\Registration"
	}

	// Get the Registry Key
	key, _ := registry.OpenKey(registry.CURRENT_USER, licenseLocation, registry.QUERY_VALUE)
	defer func() {
		_ = key.Close()
	}()
	if softwareName == "Vectorworks" {
		serial, _, _ := key.GetStringValue("User Serial Number")
		return serial
	} else if softwareName == "Vision" {
		serial, _, _ := key.GetStringValue("")
		return serial
	}
	return ""
}

// convertYearToVersion returns a version number as opposed to a version year.
// This is need in the case of a registry due to application versions being used
// in place of version years
func convertYearToVersion(appYear string) string {
	values := strings.Split(appYear, "")[2:4]
	appVersion := values[0] + values[1]
	i, err := strconv.Atoi(appVersion)
	if err != nil {
		log.Fatal(err)
	}
	versionMath := i + 5
	appVersion = strconv.Itoa(versionMath)
	return appVersion
}

//func inputNewSerial(serial string) string {
//	newSerial, _, err := dlgs.Entry("Input New Serial", "Please input a new Series E Serial:", serial)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return newSerial
//}
//
//
//
//func replaceOldSerial(newSerial string, config toBeCleaned) {
//	// TODO: Clean the input before replacing.
//	key, err := registry.OpenKey(registry.CURRENT_USER, config.license, registry.SET_VALUE)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer func() {
//		err = key.Close()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}()
//
//	err = key.SetStringValue("User Serial Number", newSerial)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
