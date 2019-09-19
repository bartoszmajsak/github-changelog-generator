package check

import (
	"fmt"
	"os"
	"strings"
)

// IfError checks if error occurred, logs it and exits
func IfError(err error) {
	if err == nil {
		return
	}
	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func RepoFormat(repo string) {
	repoParts := strings.Split(repo, "/")
	if len(repoParts) != 2 {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("wrong repo format: %s. please make sure it's in owner/repo format", repo))
		os.Exit(1)
	}
}
