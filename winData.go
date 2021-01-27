package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func winData(softwareSelect string, licenseYear string) workingData {
	winAppData := os.Getenv("APPDATA") + "\\"
	winLocalAppData := os.Getenv("LOCALAPPDATA") + "\\"

	appVersion := doTheMath(licenseYear)
	fmt.Println(appVersion)

	if softwareSelect == "Vectorworks" { // Run if Vectorworks was picked
		registry := []string{
			"SOFTWARE\\Nemetschek\\Vectorworks " + appVersion,
		}
		directories := []string{
			winAppData + softwareSelect + "\\" + licenseYear,
			winAppData + softwareSelect + " " + licenseYear + " Installer",
			winAppData + softwareSelect + " " + licenseYear + " Updater",
			winAppData + "Nemetschek\\" + softwareSelect + "\\" + licenseYear,
			winAppData + "Nemetschek\\" + softwareSelect + "\\accounts",
			winAppData + "Nemetschek\\Vectorworks RMCache\\rm" + licenseYear,
			winAppData + "Nemetschek\\Vectorworks Web Cache",
			winAppData + "vectorworks-installer",
			winAppData + "vectorworks-updater",
			winAppData + "vectorworks-updater-updater",
			winAppData + "vectorworks-cloud-services-beta",
			winAppData + "vectorworks-cloud-services",
			winLocalAppData + "vectorworks-updater-updater",
			winLocalAppData + "vectorworks-cloud-services-beta-updater",
			winLocalAppData + "Nemetschek",
		}

		return workingData{
			registry: registry,
			directories: directories,
		}

	} else { // Run if Vision was picked
		registry := []string{
			"",
		}
		directories := []string{
			"",
		}
		return workingData{
			registry: registry,
			directories: directories,
		}
	}
}

func doTheMath(appYear string) string {
	values := strings.Split(appYear, "")[2:4]
	appVersion := values[0] + values[1]
	i, err := strconv.Atoi(appVersion)
	if err != nil {
		log.Fatal(err)
	}
	versionMath := i + 6
	appVersion = strconv.Itoa(versionMath)
	return appVersion
}