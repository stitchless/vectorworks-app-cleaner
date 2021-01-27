package main

import (
	"fmt"
	"github.com/gen2brain/dlgs"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
)

type workingData struct {
	Plist 		[]string
	Directories	[]string
}

var homeDir, _ = os.UserHomeDir()

func main() {
	softwareSelect, closeDiag, err := dlgs.List("Vectorworks, Inc. - App Cleaner", "What software package are you attempting to edit?", []string{"Vectorworks","Vision"})
	if err != nil {
		panic(err)
	}
	if !closeDiag{
		fmt.Println("Closed by user...")
	} else {
		getLicenses(softwareSelect)
	}
}

func getLicenses(softwareSelect string){
	if runtime.GOOS == "darwin" {
		licenses := getMacLicenses(softwareSelect)
		license := chooseLicense(softwareSelect, licenses)
		getData(softwareSelect, license)
	} else {
		fmt.Println("Running On Windows.")
	}
	//runSoftware(softwareSelect)
}
/*

 */
func getMacLicenses(softwareSelect string) []string {
	var licenses []string
	re := regexp.MustCompile("[0-9]+") // Find all digits for plist file names
	files, err := ioutil.ReadDir(homeDir + "/Library/Preferences")
	if err != nil {
		log.Fatal(err)
	}


	for _, f := range files {
		file := strings.Contains(f.Name(), "vectorworks.license.")
		if file {
			appYear := re.FindAllString(f.Name(), -1)
			licenses = append(licenses, appYear[0])
		}
	}
	return licenses
}

func getData(softwareSelect string, licenseYear string) []string {
	if softwareSelect == "Vectorworks" {
		plist := []string{
			"net.nemetschek.vectorworks.license." + licenseYear + ".plist",
			"net.nemetschek.vectorworks." + licenseYear + ".plist",
			"net.nemetschek.vectorworks.spotlightimporter.plist",
			"net.nemetschek.vectorworks.plist",
			"net.nemetschek.vectorworksinstaller.helper.plist",
			"net.nemetschek.vectorworksinstaller.plist",
			"net.vectorworks.vectorworks." + licenseYear + ".plist",
		}
		for _, file := range plist {
			workingData.Plist = append(workingData.plist, file)
		}
			dataDirectories := []string{
			homeDir + "/Library/Application\\ Support/Vectorworks\\ RMCache/rm" + licenseYear,
			homeDir + "/Library/Application\\ Support/Vectorworks\\ Cloud\\ Services",
			homeDir + "/Library/Application\\ Support/Vectorworks/" + licenseYear,
			homeDir + "/Library/Application\\ Support/vectorworks-installer-wrapper",
		}
		for _, name := range plist {
			fmt.Println(name)
		}
		for _, name := range dataDirectories {
			fmt.Println(name)
		}
		return plist
	} else {
		plist := []string{
			"com.qtproject.plist",
			"com.vwvision.Vision" + licenseYear + ".plist",
			"com.yourcompany.Vision.plist",
			"net.vectorworks.Vision.plist",
			"net.vectorworks.vision.license." + licenseYear + ".plist",
		}
		dataDirectories := []string{
			homeDir + "/Library/Application\\ Support/Vision/" + licenseYear,
			homeDir + "/Library/Application\\ Support/VisionUpdater",
			"/Library/Frameworks/QtConcurrent.framework",
			"/Library/Frameworks/QtCore.framework",
			"/Library/Frameworks/QtDBus.framework",
			"/Library/Frameworks/QtGui.framework",
			"/Library/Frameworks/QtNetwork.framework",
			"/Library/Frameworks/QtOpenGL.framework",
			"/Library/Frameworks/QtPlugins",
			"/Library/Frameworks/QtPositioning.framework",
			"/Library/Frameworks/QtPrintSupport.framework",
			"/Library/Frameworks/QtQml.framework",
			"/Library/Frameworks/QtQuick.framework",
			"/Library/Frameworks/QtWebChannel.framework",
			"/Library/Frameworks/QtWebEngine.framework",
			"/Library/Frameworks/QtWebEngineCore.framework",
			"/Library/Frameworks/QtWebEngineWidgets.framework",
			"/Library/Frameworks/QtWidgets.framework",
			"/Library/Frameworks/QtXml.framework",
			"/Library/Frameworks/rpath_manipulator.sh",
			"/Library/Frameworks/setup_qt_frameworks.sh",
		}
		return plist
	}
}

func (data *workingData) constructData(testing []string) []workingData {
	for 
}

func chooseLicense(softwareName string, licenses []string) string {
	pickedLicense, _, err := dlgs.List("Chose your license", "Please pick from the list of found " + softwareName + " licenses.", licenses)
	if err != nil {
		log.Fatal(err)
	}
	return pickedLicense
}