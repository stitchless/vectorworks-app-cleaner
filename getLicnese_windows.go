package main

import (
	"fmt"
	"log"
	"os/user"
	"golang.org/x/sys/windows/registry"
	"strconv"
)

func getWindowsLicenses(softwareSelect string) []string {
	var licenses []string

	re := regexp.MustCompile("[0-9]+")

	folders, err := ioutil.ReadDir(os.Getenv("APPDATA") + "/" + softwareSelect)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range folders {
		appYear := re.FindString(f.Name())
		licenses = append(licenses, appYear)
	}
	return chooseLicense(softwareName, licenses)
}

func getLicense(config workingData) string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	userUIDStr := currentUser.Uid[4:4]
	uid, err := strconv.Atoi(userUIDStr)
	if err != nil {
		log.Fatal(err)
	}

	key, err := registry.OpenKey(registry.CURRENT_USER, config.license, uint32(uid))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(key)
	return key
}

// convert Software License year to version number.
func doTheMath(appYear string) string {
	values := strings.Split(appYear, "")[2:4]
	appVersion := values[0] + values[1]
	i, err := strconv.Atoi(appVersion)
	if err != nil {
		log.Fatal(err)
	}
	versionMath := i + 6
	appVersion = strconv.Itoa(versionMath)
	return appVersion
}
