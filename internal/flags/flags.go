package flags

import (
	"strings"

	cli "github.com/urfave/cli/v2"
)

// FlagNamesToEnv converts flags to a ENV format
func FlagNamesToEnv(names ...string) []string {
	out := []string{}
	for _, name := range names {
		out = append(out, FlagNameToEnv(name))
	}
	return out
}

// FlagNameToEnv converts a flag to an ENV format
func FlagNameToEnv(name string) string {
	return strings.ReplaceAll(strings.ToUpper(name), "-", "_")
}

// Join mulitple sets of flags
func Join(flags ...[]cli.Flag) []cli.Flag {
	var out []cli.Flag
	for _, f := range flags {
		out = append(out, f...)
	}
	return out
}
