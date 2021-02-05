package main

import (
	"net/http"
)

// TODO: Determine if VCS is installed

// homePageHandler is the initial page with all software information held on it.
// Each time an action is done the user is returned to this screen
// From this screen you can edit license info or clean up application data.
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	vectorworksVersions := fetchAppInfo("Vectorworks")
	visVersions := fetchAppInfo("Vision")

	templateValues := htmlValues{
		Preloader:   true,
		Title:       "Welcome to the Vectorworks Utility Tool",
		Description: "This utility will allow you to make a variety of changes to Vectorworks, Vision, and Vectorworks Cloud Services Desktop App.  Simply select an action from the list below...",
		Software: []Software{
			{Name: "Vectorworks", Versions: vectorworksVersions},
			{Name: "Vision", Versions: visVersions},
		},
	}

	err := tmpl.ExecuteTemplate(w, "homePage", templateValues)
	check(err)
}

// TODO: Show localizations via Tabs
// TODO: Show Actions as Modals?
// TODO: Illustrate license types (Private Repo)

// editSerialHandler The screen to chose the user a text field to update a selected serial number
func editSerialHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var softwareName string
	var appYear string
	var serial string

	for key, value := range r.Form {
		switch key {
		case "softwareName":
			softwareName = value[0]
		case "appYear":
			appYear = value[0]
		case "serial":
			serial = value[0]
		}
	}

	//var formData string
	formData := FormData{
		Name: softwareName,
		Version: Version{
			Year:   appYear,
			Serial: serial,
		},
	}

	vectorworksVersions := fetchAppInfo("Vectorworks")
	visVersions := fetchAppInfo("Vision")

	// Serve the screen
	templateValues := htmlValues{
		Preloader:   false,
		Title:       "Welcome to the Vectorworks Utility Tool!",
		Description: "This utility will allow you to make a variety of changes to Vectorworks, Vision, and Vectorworks Cloud Services Desktop App.  Simply select an action from the list below...",
		Software: []Software{
			{Name: "Vectorworks", Versions: vectorworksVersions},
			{Name: "Vision", Versions: visVersions},
		},
		FormData: formData,
	}

	err := tmpl.ExecuteTemplate(w, "editSerial", templateValues)
	check(err)
}

// updateSerialHandler will send the filled in text field and update the serial
// Once updated, the home homePageHandler is called and the home screen is shown
func updateSerialHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var softwareName string
	var appYear string
	var serial string

	for key, value := range r.Form {
		switch key {
		case "softwareName":
			softwareName = value[0]
		case "appYear":
			appYear = value[0]
		case "serial":
			serial = value[0]
		}
	}

	replaceOldSerial(softwareName, appYear, serial)

	//var formData string
	formData := FormData{
		Name: softwareName,
		Version: Version{
			Year:   appYear,
			Serial: serial,
		},
	}

	vectorworksVersions := fetchAppInfo("Vectorworks")
	visVersions := fetchAppInfo("Vision")

	templateValues := htmlValues{
		Preloader:   false,
		Title:       "Welcome to the Vectorworks Utility Tool!",
		Description: "This utility will allow you to make a variety of changes to Vectorworks, Vision, and Vectorworks Cloud Services Desktop App.  Simply select an action from the list below...",
		Software: []Software{
			{Name: "Vectorworks", Versions: vectorworksVersions},
			{Name: "Vision", Versions: visVersions},
		},
		FormData: formData,
	}

	err := tmpl.ExecuteTemplate(w, "homePage", templateValues)
	check(err)
}
