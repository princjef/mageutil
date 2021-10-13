package shellcmd

import (
	"fmt"
	"os"
	"os/exec"
)

// Option are functions that are passed into Run to modify the behaviour of the executed command.
type Option func(*exec.Cmd)

// WithDir returns an Option to specify the working directory of the command.
// If dir is the empty string, runs the command in the calling process's current directory.
func WithDir(dir string) Option {
	return func(cmd *exec.Cmd) {
		cmd.Dir = dir
	}
}

// WithEnv returns an Option to specify the environment of the command.
// First retrieves the value of the environment variable named by the key, if not set, uses the def value.
func WithEnv(key, def string) Option {
	return func(cmd *exec.Cmd) {
		val, found := os.LookupEnv(key)
		if !found {
			val = def
		}
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%v", key, val))
	}
}
