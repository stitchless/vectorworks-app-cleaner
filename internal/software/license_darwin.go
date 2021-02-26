package software

type LicenseOpts struct {
	serial map[string]string `plist:"NNA User License"`
}

// getSerial will read in a plist, decode it and return a keyed value as a string value
func getSerial(installation Installation) string {
	var licenseLocation string
	switch installation.Software {
	case SoftwareVectorworks:
		licenseLocation = homeDir + "/Library/Preferences/net.nemetschek.vectorworks.license." + installation.Year + ".plist"
	case SoftwareVision:
		licenseLocation = homeDir + "/Library/Preferences/net.vectorworks.vision.license." + installation.Year + ".plist"
	}

	// Read in plist
	plistFile, err := ioutil.ReadFile(licenseLocation)
	buffer := bytes.NewReader(plistFile)
	check(err)

	// parse and return plist serial
	var plistData LicenseOpts
	decoder := plist.NewDecoder(buffer)
	err = decoder.Decode(&plistData.serial)
	check(err)

	return plistData.serial[`NNA User License`]
}

// replaceOldSerial
func replaceOldSerial(softwareName string, appYear string, newSerial string) {
	licenseLocation := getSerialLocation(softwareName, appYear)
	plistFile, err := os.Open(licenseLocation)
	check(err)
	err = plistFile.Truncate(0)

	newSerial = cleanSerial(newSerial) // Clean and verify serial

	plistData := &LicenseOpts{
		serial: map[string]string{"NNA User License": newSerial},
	}

	fmt.Println(plistData.serial)
	buffer := &bytes.Buffer{}
	encoder := plist.NewEncoder(buffer)

	err = encoder.Encode(plistData.serial)
	check(err)

	err = os.WriteFile(licenseLocation, buffer.Bytes(), 0644)
	check(err)

	w := bufio.NewWriter(buffer)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	err = w.Flush()
	check(err)

	refreshPList()
}

func refreshPList() {
	fmt.Println("Refreshing plist files...")
	// osascript -e 'do shell script "sudo killall -u $USER cfprefsd" with administrator privileges'
	cmd := exec.Command(`osascript`, "-s", "h", "-e", `do shell script "sudo killall -u $USER cfprefsd" with administrator privileges`)
	stderr, err := cmd.StderrPipe()
	log.SetOutput(os.Stderr)
	check(err)

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}

	cmdErrOutput, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", cmdErrOutput)

	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

// cleanSerial will take in a string, remove any empty strings
// and confirm a regex pattern.  If regex is valid the string is returned.
func cleanSerial(serial string) string {
	r := regexp.MustCompile(`(.{6})-(.{6})-(.{6})-(.{6})`)
	parseSerial := r.FindAllString(serial, -1)
	if len(parseSerial) != 0 {
		return parseSerial[0]
		// TODO: REFER TO THIS WHEN PARSING OUT LICENSE MEANING!
	}
	panic("ERROR: REPLACE THIS WITH A TOAST SHOWING INVALID INPUT!")
}
