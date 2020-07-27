package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

// version should be the tag of the release this version of the code will
// belong to
const version = "0.0.1"
const owner = "digitalronin"
const repoName = "self-updating-golang-binary"

const releasesUrl = "https://api.github.com/repos/" + owner + "/" + repoName + "/releases"
const latestReleaseUrl = releasesUrl + "/latest"


func main() {
	fmt.Println("Self-updating golang binary")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

  rd := releaseData{CurrentVersion: version}

	_, latest := rd.isLatestVersion()

	if latest {
		fmt.Println("This is the latest version")
	} else {
		fmt.Println("Need to update")
	}

	filename, _ := os.Executable()
	fmt.Println(filename)
}

// --------------------------------------------------------


// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
