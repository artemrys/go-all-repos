package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewFromFlags_Correct(t *testing.T) {
	progname := "go-all-repos"
	tests := []struct {
		desc       string
		args       []string
		wantConfig *Config
	}{
		{
			desc: "username with dry run, gofmt action",
			args: []string{"-username", "artemrys", "-dry-run", "-action", "gofmt"},
			wantConfig: &Config{
				Username: "artemrys",
				Action:   "gofmt",
				DryRun:   true,
			},
		},
		{
			desc: "username with github access token and gofmt action",
			args: []string{"-username", "artemrys", "-action", "gofmt", "-github-access-token", "github-access-token"},
			wantConfig: &Config{
				Username:          "artemrys",
				Action:            "gofmt",
				GithubAccessToken: "github-access-token",
			},
		},
		{
			desc: "username with a repo, github access token and action",
			args: []string{"-username", "artemrys", "-repos", "go-all-repos-demo", "-action", "gofmt", "-github-access-token", "github-access-token"},
			wantConfig: &Config{
				Username:          "artemrys",
				Repos:             "go-all-repos-demo",
				RepoNames:         []string{"go-all-repos-demo"},
				Action:            "gofmt",
				GithubAccessToken: "github-access-token",
			},
		},
		{
			desc: "username with multiple repos, github access token and gofmt action",
			args: []string{"-username", "artemrys", "-repos", "go-all-repos-demo,go-all-repos", "-action", "gofmt", "-github-access-token", "github-access-token"},
			wantConfig: &Config{
				Username:          "artemrys",
				Repos:             "go-all-repos-demo,go-all-repos",
				RepoNames:         []string{"go-all-repos-demo", "go-all-repos"},
				Action:            "gofmt",
				GithubAccessToken: "github-access-token",
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			gotConfig, _, err := NewFromFlags(progname, tc.args)
			if diff := cmp.Diff(tc.wantConfig, gotConfig); diff != "" || err != nil {
				t.Errorf("NewFromFlags(_, %v) returned diff (-want, +got):%s\n", tc.args, diff)
				t.Errorf("NewFromFlags(_, %v) returned error %v, wanted nil", tc.args, err)
			}
		})
	}
}

func TestNewFromFlags_Incorrect(t *testing.T) {
	progname := "go-all-repos"
	tests := []struct {
		desc string
		args []string
	}{
		{
			desc: "empty args",
			args: []string{},
		},
		{
			desc: "not dry run, no github access token, gofmt action",
			args: []string{"-username", "artemrys", "-action", "gofmt"},
		},
		{
			desc: "not supported action",
			args: []string{"-username", "artemrys", "-action", "notsupportedaction"},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			_, _, err := NewFromFlags(progname, tc.args)
			if err == nil {
				t.Errorf("NewFromFlags(_, %v) returned error nil, wanted error", tc.args)
			}
		})
	}
}
