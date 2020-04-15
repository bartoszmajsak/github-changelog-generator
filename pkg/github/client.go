package github

import (
	"github.com/bartoszmajsak/github-changelog-generator/pkg/check"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"

	"context"
	"fmt"
	"os"
)

func CreateClient() *githubv4.Client {
	const ghToken = "GHC_GITHUB_TOKEN" //nolint[:gosec] G101: Potential hardcoded credential
	var src oauth2.TokenSource
	if token, exists := os.LookupEnv(ghToken); exists {
		src = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	} else {
		fmt.Printf(check.RedFmt, "Missing GitHub token.")
		fmt.Println("Please set it using env variable " + ghToken +
			", otherwise you might not be able to query the data due to rate limits.\n" +
			"For more information on how to create one see https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line")
	}
	httpClient := oauth2.NewClient(context.Background(), src)
	return githubv4.NewClient(httpClient)
}

type AssociatedPRsQuery struct {
	Repository struct {
		DefaultBranch struct {
			Target struct {
				Commit struct {
					Oid     string
					History struct {
						Nodes []struct {
							Oid             string
							MessageHeadline string
							AuthoredDate    githubv4.DateTime
							Author          struct {
								User struct {
									Login string
								}
							}
							AssociatedPullRequests struct {
								Nodes []struct {
									Title     string
									Number    int
									Permalink string
									Author    struct {
										Login string
									}
									Labels struct {
										Nodes []struct {
											Name string
										}
									} `graphql:"labels(first: 8)"`
								}
							} `graphql:"associatedPullRequests(first: 4)"`
						}
					} `graphql:"history(since: $createdAt)"`
				} `graphql:"... on Commit"`
			} `graphql:"target"`
		} `graphql:"defaultBranchRef"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

type MatchingCommitQuery struct {
	Repository struct {
		Object struct {
			Commit struct {
				Oid           string
				CommittedDate githubv4.DateTime
			} `graphql:"... on Commit"`
		} `graphql:"object(expression: $expr)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

func FindMatchingCommit(client *githubv4.Client, repo []string, ref string) MatchingCommitQuery {
	var matchingCommit MatchingCommitQuery
	err := client.Query(context.Background(), &matchingCommit, map[string]interface{}{
		"owner": githubv4.String(repo[0]),
		"name":  githubv4.String(repo[1]),
		"expr":  githubv4.String(ref),
	})
	check.IfError(err)
	return matchingCommit
}

func FindAssociatedPRs(client *githubv4.Client, repo []string, matchingCommit MatchingCommitQuery) []PullRequest {
	var associatedPRs AssociatedPRsQuery
	err := client.Query(context.Background(), &associatedPRs, map[string]interface{}{
		"owner":     githubv4.String(repo[0]),
		"name":      githubv4.String(repo[1]),
		"createdAt": githubv4.GitTimestamp{Time: matchingCommit.Repository.Object.Commit.CommittedDate.Time},
	})
	check.IfError(err)

	var (
		prs          []PullRequest
		unlabeledPRs bool
	)
	for _, node := range associatedPRs.Repository.DefaultBranch.Target.Commit.History.Nodes {
		if node.Oid != matchingCommit.Repository.Object.Commit.Oid && node.MessageHeadline != "release: next iteration" {
			for _, pr := range node.AssociatedPullRequests.Nodes {
				if len(pr.Labels.Nodes) == 0 {
					unlabeledPRs = true
				}
				prs = append(prs, PullRequest{
					RelatedCommit: Commit{
						Hash:            node.Oid,
						Author:          node.Author.User.Login,
						MessageHeadline: node.MessageHeadline,
						CreatedAt:       node.AuthoredDate,
					},
					Title:     pr.Title,
					Number:    pr.Number,
					Permalink: pr.Permalink,
					Author:    pr.Author.Login,
					Labels:    extractLabels(pr.Labels.Nodes),
				})
			}
		}
	}
	withoutDuplicates := removeDuplicates(prs)
	if unlabeledPRs {
		_, _ = fmt.Fprint(os.Stderr, "#### Found unlabeled PRs\n\n")
		for i := range withoutDuplicates {
			_, _ = fmt.Fprintf(os.Stderr, "* %s\n", withoutDuplicates[i].Permalink)
		}
	}
	return withoutDuplicates
}
