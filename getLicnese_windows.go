package main

import (
	"github.com/gen2brain/dlgs"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func FindAndChooseLicense(softwareName string) string {
	var licenses []string

	re := regexp.MustCompile("[0-9]+")

	folders, err := ioutil.ReadDir(os.Getenv("APPDATA") + "/Nemetschek/" + softwareName)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range folders {
		appYear := re.FindString(f.Name())
		licenses = append(licenses, appYear)
	}
	return chooseLicense(softwareName, licenses)
}

func getSerial(config softwareConfig) string {
	key, err := registry.OpenKey(registry.CURRENT_USER, config.license, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := key.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	serial, _, err := key.GetStringValue("User Serial Number")
	if err != nil {
		log.Fatal(err)
	}
	// TESTING
	// TODO: Remove Printf
	//fmt.Printf("License is: %q\n", serial)
	//err = key.Close()
	//if err != nil {
	//	log.Fatalf("Error Closing Registry: %e", err)
	//}
	return serial
}

func inputNewSerial(serial string) string {
	newSerial, _, err := dlgs.Entry("Input New Serial", "Please input a new Series E Serial:", serial)
	if err != nil {
		log.Fatal(err)
	}
	return newSerial
}

func replaceOldSerial(newSerial string, config softwareConfig) {
	// TODO: Clean the input before replacing.
	key, err := registry.OpenKey(registry.CURRENT_USER, config.license, registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := key.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = key.SetStringValue("User Serial Number", newSerial)
	if err != nil {
		log.Fatal(err)
	}
}