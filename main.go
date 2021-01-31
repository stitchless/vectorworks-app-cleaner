package main

import (
	"fmt"
	"github.com/gen2brain/dlgs"
	"github.com/webview/webview"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Globals
var homeDir, _ = os.UserHomeDir()
var dir string
var tmpl *template.Template

type htmlValues struct {
	Title       string
	Description string
	Software    []Software
	Footer      string
}

type Software struct {
	Name     string
	Versions []Version
}

type Version struct {
	Year   string
	Serial string
	Config softwareConfig
}

// software Information
type softwareConfig struct {
	plist       []string
	registry    []string
	directories []string
	license     string
	vcs         []string
	vision      []string
}

// Set up and run the webview.
func main() {
	// Define users home directory
	var err error
	dir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Needed
	funcMap := template.FuncMap{
		// Increments int by 1 (Used to illustrate table views)
		"inc": func(i int) int {
			return i + 1
		},
	}

	tmpl = template.Must(template.ParseGlob(dir + "/templates/*.gohtml")).Funcs(funcMap)
	template.Must(tmpl.ParseGlob(dir + "/views/*.gohtml")).Funcs(funcMap)

	go webApp()

	// Set up Webview
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Vectorworks, Inc. - Application Utility Tool")
	w.SetSize(800, 600, webview.HintFixed)
	w.Navigate("http://127.0.0.1:12346")
	w.Run()
}

// Error Checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func webApp() {
	mux := http.NewServeMux()
	// Routes
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static"))))
	mux.HandleFunc("/", homePageHandler) // Also catch all
	//mux.HandleFunc("/chooseYear", chooseYearHandler)
	//mux.HandleFunc("/chooseAction", chooseActionHandler)
	//mux.HandleFunc("/replaceLicense", replaceLicenseHandler)
	//mux.HandleFunc("/cleanApp", cleanAppHandler)
	//mux.HandleFunc("/cancel", cancelHandler)

	// Configure the webserver
	webServer := &http.Server{
		Addr:    "127.0.0.1:12346",
		Handler: mux,
	}
	//Start the web server
	err := webServer.ListenAndServe()
	check(err)
}

// TODO: Get VW Versions, Get Vision Version, Determine if VCS is installed

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	vectorworksVersions := fetchAppInfo("Vectorworks")
	visVersions := fetchAppInfo("Vision")
	fmt.Println(visVersions)

	homeScreen := htmlValues{
		Title:       "Welcome to the Vectorworks Utility Tool",
		Description: "This utility will allow you to make a variety of changes to Vectorworks, Vision, and Vectorworks Cloud Services Desktop App.  Simply select an action from the list below...",
		Software: []Software{
			{Name: "Vectorworks", Versions: vectorworksVersions},
			{Name: "Vision", Versions: visVersions},
		},
		Footer: "This application will work for Vectorworks, Vision, and Vectorworks Cloud Services Desktop App.",
	}

	err := tmpl.ExecuteTemplate(w, "homePageAlt", homeScreen)
	check(err)
}

// TODO: Separate views based on license type or localizations via Tabs
// TODO: Shows all results on screen with actions next to each version
// TODO: Show Actions as Modals?

// Allow user to choose which licence to start working with.
func chooseLicense(softwareName string, licenses []string) string {
	pickedLicense, _, err := dlgs.List("Choose your license", "Please pick from the list of found "+softwareName+" licenses.", licenses)
	if err != nil {
		log.Fatal(err)
	}
	return pickedLicense // return string with 4 digits representing the application license year.
}

func chooseAction(config softwareConfig) {
	items := []string{"Clean Application", "Clean VCS", "Replace License"}
	choice, _, err := dlgs.List("Chose your action", "What would you like to do?", items)
	if err != nil {
		log.Fatal(err)
	}

	switch choice {
	// Replace old license with new one
	case "Replace License":
		//serial := getSerial(config)
		//newSerial := inputNewSerial(serial)
		//replaceOldSerial(newSerial, config)
	// Remove all VCS directories
	// TODO: Move this to first step.
	case "Clean VCS":
		cleanVCS(config)
	// Removes all properties and files/folders for the given software/version
	case "Clean Application":
		cleanApplication(config)
	}

	_, _ = dlgs.Info("Finished!", choice+" is finished running")
}
