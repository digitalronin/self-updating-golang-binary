package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// version should be the tag of the release this version of the code will
// belong to
const version = "0.0.2"

const releasesUrl = "https://api.github.com/repos/digitalronin/self-updating-golang-binary/releases/latest"

type releaseData struct {
	LatestTag string `json:"tag_name"`
}

func main() {
	fmt.Println("Self-updating golang binary")

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
	response, err := http.Get(releasesUrl)
	if err != nil {
		return err, nil
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err, nil
	}
	return nil, body
}

// // parse arbitrary JSON into a map
// func getLatestRelease() error {
// 	response, err := http.Get(releasesUrl)
// 	if err != nil {
// 		return err
// 	}
// 	body, _ := ioutil.ReadAll(response.Body)
//
// 	var result map[string]interface{}
//
// 	json.Unmarshal(body, &result)
//
// 	fmt.Println("TAG: ", result["tag_name"])
//
// 	for key, value := range result {
// 		// Each value is an interface{} type, that is type asserted as a string
//
// 		switch value.(interface{}).(type) {
// 		case string:
// 			fmt.Println(key, value.(string))
// 		case bool:
// 			fmt.Println(key, value.(bool))
// 		}
// 	}
// 	// fmt.Println("tag: %s", result.tag_name.(string))
//
// 	// return nil, string(body)
// 	return nil
// }
