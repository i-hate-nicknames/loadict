package fetch

import (
	"bytes"
	"strings"
	"text/template"
)

const cardTemplate = `
<div class="word"><span class="type">{{.Word}}<span></div>
{{range .Results}}
{{range .LexicalEntries}}
<div class="lex-entry">
	<div class="lex-category">Type: {{ .LexicalCategory.Text }}</div>
	<div class="entries">
	{{range .Entries}}
	<div class="pronunciation">
	<span class="type">IPA: </span>
	<span class="value">{{.RenderPronunciations}}<span>
	</div>
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
		<div><span class="type">Examples:</span></div>
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
</div>
{{end}}
{{end}}
`

func renderCard(data *Response) (string, error) {
	tmpl, err := template.New("anki").Parse(cardTemplate)
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

// todo: add dot at the end of every example

func (e *Entry) RenderPronunciations() string {
	results := make([]string, 0)
	for _, pron := range e.Pronunciations {
		if pron.PhoneticNotation == "IPA" {
			results = append(results, pron.PhoneticSpelling)
		}
	}
	return strings.Join(results, ",")
}
