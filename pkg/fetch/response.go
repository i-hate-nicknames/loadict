package fetch

type Response struct {
	Results []*Result
	Word    string
}

type Result struct {
	LexicalEntries []*LexicalEntry
	Word           string
	Language       string
}

type LexicalEntry struct {
	LexicalCategory *SingleText
	Entries         []*Entry
}

type Entry struct {
	Senses         []*Sense
	Pronunciations []*Pronunciation
}

type Sense struct {
	Definitions []string
	Examples    []*SingleText
}

type Pronunciation struct {
	PhoneticSpelling string
	Dialects         []string
	PhoneticNotation string
}

type SingleText struct {
	Text string
}
