package release

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

type Release struct {
	Owner          string
	RepoName       string
	CurrentVersion string
	LatestTag      string `json:"tag_name"`
}

func New(owner string, repoName string, currentVersion string) Release {
  return Release{
    Owner: owner,
    RepoName: repoName,
    CurrentVersion: currentVersion,
  }
}

func (rd *Release) IsLatestVersion() (error, bool) {
	err := rd.getLatestReleaseInfo() // TODO: memoize this
	if err != nil {
		return err, false
	}

	return nil, rd.LatestTag == rd.CurrentVersion
}

func (rd *Release) SelfUpgrade() error {
	fmt.Printf("Update required. Current version: %s, Latest version: %s\n\n", rd.CurrentVersion, rd.LatestTag)

	// download tarball of latest release
	tempFilePath := "/tmp/" + rd.tarballFilename()

	fmt.Printf("Downloading latest tarball...\n  %s\n", rd.latestTarballUrl())
	rd.downloadFile(tempFilePath, rd.latestTarballUrl())

	fmt.Println("Unpacking...")
	cmd := exec.Command("tar", "xzf", tempFilePath, "--cd", "/tmp/")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	// move unpacked binary into place
	filename, _ := os.Executable()
	fmt.Printf("Replacing %s\n\n", filename)
	cmd = exec.Command("mv", "/tmp/myapp", filename)
	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Printf("Upgrade successful. Please repeat your previous command.\n")

	return nil
}

// -------------------------------------------------------------

func (rd *Release) latestReleaseUrl() string {
	releasesUrl := "https://api.github.com/repos/" + rd.Owner + "/" + rd.RepoName + "/releases"
	return releasesUrl + "/latest"
}

// TODO: get this working next
// func (rd *release) downloadLatestTarballUrl() (error, string) {
// 	downloadTo := "/tmp/" + rd.tarballFilename()
// 	err := DownloadFile(downloadTo, rd.latestTarballUrl())
// 	return err, downloadTo
// }

func (rd *Release) tarballFilename() string {
	return rd.RepoName + "_" + rd.LatestTag + "_" + runtime.GOOS + "_" + runtime.GOARCH + ".tar.gz"
}

func (rd *Release) latestTarballUrl() string {
	return "https://github.com/" + rd.Owner + "/" + rd.RepoName + "/releases/download/" + rd.LatestTag + "/" + rd.tarballFilename()
}

func (rd *Release) getLatestReleaseInfo() error {
	err, body := rd.getLatestReleaseJson()
	if err != nil {
		return err
	}

	json.Unmarshal(body, rd)

	return nil
}

func (rd *Release) getLatestReleaseJson() (error, []byte) {
	response, err := http.Get(rd.latestReleaseUrl())
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
func (rd *Release) downloadFile(filepath string, url string) error {
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
