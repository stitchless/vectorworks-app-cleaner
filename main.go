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

	// funcMap needed in order to define custom functions within go templates
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
	mux.HandleFunc("/", homePageHandler) // Also catch all
	mux.HandleFunc("/editSerial", editSerialHandler) // Also catch all

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

	homeScreen := htmlValues{
		Title:       "Welcome to the Vectorworks Utility Tool",
		Description: "This utility will allow you to make a variety of changes to Vectorworks, Vision, and Vectorworks Cloud Services Desktop App.  Simply select an action from the list below...",
		Software: []Software{
			{Name: "Vectorworks", Versions: vectorworksVersions},
			{Name: "Vision", Versions: visVersions},
		},
		Footer: "This application will work for Vectorworks, Vision, and Vectorworks Cloud Services Desktop App.",
	}

	err := tmpl.ExecuteTemplate(w, "homePage", homeScreen)
	check(err)
}
// TODO: Separate views based on license type or localizations via Tabs
// TODO: Show Actions as Modals?

func editSerialHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
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

	err := tmpl.ExecuteTemplate(w, "homePage", homeScreen)
	check(err)
}