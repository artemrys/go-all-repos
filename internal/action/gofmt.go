package action

import (
	"context"
	"log"

	"github.com/artemrys/go-all-repos/internal/repo"
	"github.com/google/go-github/github"
)

// GoFmtAction declares "go fmt" action.
type GoFmtAction struct {
	Repo         *repo.Repo
	GithubClient *github.Client
	dryRun       bool
}

// NewGoFmtAction returns a new instance of GoFmtAction.
func NewGoFmtAction(repo *repo.Repo, githubClient *github.Client, dryRun bool) *GoFmtAction {
	return &GoFmtAction{
		Repo:         repo,
		GithubClient: githubClient,
		dryRun:       dryRun,
	}
}

// Do does "go fmt" action for a particular repo.
func (a GoFmtAction) Do() {
	runPath := a.Repo.ClonedPath
	RunGit(
		[]string{
			"checkout",
			"-b",
			"go-all-repos-update",
		},
		runPath)
	Run(
		"/usr/local/go/bin/go",
		[]string{
			"/usr/local/go/bin/go",
			"fmt",
		},
		runPath,
	)
	RunGit(
		[]string{
			"add",
			".",
		},
		runPath,
	)
	RunGit(
		[]string{
			"commit",
			"-m",
			"[go-all-repos] update",
		},
		runPath,
	)
	if !a.dryRun {
		RunGit(
			[]string{
				"push",
				"origin",
				"-u",
				"go-all-repos-update",
			},
			runPath,
		)
		pullRequest := &github.NewPullRequest{
			Title: github.String("[go-all-repos] update"),
			Head:  github.String("go-all-repos-update"),
			Base:  github.String("master"),
		}
		resp, _, err := a.GithubClient.PullRequests.Create(context.Background(), "artemrys", a.Repo.Name, pullRequest)
		if err != nil {
			log.Printf("Error while creating pull request for %v: %v\n", a.Repo, err)
		} else {
			log.Printf("Created pull request for %v: %v\n", a.Repo, resp)
		}
	}
}
