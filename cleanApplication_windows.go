package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
)

func cleanApplication(config softwareConfig) {
	// Deletes relevant registry entries for select software/version
	for _, property := range config.registry {
		k, err := registry.OpenKey(registry.CURRENT_USER, property, registry.ALL_ACCESS)
		if err != nil {
			log.Fatal(err)
		}
		defer k.Close()

		names, _ := k.ReadSubKeyNames(50)

		for _, name := range names {
			_ = registry.DeleteKey(k, name)
		}
		_ = registry.DeleteKey(k, "")
	}
	// TODO: Check for directory after as a way to verify deletion.

	for _, directory := range config.directories {
		fmt.Println(directory)
	}
}

