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
		repo,
		format,
		tag string
	)
	generateCmd := &cobra.Command{
		Use:          "generate",
		Short:        "Generates changelog for a given from",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error { //nolint[:unparam]
			pullRequests := fetchPRsSinceLastRelease(repo)
			simplifyDepsPRs(pullRequests)
			tpl := Default
			if cmd.Flag("format").Value.String() == "adoc" {
				tpl = DefaultAdoc
			}
			t, err := template.New("changelog").Parse(tpl)
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
	generateCmd.Flags().StringVar(&format, "format", "md", "format of generated release notes")
	generateCmd.Flags().StringVarP(&repo, "repository", "r", "", "repository URL")

	_ = generateCmd.MarkFlagRequired("repository")
	return generateCmd
}

func fetchPRsSinceLastRelease(repoName string) map[string][]github.PullRequest {
	check.RepoFormat(repoName)
	repo := strings.Split(repoName, "/")
	client := github.CreateClient()
	previousRelease, _ := github.LatestReleaseOf(repo[0], repo[1])
	matchingCommit := github.FindMatchingCommit(client, repo, previousRelease)
	prs := github.FindAssociatedPRs(client, repo, matchingCommit)

	prsByLabels := make(map[string][]github.PullRequest)
	for i := range prs {
		pr := prs[i]
		label := "misc"
		if shouldSkipInChangelog(pr.Labels) {
			continue
		}
		if len(pr.Labels) > 0 {
			label = pr.Labels[0]
		}
		prsByLabels[label] = append(prsByLabels[label], pr)
	}

	return prsByLabels
}

const skipLabel = "skip-changelog"

func shouldSkipInChangelog(labels []string) bool {
	for _, label := range labels {
		if label == skipLabel {
			return true
		}
	}
	return false
}

const dependabotPrefix = "build(deps): bump "

func simplifyDepsPRs(prsByLabels map[string][]github.PullRequest) {
	dependencies := prsByLabels["dependencies"]
	sort.SliceStable(dependencies, func(i, j int) bool {
		return strings.Compare(dependencies[i].Title, dependencies[j].Title) < 0
	})

	latestDeps := make(map[string][]github.PullRequest)
	for i := 0; i < len(dependencies); i++ {
		prTitle := strings.Split(strings.TrimPrefix(dependencies[i].Title, dependabotPrefix), " ")
		dep := prTitle[0]
		version := prTitle[4]
		dependencies[i].Title = dep + " to " + version
		latestDeps[dep] = append(latestDeps[dep], dependencies[i])
	}

	latestPrs := make([]github.PullRequest, 0)
	for key := range latestDeps {
		key := key
		sort.SliceStable(latestDeps[key], func(i, j int) bool {
			return latestDeps[key][i].RelatedCommit.CreatedAt.After(latestDeps[key][j].RelatedCommit.CreatedAt.Time)
		})
		latestPrs = append(latestPrs, latestDeps[key][0])
	}

	sort.SliceStable(latestPrs, func(i, j int) bool {
		return strings.Compare(latestPrs[i].Title, latestPrs[j].Title) < 0
	})

	prsByLabels["dependencies"] = latestPrs
}
