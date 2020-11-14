// Package repo contains structs and methods to work with repository.
package repo

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/google/go-github/github"
)

// Repo holds info about repository.
type Repo struct {
	Name       string
	CloneURL   string
	ClonedPath string
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

// Clone clones repository to the current folder with "go-all-repos" prefix.
func (r *Repo) Clone(prefix string) {
	log.Printf("Cloning repo %q @ %s\n", r.Name, r.CloneURL)
	clonedPath := fmt.Sprintf("%s-%s", prefix, r.Name)
	if _, err := exec.Command(
		"git",
		"clone",
		r.CloneURL,
		clonedPath,
		"--depth",
		"1",
	).Output(); err != nil {
		clonedPath = ""
		log.Printf("Error while cloning repo %q: %v\n", r.Name, err)
	}
	r.ClonedPath = clonedPath
}

func buildCloneURL(username, repo string) string {
	return fmt.Sprintf("https://github.com/%s/%s.git", username, repo)
}
