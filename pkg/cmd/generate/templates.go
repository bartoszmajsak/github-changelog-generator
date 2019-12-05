package generate

import "github.com/bartoszmajsak/github-changelog-generator/pkg/github"

type Changelog struct {
	Release      string
	PullRequests map[string][]github.PullRequest
}

const Default = `
{{- with $prs := (index .PullRequests "enhancement") -}}
{{ if $prs }}
### New features
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (index .PullRequests "bug") -}}
{{ if $prs }}
### Bug fixes
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (index .PullRequests "dependencies") -}}
{{ if $prs }}
### Dependencies update
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}}))
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (index .PullRequests "infra") -}}
{{ if $prs }}
### Project infrastructure
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (index .PullRequests "test-infra") -}}
{{ if $prs }}
### Testing
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (index .PullRequests "misc") -}}
{{ if $prs }}
### Misc
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}
{{ end }}
`

const DefaultAdoc = `
{{- with $prs := (index .PullRequests "enhancement") -}}
{{ if $prs }}
=== New features
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}]), by https://github.com/{{$pr.Author}}[@{{$pr.Author}}]
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (index .PullRequests "bug") -}}
{{ if $prs }}
=== Bug fixes
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}]), by https://github.com/{{$pr.Author}}[@{{$pr.Author}}]
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (index .PullRequests "dependencies") -}}
{{ if $prs }}
=== Dependencies update
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}])
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (index .PullRequests "infra") -}}
{{ if $prs }}
=== Project infrastructure
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}]), by https://github.com/{{$pr.Author}}[@{{$pr.Author}}]
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (index .PullRequests "test-infra") -}}
{{ if $prs }}
=== Testing
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}]), by https://github.com/{{$pr.Author}}[@{{$pr.Author}}]
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (index .PullRequests "misc") -}}
{{ if $prs }}
=== Misc
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}
{{ end }}
`
