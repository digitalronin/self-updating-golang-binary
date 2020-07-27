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
	innerStruct myRelease
}

// These attributes need to be exported so that the json.Unmarshal call works
// correctly. But we don't want them to be exported to callers of this package,
// so we wrap them in a private, innner struct which is not exported.
type myRelease struct {
	Owner          string
	RepoName       string
	CurrentVersion string
	LatestTag      string `json:"tag_name"`
}

func New(owner string, repoName string, currentVersion string) Release {
	return Release{
		myRelease{
			Owner:          owner,
			RepoName:       repoName,
			CurrentVersion: currentVersion,
		},
	}
}

func (rd *Release) IsLatestVersion() (error, bool) {
	err := rd.innerStruct.getLatestReleaseInfo() // TODO: memoize this
	if err != nil {
		return err, false
	}

	return nil, rd.innerStruct.LatestTag == rd.innerStruct.CurrentVersion
}

func (rd *Release) SelfUpgrade() error {
	fmt.Printf("Update required. Current version: %s, Latest version: %s\n\n", rd.innerStruct.CurrentVersion, rd.innerStruct.LatestTag)

	// download tarball of latest release
	tempFilePath := "/tmp/" + rd.innerStruct.tarballFilename()

	fmt.Printf("Downloading latest tarball...\n  %s\n", rd.innerStruct.latestTarballUrl())
	rd.innerStruct.downloadFile(tempFilePath, rd.innerStruct.latestTarballUrl())

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

func (rd *myRelease) getLatestReleaseInfo() error {
	err, body := rd.getLatestReleaseJson()
	if err != nil {
		return err
	}

	json.Unmarshal(body, rd)

	return nil
}

func (rd *myRelease) getLatestReleaseJson() (error, []byte) {
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
func (rd *myRelease) downloadFile(filepath string, url string) error {
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

func (rd *myRelease) tarballFilename() string {
	return rd.RepoName + "_" + rd.LatestTag + "_" + runtime.GOOS + "_" + runtime.GOARCH + ".tar.gz"
}

func (rd *myRelease) latestTarballUrl() string {
	return fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s", rd.Owner, rd.RepoName, rd.LatestTag, rd.tarballFilename())
}

func (rd *myRelease) latestReleaseUrl() string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", rd.Owner, rd.RepoName)
}
