package main

type response struct {
	results []*Result
}

type Result struct {
	LexicalEntries []*LexicalEntry
	Word           string
}

type LexicalEntry struct {
	LexicalCategory string
	Entries         []*Entry
	Pronunciations  []*Pronunciation
}

type Entry struct {
	GrammaticalFeatures []*GrammaticalFeature
	Senses              []*Sense
}

type GrammaticalFeature struct {
	Text string
	Type string
}

type Sense struct {
	Definitions []string
	Examples    []*Example
}

type Example struct {
	Text string
}

type Pronunciation struct {
	PhoneticSpelling string
	Dialects         []string
}
