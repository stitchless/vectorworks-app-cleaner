package main

import (
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

	tmpl = template.Must(template.ParseGlob(dir+"/templates/*.gohtml"))
	template.Must(tmpl.ParseGlob(dir+"/views/*.gohtml"))

	go webApp()

	// Set up Webview
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Vectorworks, Inc.")
	w.SetSize(1200, 850, webview.HintNone)
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
	mux.HandleFunc("/", basePageHandler)
	//mux.HandleFunc("/chooseYear", chooseYearHandler)
	//mux.HandleFunc("/chooseAction", chooseActionHandler)
	//mux.HandleFunc("/replaceLicense", replaceLicenseHandler)
	//mux.HandleFunc("/cleanApp", cleanAppHandler)
	//mux.HandleFunc("/cancel", cancelHandler)
	//mux.HandleFunc("/", catchAllHandler)

	// Configure the webserver
	webServer := &http.Server{
		Addr:    "127.0.0.1:12346",
		Handler: mux,
	}
	//Start the web server
	err := webServer.ListenAndServe()
	check(err)
}

type MyList struct {
	Option	string
	Done	bool
}

type testData struct {
	Title	string
	MyList	*[]MyList

}

func basePageHandler(w http.ResponseWriter, r *http.Request) {
	// Get Data
	derpData := testData{
		Title: "Herp Derp",
		MyList: &[]MyList{
			{Option: "Option 1", Done: true},
			{Option: "Option 2", Done: false},
			{Option: "Option 3", Done: true},
		},
	}
	err := tmpl.ExecuteTemplate(w, "homePage", derpData)
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
