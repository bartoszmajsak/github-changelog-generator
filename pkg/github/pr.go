package github

func extractLabels(nodes []struct{ Name string }) []string {
	labels := []string{}
	for _, node := range nodes {
		labels = append(labels, node.Name)
	}
	return labels
}

func removeDuplicates(elements []PullRequest) []PullRequest {
	found := map[int]bool{}
	var result []PullRequest

	for v := range elements {
		if !found[elements[v].Number] {
			found[elements[v].Number] = true
			result = append(result, elements[v])
		}
	}
	return result
}
