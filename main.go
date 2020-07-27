package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Self-updating golang binary demo.\n\n")

	// CurrentVersion should be the tag of the release this version of the code
	// will belong to
	rd := releaseData{
		RepoName:       "self-updating-golang-binary",
		Owner:          "digitalronin",
		CurrentVersion: "0.0.4",
	}

	_, latest := rd.isLatestVersion()

	if latest {
		fmt.Printf("This is the latest version: %s\n", rd.LatestTag)
	} else {
		err := rd.SelfUpdate()
		if err != nil {
			fmt.Printf("Unexpected error: %s\n", err)
			os.Exit(1)
		}
	}
}
