package generate

import (
	"os"
	"regexp"
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
			dependencies, otherPRs := extractDepPRs(pullRequests)
			dependencies = simplifyDepsPRs(dependencies)
			tpl := Default
			if cmd.Flag("format").Value.String() == "adoc" {
				tpl = DefaultAdoc
			}
			t, err := template.New("changelog").Funcs(map[string]interface{}{
				"withLabel": func(prs []github.PullRequest, label string) []github.PullRequest {
					prsWithLabel := make([]github.PullRequest, 0)
					for i := range prs {
						pr := &prs[i]
						if Contains(pr.Labels, label) {
							prsWithLabel = append(prsWithLabel, *pr)
						}
					}
					return prsWithLabel
				},
			}).Parse(tpl)
			if err != nil {
				return err
			}
			if err := t.Execute(os.Stdout, &Changelog{Release: tag, PullRequests: append(otherPRs, dependencies...)}); err != nil {
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

func fetchPRsSinceLastRelease(repoName string) []github.PullRequest {
	check.RepoFormat(repoName)
	repo := strings.Split(repoName, "/")
	client := github.CreateClient()
	previousRelease, _ := github.LatestReleaseOf(repo[0], repo[1])
	matchingCommit := github.FindMatchingCommit(client, repo, previousRelease)
	prs := github.FindAssociatedPRs(client, repo, matchingCommit)

	filteredPRs := make([]github.PullRequest, 0)
	for i := range prs {
		pr := prs[i]
		if shouldSkipInChangelog(pr.Labels) {
			continue
		}
		filteredPRs = append(filteredPRs, pr)
	}

	return filteredPRs
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

var dependabotPrefix = regexp.MustCompile(`^build\(deps\): [B|b]ump `)

func simplifyDepsPRs(dependencies []github.PullRequest) []github.PullRequest {
	sort.SliceStable(dependencies, func(i, j int) bool {
		return strings.Compare(dependencies[i].Title, dependencies[j].Title) < 0
	})

	latestDeps := make(map[string][]github.PullRequest)
	for i := 0; i < len(dependencies); i++ {
		strippedPrefix := dependabotPrefix.ReplaceAllString(dependencies[i].Title, "")
		prTitle := strings.Split(strippedPrefix, " ")
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

	return latestPrs
}

func extractDepPRs(prs []github.PullRequest) (depPRs, otherPRs []github.PullRequest) {
	depPRs = make([]github.PullRequest, 0)
	otherPRs = make([]github.PullRequest, 0)
	for i := range prs {
		pr := &prs[i]
		if Contains(pr.Labels, "dependencies") {
			depPRs = append(depPRs, prs[i])
		} else {
			otherPRs = append(otherPRs, prs[i])
		}
	}
	return
}
