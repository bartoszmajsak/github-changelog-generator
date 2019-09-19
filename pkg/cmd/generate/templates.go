package generate

import "github.com/bartoszmajsak/github-changelog-generator/pkg/github"

type Changelog struct {
	Release      string
	PullRequests map[string][]github.PullRequest
}

const Default = `
# Highlights of {{.Release }} release

## New features

{{- with $prs := (index .PullRequests "enhancement") -}}
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}

## Bug fixes

{{- with $prs := (index .PullRequests "bug") -}}
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}

## Dependencies update

{{- with $prs := (index .PullRequests "dependencies") -}}
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}}))
{{- end -}}
{{ end }}

## Project infrastructure

{{- with $prs := (index .PullRequests "infra") -}}
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}

## Testing

{{- with $prs := (index .PullRequests "test-infra") -}}
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}
`
