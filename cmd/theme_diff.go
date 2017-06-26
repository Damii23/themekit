package cmd

import (
	"bytes"
	"text/template"
)

type themeDiff struct {
	Created []string
	Updated []string
	Removed []string
}

var themeDiffErrorTmplt = template.Must(template.New("themeDiffError").Parse(`Remote files are inconsistent with manifest
Diff:
{{- if .Created }}
	New Files:
		{{- range .Created }}
		- {{ print . }}
		{{- end }}
{{- end }}
{{- if .Updated }}
	Updated Files:
		{{- range .Updated }}
		- {{ print . }}
		{{- end }}
{{- end }}
{{- if .Removed }}
	Removed Files:
		{{- range .Removed }}
		- {{ print . }}
		{{- end }}
{{- end }}

You can solve this by running theme download and merging the remote changes
using your favourite diff tool or if you are certain about what you are doing
then use the --force flag
`))

func newDiff() *themeDiff {
	return &themeDiff{
		Created: []string{},
		Updated: []string{},
		Removed: []string{},
	}
}

func (diff *themeDiff) Any() bool {
	return len(diff.Created) > 0 || len(diff.Updated) > 0 || len(diff.Removed) > 0
}

func (diff *themeDiff) Error() string {
	var tpl bytes.Buffer
	themeDiffErrorTmplt.Execute(&tpl, diff)
	return tpl.String()
}