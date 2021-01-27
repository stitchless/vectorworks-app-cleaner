package main

import "fmt"

func cleanApplication(config softwareConfig) {
	// Deletes relevant registry entries for select software/version
	for _, property := range config.registry {
		fmt.Println(property)
	}

	for _, directory := range config.directories {
		fmt.Println(directory)
	}
}
