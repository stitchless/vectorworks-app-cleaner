package main

import (
	"embed"
	"github.com/webview/webview"
	"html/template"
	"log"
	"net/http"
	"os"
)

// EMBEDS
//go:embed webview.dll WebView2Loader.dll
//go:embed templates/* static/* views/*
var content embed.FS

// Globals
var homeDir, _ = os.UserHomeDir()
var dir string
var tmpl *template.Template

type htmlValues struct {
	Title       string
	Preloader   bool
	Description string
	Softwares []Software
	FormData    FormData
}

type FormData struct {
	Name    string
	Year   string
	Serial string
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
		"FindInstallationsBySoftware": FindInstallationsBySoftware,
	}

	tmpl = template.Must(template.ParseGlob(dir + "/templates/*.html.tmpl")).Funcs(funcMap)
	template.Must(tmpl.ParseGlob(dir + "/views/*.html.tmpl")).Funcs(funcMap)

	go runWebserver()

	// Set up Webview
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Vectorworks, Inc. - Application Utility Tool")
	w.SetSize(1000, 700, webview.HintFixed)
	w.Navigate("http://127.0.0.1:12346")
	w.Run()
}

// Creates all Mux objects, Handlers, configures the web server and
// deploys the http server on http://127.0.0.1:12346
func runWebserver() {
	mux := http.NewServeMux()
	// Routes
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static"))))
	mux.HandleFunc("/", homePageHandler) // Handle home page, also catch all
	mux.HandleFunc("/editSerial", editSerialHandler)
	mux.HandleFunc("/updateSerial", updateSerialHandler)

	// Configure the webserver
	webServer := &http.Server{
		Addr:    "127.0.0.1:12346",
		Handler: mux,
	}
	//Start the web server
	err := webServer.ListenAndServe()
	check(err)
}
