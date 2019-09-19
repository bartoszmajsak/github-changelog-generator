package main

import (
	"github.com/bartoszmajsak/github-changelog-generator/pkg/check"
	"github.com/bartoszmajsak/github-changelog-generator/pkg/cmd"
	"github.com/bartoszmajsak/github-changelog-generator/pkg/cmd/generate"
	"github.com/bartoszmajsak/github-changelog-generator/pkg/cmd/version"
)

func main() {
	rootCmd := cmd.NewCmd()
	rootCmd.AddCommand(generate.NewCmd(), version.NewCmd())
	if err := rootCmd.Execute(); err != nil {
		check.IfError(err)
	}
}
