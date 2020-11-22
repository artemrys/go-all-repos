package helpers

import (
	"log"
	"os/exec"
)

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
