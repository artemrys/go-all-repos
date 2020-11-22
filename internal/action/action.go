// Package action declares Action interface.
package action

import (
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
