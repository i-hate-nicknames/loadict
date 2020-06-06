package main

import (
	"bytes"
	"strings"
	"text/template"
)

const card = `
<div class="word"><span class="type">{{.Word}}<span></div>
{{range .Results}}
{{range .LexicalEntries}}
<div class="lex-entry">

	<div class="lex-category">{{ .LexicalCategory.Text }}</div>
	<div class="entries">
	{{range .Entries}}
	{{range .Senses}}
	<div class="sense">
		<div class="definitions">
		{{range .Definitions}}
			<div class="definition">
				{{.}}
			</div>
		{{end}}
		</div>
		<div class="examples">
		{{range .Examples}}
			<div class="example">
				{{.Text}}
			</div>
		{{end}}
		</div>
	</div>
	{{end}}
	{{end}}
	</div>
	<div class="pronunciation">
		<span class="type">IPA: </span>
		<span class="value">{{.RenderPronunciations}}<span>
	</div>

</div>
{{end}}
{{end}}
`

func renderCard(data *Response) (string, error) {
	tmpl, err := template.New("anki").Parse(card)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (le *LexicalEntry) RenderPronunciations() string {
	results := make([]string, 0)
	for _, pron := range le.Pronunciations {
		if pron.PhoneticNotation == "IPA" {
			results = append(results, pron.PhoneticSpelling)
		}
	}
	return strings.Join(results, ",")
}
