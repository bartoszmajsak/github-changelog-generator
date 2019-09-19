package generate

import (
	"os"
	"strings"
	"text/template"

	"github.com/bartoszmajsak/github-changelog-generator/pkg/github"

	"github.com/bartoszmajsak/github-changelog-generator/pkg/check"

	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	var (
		from,
		repo,
		tag string
	)
	generateCmd := &cobra.Command{
		Use:          "generate",
		Short:        "Generates changelog for a given from",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error { //nolint[:unparam]
			pullRequests := fetchRelatedPRs(repo, from)
			t, err := template.New("changelog").Parse(Default)
			if err != nil {
				return err
			}
			if err := t.Execute(os.Stdout, &Changelog{Release: tag, PullRequests: pullRequests}); err != nil {
				return err
			}
			return nil
		},
	}

	generateCmd.Flags().StringVarP(&tag, "tag", "t", "", "tag used for current release")
	generateCmd.Flags().StringVarP(&from, "from", "f", "", "from for which changelog should be generated")
	generateCmd.Flags().StringVarP(&repo, "repository", "r", "", "repository URL")

	_ = generateCmd.MarkFlagRequired("for")
	_ = generateCmd.MarkFlagRequired("tag")
	_ = generateCmd.MarkFlagRequired("repository")
	return generateCmd
}

func fetchRelatedPRs(repoName, ref string) map[string][]github.PullRequest {
	check.RepoFormat(repoName)
	repo := strings.Split(repoName, "/")
	client := github.CreateClient()

	matchingCommit := github.FindMatchingCommit(client, repo, ref)
	prs := github.FindAssociatedPRs(client, repo, matchingCommit)

	prsByLabels := make(map[string][]github.PullRequest)
	for i := range prs {
		pr := prs[i]
		prsByLabels[pr.Labels[0]] = append(prsByLabels[pr.Labels[0]], pr)
	}

	return prsByLabels
}
