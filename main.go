package main

import (
	"fmt"
	release "github.com/digitalronin/self-updating-golang-binary/pkg/github/release"
)

const version = "0.0.3"

func main() {
	fmt.Printf("Self-updating golang binary demo.\n\n")

	// CurrentVersion should be the tag of the release this version of the code
	// will belong to
	r := release.New("digitalronin", "self-updating-golang-binary", version)
  r.UpgradeIfNotLatest()
}
