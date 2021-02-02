package main

func generateConfig(softwareName string, licenseYear string) toBeCleaned {
	var properties[]string
	var directories []string
	switch softwareName {
	case "Vectorworks":
		properties = []string{
			"net.nemetschek.vectorworks.license." + licenseYear + ".plist",
			"net.nemetschek.vectorworks." + licenseYear + ".plist",
			"net.nemetschek.vectorworks.spotlightimporter.plist",
			"net.nemetschek.vectorworks.plist",
			"net.nemetschek.vectorworksinstaller.helper.plist",
			"net.nemetschek.vectorworksinstaller.plist",
			"net.vectorworks.vectorworks." + licenseYear + ".plist",
		}
		directories = []string{
			homeDir + "/Library/Application\\ Support/Vectorworks\\ RMCache/rm" + licenseYear,
			homeDir + "/Library/Application\\ Support/Vectorworks\\ Cloud\\ Services",
			homeDir + "/Library/Application\\ Support/Vectorworks/" + licenseYear,
			homeDir + "/Library/Application\\ Support/vectorworks-installer-wrapper",
		}
	case "Vision":
		properties = []string{
			"com.qtproject.plist",
			"com.vwvision.Vision" + licenseYear + ".plist",
			"com.yourcompany.Vision.plist",
			"net.vectorworks.Vision.plist",
			"net.vectorworks.vision.license." + licenseYear + ".plist",
		}
		directories = []string{
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
	case "VCS":
	}
	return toBeCleaned{
		Properties: properties,
		Directories: directories,
	}
}
