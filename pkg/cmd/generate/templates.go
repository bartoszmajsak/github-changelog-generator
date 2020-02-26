package generate

import "github.com/bartoszmajsak/github-changelog-generator/pkg/github"

type Changelog struct {
	Release      string
	PullRequests []github.PullRequest
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

const Default = `
{{- with $prs := (withLabel .PullRequests "kind/enhancement") -}}
{{ if $prs }}
### New features
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (withLabel .PullRequests "kind/bug") -}}
{{ if $prs }}
### Bug fixes
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (withLabel .PullRequests "dependencies") -}}
{{ if $prs }}
### Latest dependencies update
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}}))
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (withLabel .PullRequests "internal/infra") -}}
{{ if $prs }}
### Project infrastructure
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (withLabel .PullRequests "internal/test-infra") -}}
{{ if $prs }}
### Testing
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}
{{ end }}
`

const DefaultAdoc = `
{{- with $prs := (withLabel .PullRequests "kind/enhancement") -}}
{{ if $prs }}
=== New features
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}]), by https://github.com/{{$pr.Author}}[@{{$pr.Author}}]
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (withLabel .PullRequests "kind/bug") -}}
{{ if $prs }}
=== Bug fixes
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}]), by https://github.com/{{$pr.Author}}[@{{$pr.Author}}]
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (withLabel .PullRequests "dependencies") -}}
{{ if $prs }}
=== Latest dependencies update
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}])
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (withLabel .PullRequests "internal/infra") -}}
{{ if $prs }}
=== Project infrastructure
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}]), by https://github.com/{{$pr.Author}}[@{{$pr.Author}}]
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (withLabel .PullRequests "internal/test-infra") -}}
{{ if $prs }}
=== Testing
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}]), by https://github.com/{{$pr.Author}}[@{{$pr.Author}}]
{{- end -}}
{{ end }}
{{ end }}
`
