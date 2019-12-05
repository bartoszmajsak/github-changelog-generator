package github

import (
	"net/http"

	"github.com/bartoszmajsak/github-changelog-generator/version"

	"github.com/google/go-github/github"
	"golang.org/x/net/context"
)

func LatestRelease() (string, error) {
	httpClient := http.Client{}
	defer httpClient.CloseIdleConnections()

	client := github.NewClient(&httpClient)
	latestRelease, _, err := client.Repositories.
		GetLatestRelease(context.Background(), "bartoszmajsak", "github-changelog-generator")
	if err != nil {
		return "", err
	}
	return *latestRelease.Name, nil
}

func IsLatestRelease(latestRelease string) bool {
	return latestRelease == version.Version
}
