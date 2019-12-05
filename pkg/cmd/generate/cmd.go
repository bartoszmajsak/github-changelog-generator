package generate

import (
	"os"
	"sort"
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
			sortDependencyPRs(pullRequests)
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

	generateCmd.Flags().StringVarP(&tag, "tag", "t", "UNRELEASED", "tag used for current release")
	generateCmd.Flags().StringVarP(&from, "from", "f", "", "from for which changelog should be generated")
	generateCmd.Flags().StringVarP(&repo, "repository", "r", "", "repository URL")

	_ = generateCmd.MarkFlagRequired("repository")
	_ = generateCmd.MarkFlagRequired("from")
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
		label := "misc"
		if len(pr.Labels) > 0 {
			label = pr.Labels[0]
		}
		prsByLabels[label] = append(prsByLabels[label], pr)
	}

	return prsByLabels
}

func sortDependencyPRs(prsByLabels map[string][]github.PullRequest) {
	dependencies := prsByLabels["dependencies"]
	sort.SliceStable(dependencies, func(i, j int) bool {
		return strings.Compare(dependencies[i].Title, dependencies[j].Title) < 0
	})
}
