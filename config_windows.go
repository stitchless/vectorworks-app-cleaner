package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func generateConfig(softwareName string, licenseYear string) softwareConfig {
	// define system variables
	winAppData := os.Getenv("APPDATA") + "\\"
	winLocalAppData := os.Getenv("LOCALAPPDATA") + "\\"
	appVersion := convertYearToVersion(licenseYear)

	if softwareName == "Vectorworks" { // Run if Vectorworks was picked
		license := "SOFTWARE\\Nemetschek\\Vectorworks " + appVersion + "\\Registration"
		registry := []string{
			"SOFTWARE\\Nemetschek\\Vectorworks " + appVersion,
			"SOFTWARE\\VectorWorks",
		}
		directories := []string{
			winAppData + softwareName + "\\" + licenseYear,
			winAppData + softwareName + " " + licenseYear + " Installer",
			winAppData + softwareName + " " + licenseYear + " Updater",
			winAppData + "Nemetschek\\" + softwareName + "\\" + licenseYear,
			winAppData + "Nemetschek\\" + softwareName + "\\accounts",
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

		return softwareConfig{
			registry:    registry,
			directories: directories,
			license:     license,
		}

	} else { // Run if Vision was picked
		registry := []string{
			"",
		}
		directories := []string{
			"",
		}
		return softwareConfig{
			registry:    registry,
			directories: directories,
		}
	}
}

// convert Software License year to version number.
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
