package main

import (
	"os"
)

func cleanApplication(config softwareConfig) {
	plistPath := homeDir + "/Library/Preferences/"
	// Deletes relevant plist files for select software/version
	for _, plist := range config.plist {
		_ = os.RemoveAll(plistPath + plist)
		//TODO: Add logging for user feedback.
	}
	// TODO: Check for directory after as a way to verify deletion.

	for _, directory := range config.directories {
		_ = os.RemoveAll(directory)
		// TODO: Implement error checking.
	}
}

func cleanVCS(config softwareConfig) {
	for _, directory := range config.vcs {
		_ = os.RemoveAll(directory)
	}
}
