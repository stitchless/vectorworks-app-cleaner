package main

import (
	"fmt"
	"net/http"
)

// TODO: Get Vectorworks Info [Done]
// TODO: Get Vision Info [WIP]
// TODO: Determine if VCS is installed

// homePageHandler is the initial page with all software information held on it.
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

func updateSerialHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var softwareName string
	var appYear string
	var serial string

	fmt.Println(r.Form)
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
