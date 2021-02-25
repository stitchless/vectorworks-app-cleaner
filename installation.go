package main

type Installation struct {
	License License
	Software Software
	Properties []string
	Directories []string
	Year string
}

func FindInstallationsBySoftware(software Software) ([]Installation, error) {
	var installations []Installation

	years := findInstallationYears(software)

	// Attach configs, versions, and app years all into on object then return that object
	for _, year := range years {
		installation := Installation{
			Software:    software,
			Year:        year,
		}

		installation.Properties = findProperties(installation)
		installation.Directories = findDirectories(installation)
		installation.License = License{
			Serial: getSerial(installation),
		}

		installations = append(installations, installation)
	}

	return installations, nil
}
