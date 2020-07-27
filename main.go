package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Self-updating golang binary")

	// CurrentVersion should be the tag of the release this version of the code
	// will belong to
	rd := releaseData{
		RepoName:       "self-updating-golang-binary",
		Owner:          "digitalronin",
		CurrentVersion: "0.0.1",
	}

	_, latest := rd.isLatestVersion()

	if latest {
		fmt.Printf("This is the latest version: %s\n", rd.LatestTag)
	} else {
		fmt.Printf("Current version: %s, Latest version: %s\n", rd.CurrentVersion, rd.LatestTag)
		fmt.Println("Update required")

		err := rd.SelfUpdate()
		if err != nil {
			fmt.Printf("Unexpected error: %s\n", err)
			os.Exit(1)
		}
	}

	filename, _ := os.Executable()
	fmt.Println(filename)
}
