package generate

import "github.com/bartoszmajsak/github-changelog-generator/pkg/github"

type Changelog struct {
	Release      string
	Areas        map[string]string
	PullRequests []github.PullRequest
}

type ChangeGroup struct {
	Title        string
	PullRequests []github.PullRequest
}

// TODO withLabels moved here

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

const ChangeSection = `{{- if .PullRequests -}}
#### {{ .Title }}
{{range $pr := .PullRequests -}}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}})), by [@{{$pr.Author}}](https://github.com/{{$pr.Author}})
{{ end -}}
{{- end -}}`

const ChangeSectionAdoc = `{{- if .PullRequests -}}
==== {{ .Title }}
{{range $pr := .PullRequests -}}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}]), by https://github.com/{{$pr.Author}}[@{{$pr.Author}}]
{{ end -}}
{{- end -}}`

const Default = `
{{- range $areaName, $areaLabel := .Areas -}}
{{- $bugs := (withLabels $.PullRequests "kind/bug" $areaLabel) -}}
{{- $features := (withLabels $.PullRequests "kind/enhancement" $areaLabel) -}}
{{- if or $bugs $features -}}

### {{ $areaName }}

{{ template "section" (combine $features "New features") }}
{{ template "section" (combine $bugs "Bugs") }}
{{ end -}}
{{- end -}}

{{- with $prs := (withLabels .PullRequests "dependencies") -}}
{{- if $prs -}}
## Latest dependencies update
{{range $pr := $prs }}
 * {{$pr.Title}} ([#{{$pr.Number}}]({{$pr.Permalink}}))
{{- end -}}
{{- end -}}
{{- end }}
`

const DefaultAdoc = `
{{- range $areaName, $areaLabel := .Areas -}}
{{- $bugs := (withLabels $.PullRequests "kind/bug" $areaLabel) -}}
{{- $features := (withLabels $.PullRequests "kind/enhancement" $areaLabel) -}}
{{- if or $bugs $features -}}

=== {{ $areaName }}

{{ template "section" (combine $features "New features") }}
{{ template "section" (combine $bugs "Bugs") }}
{{ end -}}
{{- end -}}

{{- with $prs := (withLabels .PullRequests "dependencies") -}}
{{- if $prs -}}
== Latest dependencies update
{{range $pr := $prs }}
 * {{$pr.Title}} ({{$pr.Permalink}}[#{{$pr.Number}}]), by https://github.com/{{$pr.Author}}[@{{$pr.Author}}]
{{- end -}}
{{- end -}}
{{- end }}
`
