package github

// PullRequest is a type holding relevant information about pull request used in generating changelog
type PullRequest struct {
	RelatedCommit Commit
	Title         string
	Number        int
	Permalink     string
	Author        string
	Labels        []string
}

type Commit struct {
	Hash            string
	MessageHeadline string
	Author          string
}
