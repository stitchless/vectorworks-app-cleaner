package main

import (
	"github.com/gen2brain/dlgs"
	"github.com/webview/webview"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

// TODO: Refactor to use WebUI as opposed to a series of dialogs.

// Data
type softwareConfig struct {
	plist       []string
	registry    []string
	directories []string
	license     string
	vcs         []string
	vision      []string
}

// Home Directory OS dependent
var homeDir, _ = os.UserHomeDir()

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var dir string

func init() {
	//events := make(chan string, 1000)

	var err error
	dir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
}

type PageVariables struct {
	Date         string
	Time         string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	now := time.Now() // find the time right now
	HomePageVars := PageVariables{ //store the date and time in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("index.html") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func main() {
	go app()
	// Set up Webview
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Vectorworks, Inc.")
	w.SetSize(1200, 850, webview.HintNone)
	w.Navigate("http://127.0.0.1:12346/public/html/index.html")
	w.Run()
}

func app() {
	mux := http.NewServeMux()
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(dir+"/public"))))
	webServer := &http.Server{
		Addr:    "127.0.0.1:12346",
		Handler: mux,
	}
	err := webServer.ListenAndServe()
	check(err)
}

//// Start by picking between "Vectorworks" and "Vision"
//softwareName, cancelDiag, err := dlgs.List("Vectorworks, Inc. - App Cleaner", "What software package are you attempting to edit?", []string{"Vectorworks", "Vision"})
//if err != nil {
//	log.Fatalf("cannot use the dialog as expected: %e", err)
//}
//
//if !cancelDiag {
//	log.Print("Closed by user...")
//}
//
//license := FindAndChooseLicense(softwareName)   // Find and Choose Versions of software based on license
//config := generateConfig(softwareName, license) // generate proper data for select license version
//chooseAction(config)


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
		serial := getSerial(config)
		newSerial := inputNewSerial(serial)
		replaceOldSerial(newSerial, config)
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
