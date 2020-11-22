// Package repo contains structs and methods to work with repository.
package repo

import (
	"fmt"
	"log"

	"github.com/artemrys/go-all-repos/internal/helpers"
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

// Clone clones repository to the current folder with `prefix` prefix.
func (r *Repo) Clone(prefix string) {
	log.Printf("Cloning repo %q from %s\n", r.Name, r.CloneURL)
	clonedPath := fmt.Sprintf("%s-%s", prefix, r.Name)
	helpers.RunGit(
		[]string{
			"clone",
			r.CloneURL,
			clonedPath,
			"--depth",
			"1",
		},
		"",
	)
	r.ClonedPath = clonedPath
}

func buildCloneURL(username, repo string) string {
	return fmt.Sprintf("https://github.com/%s/%s.git", username, repo)
}
