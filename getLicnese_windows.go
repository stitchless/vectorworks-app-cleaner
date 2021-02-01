package main

import (
	"github.com/gen2brain/dlgs"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

// fetchAppInfo fetches the application Year[string], License[string] via getSerial, and a config[map]
// This config will hold license location registries and directories
// In case no entries are found a nil will be returned
func fetchAppInfo(softwareName string) []Version {
	var appYears []string
	var versions []Version
	configs := map[string]softwareConfig{}
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
		version := Version{Year: year, Serial: getSerial(configs[year], softwareName), Config: configs[year]}
		versions = append(versions, version)
	}

	return versions
}


// getSerial will search the registry for any valid serials.
func getSerial(config softwareConfig, softwareName string) string {
	key, _ := registry.OpenKey(registry.CURRENT_USER, config.license , registry.QUERY_VALUE)

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



func inputNewSerial(serial string) string {
	newSerial, _, err := dlgs.Entry("Input New Serial", "Please input a new Series E Serial:", serial)
	if err != nil {
		log.Fatal(err)
	}
	return newSerial
}



func replaceOldSerial(newSerial string, config softwareConfig) {
	// TODO: Clean the input before replacing.
	key, err := registry.OpenKey(registry.CURRENT_USER, config.license, registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = key.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = key.SetStringValue("User Serial Number", newSerial)
	if err != nil {
		log.Fatal(err)
	}
}