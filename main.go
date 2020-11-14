package main

import (
	"context"
	"flag"
	"log"
	"strings"

	"github.com/artemrys/go-all-repos/internal/action"
	"github.com/artemrys/go-all-repos/internal/repo"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const clonePrefix = "go-all-repos"

var (
	username          = flag.String("username", "", "Github username")
	reposFlag         = flag.String("repos", "", "Github repos to update, comma separeted")
	githubAccessToken = flag.String("github-access-token", "", "Github access token")
	dryRun            = flag.Bool("dry-run", false, "Dry run (do not push anything if true)")
	reposNames        []string
)

func init() {
	flag.Parse()
	if *username == "" {
		log.Fatalln("Username is blank, should be specified.")
	}
	if *githubAccessToken == "" {
		log.Fatalln("Github access token is blank, should be specified.")
	}
	if *reposFlag != "" {
		reposNames = strings.Split(*reposFlag, ",")
	}
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *githubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	githubClient := github.NewClient(tc)
	repos := []*repo.Repo{}
	if len(reposNames) == 0 {
		ctx := context.Background()
		opt := &github.RepositoryListOptions{Type: "public"}
		githubRepos, _, err := githubClient.Repositories.List(ctx, *username, opt)
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
		e := action.NewGoFmtAction(r, githubClient, *dryRun)
		e.Do()
	}
}
