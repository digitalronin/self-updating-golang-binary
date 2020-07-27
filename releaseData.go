package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

type releaseData struct {
	CurrentVersion string
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

func (rd *releaseData) isLatestVersion() (error, bool) {
  // move this method into this class
	err := rd.getLatestReleaseInfo()
	if err != nil {
		return err, false
	}
	fmt.Println(rd.latestTarballUrl())
	return nil, rd.LatestTag == version
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

