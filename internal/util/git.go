package util

import (
	"github.com/go-git/go-git/v5"
	"strings"
)

func GetGitCloneUrl() ([]string, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, err
	}
	remotes, err := repo.Remotes()
	if err != nil {
		return nil, err
	}
	head, err := repo.Head()
	if err != nil {
		return nil, err
	}
	return []string{"git", "clone", "-b", head.Name().Short(), "--single-branch", convertToHTTPS(remotes[0].Config().URLs[0]), "/mnt/repo"}, nil
}

func convertToHTTPS(url string) string {
	if strings.HasPrefix(url, "git@github.com:") {
		return strings.Replace(url, "git@github.com:", "https://github.com/", 1)
	}
	return url // Already HTTPS
}
