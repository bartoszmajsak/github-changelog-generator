package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/bartoszmajsak/github-changelog-generator/pkg/github"

	"github.com/bartoszmajsak/github-changelog-generator/pkg/cmd/version"
	"github.com/bartoszmajsak/github-changelog-generator/pkg/config"
	"github.com/bartoszmajsak/github-changelog-generator/pkg/format"
	v "github.com/bartoszmajsak/github-changelog-generator/version"

	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	var configFile string
	releaseInfo := make(chan string, 1)

	rootCmd := &cobra.Command{
		Use: "cmd",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error { //nolint[:unparam]
			if v.Released() {
				go func() {
					latestRelease, _ := github.LatestRelease()
					if !github.IsLatestRelease(latestRelease) {
						releaseInfo <- fmt.Sprintf("WARN: you are using %s which is not the latest release (newest is %s).\n"+
							"Follow release notes for update info https://github.com/Maistra/istio-workspace/releases/latest", v.Version, latestRelease)
					} else {
						releaseInfo <- ""
					}
				}()
			}
			return config.SetupConfigSources(configFile, cmd.Flag("config").Changed)
		},
		RunE: func(cmd *cobra.Command, args []string) error { //nolint[:unparam]
			shouldPrintVersion, _ := cmd.Flags().GetBool("version")
			if shouldPrintVersion {
				version.PrintVersion()
			} else {
				// actual logic
				fmt.Print(cmd.UsageString())
			}
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			if v.Released() {
				timer := time.NewTimer(2 * time.Second)
				select {
				case release := <-releaseInfo:
					fmt.Println(release)
				case <-timer.C:
					// do nothing, just timeout
				}
			}
			close(releaseInfo)
			return nil
		},
	}

	rootCmd.PersistentFlags().
		StringVarP(&configFile, "config", "c", ".ghc.config.yaml",
			fmt.Sprintf("config file (supported formats: %s)", strings.Join(config.SupportedExtensions(), ", ")))
	rootCmd.Flags().Bool("version", false, "prints the version number of ike cli")
	rootCmd.PersistentFlags().String("help-format", "standard", "prints help in asciidoc table")
	if err := rootCmd.PersistentFlags().MarkHidden("help-format"); err != nil {
		fmt.Printf("failed while trying to hide a flag: %s\n", err)
	}

	format.EnhanceHelper(rootCmd)
	format.RegisterTemplateFuncs()

	return rootCmd
}
