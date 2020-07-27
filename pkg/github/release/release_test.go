package release

import (
  "io/ioutil"
	"testing"
)

func TestRunningCurrentVersion(t *testing.T) {
  r := New("owner", "reponame", "9.10.11")

	json, _ := ioutil.ReadFile("fixtures/9.10.11-version.json")
  r.innerStruct.releaseJson = json

	_, latest := r.IsLatestVersion()
  if !latest {
    t.Errorf("Expected version to be latest")
  }
}

func TestTarballFilename(t *testing.T) {
  r := New("owner", "reponame", "9.10.11")
	json, _ := ioutil.ReadFile("fixtures/9.10.11-version.json")
  r.innerStruct.releaseJson = json
	r.innerStruct.getLatestReleaseInfo()

  tarball := r.innerStruct.tarballFilename()
  expected := "reponame_9.10.11_darwin_amd64.tar.gz"
  if tarball != expected {
    t.Errorf("Expected: %s, got: %s", expected, tarball)
  }
}

