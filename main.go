package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/artemrys/go-all-repos/internal/action"
	"github.com/artemrys/go-all-repos/internal/config"
	"github.com/artemrys/go-all-repos/internal/repo"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const clonePrefix = "go-all-repos"

func main() {
	config, output, err := config.NewFromFlags(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		fmt.Println(output)
		os.Exit(2)
	} else if err != nil {
		fmt.Println("got error:", err)
		fmt.Println("output:\n", output)
		os.Exit(1)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GithubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	githubClient := github.NewClient(tc)
	repos := []*repo.Repo{}
	if len(config.RepoNames) == 0 {
		ctx := context.Background()
		opt := &github.RepositoryListOptions{Type: "public"}
		githubRepos, _, err := githubClient.Repositories.List(ctx, config.Username, opt)
		if err != nil {
			log.Fatalf("Error while getting repositories for %q: %v\n", config.Username, err)
		}
		for _, r := range githubRepos {
			repos = append(repos, repo.NewFromGithub(r))
		}
	} else {
		for _, r := range config.RepoNames {
			repos = append(repos, repo.New(config.Username, r))
		}
	}
	for _, r := range repos {
		r.Clone(clonePrefix)
		e := action.NewGoFmtAction(r, githubClient, config.DryRun)
		e.Do()
	}
}
