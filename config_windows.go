package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// generateConfig generates a series of locations and fill in missing information with teh version
// year then returns them as a slice.
func generateConfig(softwareName string, licenseYear string) softwareConfig {
	// define system variables
	winAppData := os.Getenv("APPDATA") + "\\"
	winLocalAppData := os.Getenv("LOCALAPPDATA") + "\\"
	appVersion := convertYearToVersion(licenseYear)

	switch softwareName {
	case "Vectorworks":
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
			winLocalAppData + "vectorworks-updater-updater",
			winLocalAppData + "Nemetschek",
		}
		return softwareConfig{
			license:     license,
			registry:    registry,
			directories: directories,
		}
	case "Vision":
		license := "SOFTWARE\\VectorWorks\\Vision "+ licenseYear + "\\Registration"
		registry := []string{
			"ESP Vision",
			"SOFTWARE\\VectorWorks\\Vision "+ licenseYear,
			"SOFTWARE\\VWVision\\Vision" + licenseYear,
		}
		return softwareConfig{
			license: license,
			registry: registry,
		}
	case "VCS":
		vcs := []string{
			winAppData + "vectorworks-cloud-services-beta",
			winAppData + "vectorworks-cloud-services",
			winLocalAppData + "vectorworks-cloud-services-beta-updater",
		}
		return softwareConfig{
			vcs: vcs,
		}
	}

	// Should be unreachable
	return softwareConfig{}
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
