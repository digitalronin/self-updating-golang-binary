package main

import (
	"fmt"
	release "github.com/digitalronin/self-updating-golang-binary/pkg/github/release"
	"os"
)

func main() {
	fmt.Printf("Self-updating golang binary demo.\n\n")

	// CurrentVersion should be the tag of the release this version of the code
	// will belong to
	r := release.New("digitalronin", "self-updating-golang-binary", "0.03")

	_, latest := r.IsLatestVersion()

	if latest {
		fmt.Println("This is the latest version")
	} else {
		err := r.SelfUpgrade()
		if err != nil {
			fmt.Printf("Unexpected error: %s\n", err)
			os.Exit(1)
		}
	}
}
