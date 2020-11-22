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
			desc: "username with dry run",
			args: []string{"-username", "artemrys", "-dry-run"},
			wantConfig: &Config{
				Username: "artemrys",
				DryRun:   true,
			},
		},
		{
			desc: "username with github access token",
			args: []string{"-username", "artemrys", "-github-access-token", "github-access-token"},
			wantConfig: &Config{
				Username:          "artemrys",
				GithubAccessToken: "github-access-token",
			},
		},
		{
			desc: "username with a repo and github access token",
			args: []string{"-username", "artemrys", "-repos", "go-all-repos-demo", "-github-access-token", "github-access-token"},
			wantConfig: &Config{
				Username:          "artemrys",
				Repos:             "go-all-repos-demo",
				RepoNames:         []string{"go-all-repos-demo"},
				GithubAccessToken: "github-access-token",
			},
		},
		{
			desc: "username with multiple repos and github access token",
			args: []string{"-username", "artemrys", "-repos", "go-all-repos-demo,go-all-repos", "-github-access-token", "github-access-token"},
			wantConfig: &Config{
				Username:          "artemrys",
				Repos:             "go-all-repos-demo,go-all-repos",
				RepoNames:         []string{"go-all-repos-demo", "go-all-repos"},
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
			desc: "not dry run, no github access token",
			args: []string{"-username", "artemrys"},
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
