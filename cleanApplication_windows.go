package main

import (
	"golang.org/x/sys/windows/registry"
	"os"
)

func cleanApplication(config softwareConfig) {
	// Deletes relevant registry entries for select software/version
	for _, property := range config.registry {
		k, _ := registry.OpenKey(registry.CURRENT_USER, property, registry.ALL_ACCESS)

		defer k.Close()

		names, _ := k.ReadSubKeyNames(-1)

		for _, name := range names {
			_ = registry.DeleteKey(k, name)
		}
		_ = registry.DeleteKey(k, "")
	}
	// TODO: Check for directory after as a way to verify deletion.

	for _, directory := range config.directories {
		_ = os.RemoveAll(directory)
		// TODO: Implement error checking.
	}
}

