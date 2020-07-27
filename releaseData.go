package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

type releaseData struct {
  Owner string
  RepoName string
	CurrentVersion string
	LatestTag string `json:"tag_name"`
}

func (rd *releaseData) latestReleaseUrl() string {
  releasesUrl := "https://api.github.com/repos/" + rd.Owner + "/" + rd.RepoName + "/releases"
  return releasesUrl + "/latest"
}

// TODO: get this working next
// func (rd *releaseData) downloadLatestTarballUrl() (error, string) {
// 	downloadTo := "/tmp/" + rd.tarballFilename()
// 	err := DownloadFile(downloadTo, rd.latestTarballUrl())
// 	return err, downloadTo
// }

func (rd *releaseData) tarballFilename() string {
	return rd.RepoName + "_" + rd.LatestTag + "_" + runtime.GOOS + "_" + runtime.GOARCH + ".tar.gz"
}

func (rd *releaseData) latestTarballUrl() string {
	return "https://github.com/" + rd.Owner + "/" + rd.tarballFilename()
}

func (rd *releaseData) isLatestVersion() (error, bool) {
  err := rd.getLatestReleaseInfo() // TODO: memoize this
	if err != nil {
		return err, false
	}
	fmt.Println(rd.latestTarballUrl())
	return nil, rd.LatestTag == rd.CurrentVersion
}

func (rd *releaseData) getLatestReleaseInfo() error {
	err, body := rd.getLatestReleaseJson()
	if err != nil {
		return err
	}

	json.Unmarshal(body, rd)

	return nil
}

func (rd *releaseData) getLatestReleaseJson() (error, []byte) {
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

