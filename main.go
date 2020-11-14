package main

import (
	"context"
	"flag"
	"log"
	"strings"

	"github.com/artemrys/go-all-repos/internal/repo"
	"github.com/google/go-github/github"
)

const clonePrefix = "go-all-repos"

var (
	username   = flag.String("username", "", "Github username")
	reposFlag  = flag.String("repos", "", "Github repos to update, comma separeted")
	reposNames []string
)

func init() {
	flag.Parse()
	if *username == "" {
		log.Fatalf("Username is blank, should be specified.")
	}
	if *reposFlag != "" {
		reposNames = strings.Split(*reposFlag, ",")
	}
}

func main() {
	repos := []*repo.Repo{}
	if len(reposNames) == 0 {
		ctx := context.Background()
		client := github.NewClient(nil)
		opt := &github.RepositoryListOptions{Type: "public"}
		githubRepos, _, err := client.Repositories.List(ctx, *username, opt)
		if err != nil {
			log.Fatalf("Error while getting repositories for %q: %v\n", *username, err)
		}
		for _, r := range githubRepos {
			repos = append(repos, repo.NewFromGithub(r))
		}
	} else {
		for _, r := range reposNames {
			repos = append(repos, repo.New(*username, r))
		}
	}
	for _, r := range repos {
		r.Clone(clonePrefix)
	}
}
