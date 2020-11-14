package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/google/go-github/github"
)

// Repo holds info about repository.
type Repo struct {
	Name string
	URL  string
}

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

func buildRepoURL(username, repo string) string {
	return fmt.Sprintf("https://github.com/%s/%s", username, repo)
}

func main() {
	repos := []*Repo{}
	if len(reposNames) == 0 {
		ctx := context.Background()
		client := github.NewClient(nil)
		opt := &github.RepositoryListOptions{Type: "public"}
		githubRepos, _, err := client.Repositories.List(ctx, *username, opt)
		if err != nil {
			log.Fatalf("Error while getting repositories for %q: %v\n", *username, err)
		}
		for _, r := range githubRepos {
			repos = append(repos, &Repo{
				Name: *r.Name,
				URL:  *r.CloneURL,
			})
		}
	} else {
		for _, repo := range reposNames {
			repos = append(repos, &Repo{
				Name: repo,
				URL:  buildRepoURL(*username, repo),
			})
		}
	}
	for _, repo := range repos {
		log.Printf("Cloning repo %q\n", repo.Name)
		if _, err := exec.Command(
			"git",
			"clone",
			repo.URL,
			fmt.Sprintf("go-all-repos-%s", repo.Name),
			"--depth",
			"1",
		).Output(); err != nil {
			log.Printf("Error while cloning repo %q: %v\n", repo.Name, err)
		}
	}
}
