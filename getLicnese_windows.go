package main

import (
	"github.com/gen2brain/dlgs"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func fetchAppYears(softwareName string) []Version {
	var appYears []string
	var versions []Version
	configs := map[string]softwareConfig{}

	re := regexp.MustCompile("[0-9]+")

	folders, _ := ioutil.ReadDir(os.Getenv("APPDATA") + "/Nemetschek/" + softwareName)

	for _, f := range folders {
		appYear := re.FindString(f.Name())
		if appYear != "" {
			appYears = append(appYears, appYear)
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
	key, _ := registry.OpenKey(registry.CURRENT_USER, config.license , registry.QUERY_VALUE)

	defer func() {
		_ = key.Close()
	}()

	serial, _, _ := key.GetStringValue("User Serial Number")

	return serial
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