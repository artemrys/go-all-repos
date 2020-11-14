package repo

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/github"
)

func TestNew(t *testing.T) {
	t.Parallel()
	username := "artemrys"
	repo := "go-all-repos"

	want := &Repo{
		Name:     "go-all-repos",
		CloneURL: "https://github.com/artemrys/go-all-repos.git",
	}

	got := New(username, repo)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("New(%s, %s) returned diff (-want, +got):%s\n", username, repo, diff)
	}
}

func TestNewFromGithub(t *testing.T) {
	t.Parallel()
	githubRepoName := "go-all-repos"
	githubRepoCloneURL := "https://github.com/artemrys/go-all-repos.git"
	githubRepo := github.Repository{
		Name:     &githubRepoName,
		CloneURL: &githubRepoCloneURL,
	}

	want := &Repo{
		Name:     "go-all-repos",
		CloneURL: "https://github.com/artemrys/go-all-repos.git",
	}

	got := NewFromGithub(&githubRepo)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("NewFromGithub(%v) returned diff (-want, +got):%s\n", githubRepo, diff)
	}
}
