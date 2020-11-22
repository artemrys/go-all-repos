// Package action declares Action interface.
package action

import (
	"log"
	"os/exec"

	"github.com/artemrys/go-all-repos/internal/config"
	"github.com/artemrys/go-all-repos/internal/repo"
	"github.com/google/go-github/github"
)

// Action declares interface for all actions.
type Action interface {
	Do()
}

// NewActionFunc declares a function to return a new action.
type NewActionFunc func(*repo.Repo, *github.Client, *config.Config) Action

// ActionsMap returns a mapping from a string to an Action.
func ActionsMap() map[string]NewActionFunc {
	return map[string]NewActionFunc{
		"gofmt": func(repo *repo.Repo, githubClient *github.Client, config *config.Config) Action {
			return NewGoFmtAction(repo, githubClient, config)
		},
	}
}

// NewAction returns new Action.
func NewAction(repo *repo.Repo, githubClient *github.Client, config *config.Config) Action {
	actionsMap := ActionsMap()
	newActionFunc := actionsMap[config.Action]
	return newActionFunc(repo, githubClient, config)
}

// Run is a helper function to execute commands.
func Run(path string, args []string, dir string) {
	cmd := &exec.Cmd{
		Path: path,
		Args: args,
		Dir:  dir,
	}
	out, err := cmd.Output()
	if err != nil {
		log.Printf("cmd.Stdout %s\n", cmd.Stdout)
		log.Printf("cmd.Stderr %s\n", cmd.Stderr)
		log.Printf("Error while executing command %q %q: %v\n", cmd.Path, cmd.Args, err)
	} else {
		log.Printf("Output when running command %q %q: %s\n", cmd.Path, cmd.Args, out)
	}
}

// RunGit is a helper function to execute Git commands.
func RunGit(gitArgs []string, dir string) {
	args := []string{}
	args = append(args, "/usr/bin/git")
	for _, a := range gitArgs {
		args = append(args, a)
	}
	Run("/usr/bin/git", args, dir)
}
