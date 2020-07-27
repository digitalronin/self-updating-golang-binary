package main

import (
	"fmt"
	"os"
  release "github.com/digitalronin/self-updating-golang-binary/pkg/github/release"
)

func main() {
	fmt.Printf("Self-updating golang binary demo.\n\n")

	// CurrentVersion should be the tag of the release this version of the code
	// will belong to
	rd := release.Release{
		RepoName:       "self-updating-golang-binary",
		Owner:          "digitalronin",
		CurrentVersion: "0.0.3",
	}

	_, latest := rd.IsLatestVersion()

	if latest {
		fmt.Println("This is the latest version")
	} else {
		err := rd.SelfUpdate()
		if err != nil {
			fmt.Printf("Unexpected error: %s\n", err)
			os.Exit(1)
		}
	}
}
