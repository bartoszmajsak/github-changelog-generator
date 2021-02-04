package generate

import "github.com/bartoszmajsak/github-changelog-generator/pkg/github"

type Changelog struct {
	Release      string
	Areas        map[string]string
	PullRequests []github.PullRequest
}

func Contains(s []string, es ...string) bool {
	for _, e := range es {
		eFound := false
		for _, a := range s {
			if a == e {
				eFound = true
			}
		}
		if !eFound {
			return false
		}
	}
	return true
}

const Default = `
{{- with $changeLog := . -}}
{{ range $areaName, $areaLabel := .Areas }}
### {{$areaName}}
{{- with $prs := (withLabels $changeLog.PullRequests "kind/enhancement" $areaLabel) -}}
{{ if $prs }}
#### New features
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (withLabels $changeLog.PullRequests "kind/bug" $areaLabel) -}}
{{ if $prs }}
#### Bug fixes
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{- end -}}
{{ end }}
{{ end }}
{{ end }}
{{- end -}}

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
==== New features
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}]), by https://github.com/{{$pr.Author}}[@{{$pr.Author}}]
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (withLabel .PullRequests "kind/bug") -}}
{{ if $prs }}
==== Bug fixes
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}]), by https://github.com/{{$pr.Author}}[@{{$pr.Author}}]
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (withLabel .PullRequests "dependencies") -}}
{{ if $prs }}
==== Latest dependencies update
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}])
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (withLabel .PullRequests "internal/infra") -}}
{{ if $prs }}
==== Project infrastructure
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}]), by https://github.com/{{$pr.Author}}[@{{$pr.Author}}]
{{- end -}}
{{ end }}
{{ end }}

{{- with $prs := (withLabel .PullRequests "internal/test-infra") -}}
{{ if $prs }}
==== Testing
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}]), by https://github.com/{{$pr.Author}}[@{{$pr.Author}}]
{{- end -}}
{{ end }}
{{ end }}
`
