package check

import (
	"fmt"
	"os"
	"strings"
)

const RedFmt = "\x1b[31;1m%s\x1b[0m\n"

// IfError checks if error occurred, logs it and exits
func IfError(err error) {
	if err == nil {
		return
	}
	fmt.Printf(RedFmt, fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// RepoFormat checks if repository string conforms with expected format, i.e. owner/repo
func RepoFormat(repo string) {
	repoParts := strings.Split(repo, "/")
	if len(repoParts) != 2 {
		fmt.Printf(RedFmt, fmt.Sprintf("wrong repo format: %s. please make sure it's in owner/repo format", repo))
		os.Exit(1)
	}
}
