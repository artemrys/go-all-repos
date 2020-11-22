package config

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
)

// Config contains values to configure runs.
type Config struct {
	// Username is a Github username.
	Username string
	// Repos contains repository names to run against.
	// Multiple values should be separated by comma.
	Repos string
	// RepoNames contains already parsed repository names from Repos.
	RepoNames []string
	// GithubAccessToken is a Github access token.
	GithubAccessToken string
	// dryRun runs everything as normal is set to true.
	// If set to false, does not push changes and does not create a PR.
	DryRun bool
}

// NewFromFlags parses the command-line arguments provided to the program.
// Typically os.Args[0] is provided as 'progname' and os.args[1:] as 'args'.
// Returns the Config in case parsing succeeded, or an error. In any case, the
// output of the flag.Parse is returned in output.
// A special case is usage requests with -h or -help: then the error
// flag.ErrHelp is returned and output will contain the usage message.
func NewFromFlags(progname string, args []string) (*Config, string, error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)

	config := new(Config)
	flags.StringVar(&config.Username, "username", "", "Github username")
	flags.StringVar(&config.Repos, "repos", "", "Github repos to update, comma separeted")
	flags.StringVar(&config.GithubAccessToken, "github-access-token", "", "Github access token")
	flags.BoolVar(&config.DryRun, "dry-run", false, "Dry run (do not push anything if true)")
	if err := flags.Parse(args); err != nil {
		return nil, buf.String(), err
	}
	if config.Username == "" {
		return nil, "", fmt.Errorf("username is blank, should be specified")
	}
	if !config.DryRun && config.GithubAccessToken == "" {
		return nil, "", fmt.Errorf("github-access-token is blank, should be specified")
	}
	if config.Repos != "" {
		config.RepoNames = strings.Split(config.Repos, ",")
	}
	return config, buf.String(), nil
}
