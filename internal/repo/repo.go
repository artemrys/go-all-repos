// Package repo contains structs and methods to work with repository.
package repo

import (
	"fmt"

	"github.com/google/go-github/github"
)

// Repo holds info about repository.
type Repo struct {
	Name     string
	CloneURL string
}

// New creates new Repo instance.
func New(username, repo string) *Repo {
	return &Repo{
		Name:     repo,
		CloneURL: buildCloneURL(username, repo),
	}
}

// NewFromGithub creates new Repo from Github repository.
func NewFromGithub(githubRepo *github.Repository) *Repo {
	return &Repo{
		Name:     *githubRepo.Name,
		CloneURL: *githubRepo.CloneURL,
	}
}

func buildCloneURL(username, repo string) string {
	return fmt.Sprintf("https://github.com/%s/%s.git", username, repo)
}
