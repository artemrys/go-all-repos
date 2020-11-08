package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/exec"

	"github.com/google/go-github/github"
)

var username = flag.String("username", "", "Github username")

func init() {
	flag.Parse()
	if *username == "" {
		log.Fatalf("Username is blank, should be specified.")
	}
}

func main() {
	ctx := context.Background()
	client := github.NewClient(nil)
	opt := &github.RepositoryListOptions{Type: "public"}
	repos, _, err := client.Repositories.List(ctx, *username, opt)
	if err != nil {
		log.Fatalf("Error while getting repositories for %q: %v\n", *username, err)
	}
	for _, repo := range repos {
		log.Printf("Cloning repo %q\n", *repo.Name)
		if _, err := exec.Command(
			"git",
			"clone",
			*repo.CloneURL,
			fmt.Sprintf("go-all-repos-%s", *repo.Name),
			"--depth",
			"1",
		).Output(); err != nil {
			log.Printf("Error while cloning repo %q: %v\n", *repo.Name, err)
		}
	}
}
