package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

type releaseData struct {
	LatestTag string `json:"tag_name"`
}

// TODO: get this working next
// func (rd *releaseData) downloadLatestTarballUrl() (error, string) {
// 	downloadTo := "/tmp/" + rd.tarballFilename()
// 	err := DownloadFile(downloadTo, rd.latestTarballUrl())
// 	return err, downloadTo
// }

func (rd *releaseData) tarballFilename() string {
	return repoName + "_" + rd.LatestTag + "_" + runtime.GOOS + "_" + runtime.GOARCH + ".tar.gz"
}

func (rd *releaseData) latestTarballUrl() string {
	return "https://github.com/" + owner + "/" + rd.tarballFilename()
}

func main() {
	fmt.Println("Self-updating golang binary")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	_, latest := isThisTheLatestVersion()

	if latest {
		fmt.Println("This is the latest version")
	} else {
		fmt.Println("Need to update")
	}

	filename, _ := os.Executable()
	fmt.Println(filename)
}

// --------------------------------------------------------

func isThisTheLatestVersion() (error, bool) {
	var relData releaseData
	err := getLatestReleaseInfo(&relData)
	if err != nil {
		return err, false
	}
	fmt.Println(relData.latestTarballUrl())
	return nil, relData.LatestTag == version
}

func getLatestReleaseInfo(rd *releaseData) error {
	err, body := getLatestReleaseJson()
	if err != nil {
		return err
	}

	json.Unmarshal(body, rd)

	return nil
}

func getLatestReleaseJson() (error, []byte) {
	response, err := http.Get(latestReleaseUrl)
	if err != nil {
		return err, nil
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err, nil
	}
	return nil, body
}

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
