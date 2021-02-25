package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func findInstallationYears(software Software) []string {
	var appdataFolder string
	var years []string

	// Different software has different locations
	switch software {
	case SoftwareVectorworks:
		appdataFolder = os.Getenv("APPDATA") + "/Nemetschek/Vectorworks"
	case SoftwareVision:
		appdataFolder = os.Getenv("APPDATA") + "/Vision"
	default:
		return nil
	}

	folders, _ := ioutil.ReadDir(appdataFolder)

	for _, f := range folders {
		year := regexp.MustCompile("[0-9]+").FindString(f.Name())
		if year != "" {
			years = append(years, year)
		}
	}

	return years
}

func findProperties(installation Installation) []string {
	// define system variables
	version := convertYearToVersion(installation.Year)

	switch installation.Software {
	case SoftwareVectorworks:
		return []string{
			"SOFTWARE\\Nemetschek\\Vectorworks " + version,
			"SOFTWARE\\VectorWorks",
		}
	case SoftwareVision:
		return []string{
			"ESP Vision",
			"SOFTWARE\\VectorWorks\\Vision " + installation.Year,
			"SOFTWARE\\VWVision\\Vision" + installation.Year,
		}
	}

	return nil
}

func findDirectories(installation Installation) []string {
	// define system variables
	winAppData := os.Getenv("APPDATA") + "\\"
	winLocalAppData := os.Getenv("LOCALAPPDATA") + "\\"

	switch installation.Software {
	case SoftwareVectorworks:
		return []string{
			winAppData + installation.Software + "\\" + installation.Year,
			winAppData + installation.Software + " " + installation.Year + " Installer",
			winAppData + installation.Software + " " + installation.Year + " Updater",
			winAppData + "Nemetschek\\" + installation.Software + "\\" + installation.Year,
			winAppData + "Nemetschek\\" + installation.Software + "\\accounts",
			winAppData + "Nemetschek\\" + installation.Software + " RMCache\\rm" + installation.Year,
			winAppData + "Nemetschek\\" + installation.Software + " Web Cache",
			winAppData + "vectorworks-installer",
			winAppData + "vectorworks-updater",
			winAppData + "vectorworks-updater-updater",
			winLocalAppData + "vectorworks-updater-updater",
			winLocalAppData + "Nemetschek",
		}
	case SoftwareVision:
		return []string{
			filepath.Join(winAppData, installation.Software, installation.Year),
			filepath.Join(winLocalAppData, "VisionUpdater"),
		}
	case SoftwareCloudServices:
		return []string{
			winAppData + "vectorworks-cloud-services-beta",
			winAppData + "vectorworks-cloud-services",
			winLocalAppData + "vectorworks-cloud-services-beta-updater",
		}
	}

	return nil
}
