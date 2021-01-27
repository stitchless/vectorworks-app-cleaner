package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func FindAndChooseLicense(softwareName string) string {
	var licenses []string

	re := regexp.MustCompile("[0-9]+")

	folders, err := ioutil.ReadDir(os.Getenv("APPDATA") + "/" + softwareName)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range folders {
		appYear := re.FindString(f.Name())
		licenses = append(licenses, appYear)
	}
	return chooseLicense(softwareName, licenses)
}

func getLicense(config softwareConfig) {
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

	s, _, err := key.GetStringValue("User Serial Number")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("License is: %q\n", s)
	err = key.Close()
	if err != nil {
		log.Fatal(err)
	}
}
