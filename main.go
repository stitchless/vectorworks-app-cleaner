package main

import (
	"fmt"
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
	Preloader   bool
	Description string
	Software    []Software
	FormData FormData
}

type FormData struct {
	Name string
	Version Version
}

type Software struct {
	Name     string
	Versions []Version
}

type Version struct {
	Year   string
	Serial string
}

// software Information
type toBeCleaned struct {
	Properties  []string
	Directories []string
}

// Set up and run the webview.
func main() {
	// Define users home directory
	var err error
	dir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// funcMap needed in order to define custom functions within go templates
	funcMap := template.FuncMap{
		// Increments int by 1 (Used to illustrate table views)
		"inc": func(i int) int {
			return i + 1
		},
		"comp": func(val1 string, val2 string) bool {
			if val1 == val2 {
				return true
			}
			return false
		},
	}

	tmpl = template.Must(template.ParseGlob(dir + "/templates/*.gohtml")).Funcs(funcMap)
	template.Must(tmpl.ParseGlob(dir + "/views/*.gohtml")).Funcs(funcMap)

	go webApp()

	// Set up Webview
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Vectorworks, Inc. - Application Utility Tool")
	w.SetSize(1000, 700, webview.HintFixed)
	w.Navigate("http://127.0.0.1:12346")
	w.Run()
}

// Error Checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// webApp Creates all Mux objects, Handlers, configures the web server and
// deploys the http server on http://127.0.0.1:12346
func webApp() {
	mux := http.NewServeMux()
	// Routes
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static"))))
	mux.HandleFunc("/", homePageHandler)             // Also catch all
	mux.HandleFunc("/editSerial", editSerialHandler) // Also catch all
	mux.HandleFunc("/updateSerial", updateSerialHandler) // Also catch all

	// Configure the webserver
	webServer := &http.Server{
		Addr:    "127.0.0.1:12346",
		Handler: mux,
	}
	//Start the web server
	err := webServer.ListenAndServe()
	check(err)
}

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
		Version:Version{
			Year: appYear,
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
		Version:Version{
			Year: appYear,
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
