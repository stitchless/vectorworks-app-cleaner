package main

import (
	"os"
	"path/filepath"
)

// generateConfig generates a series of locations and fill in missing information with teh version
// year then returns them as a slice.
func generateConfig(softwareName string, licenseYear string) toBeCleaned {
	var properties []string
	var directories []string
	// define system variables
	winAppData := os.Getenv("APPDATA") + "\\"
	winLocalAppData := os.Getenv("LOCALAPPDATA") + "\\"
	appVersion := convertYearToVersion(licenseYear)

	switch softwareName {
	case "Vectorworks":
		properties = []string{
			"SOFTWARE\\Nemetschek\\Vectorworks " + appVersion,
			"SOFTWARE\\VectorWorks",
		}
		directories = []string{
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
	case "Vision":
		properties = []string{
			"ESP Vision",
			"SOFTWARE\\VectorWorks\\Vision " + licenseYear,
			"SOFTWARE\\VWVision\\Vision" + licenseYear,
		}
		directories = []string{
			filepath.Join(winAppData, softwareName, licenseYear),
			filepath.Join(winLocalAppData, "VisionUpdater"),
		}
	case "VCS":
		directories = []string{
			winAppData + "vectorworks-cloud-services-beta",
			winAppData + "vectorworks-cloud-services",
			winLocalAppData + "vectorworks-cloud-services-beta-updater",
		}
	}
	return toBeCleaned{
		Properties:  properties,
		Directories: directories,
	}
}
