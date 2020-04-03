package main

type Response struct {
	Results []*Result
	word    string
}

type Result struct {
	LexicalEntries []*LexicalEntry
	Word           string
	Language       string
}

type LexicalEntry struct {
	LexicalCategory *SingleText
	Entries         []*Entry
	Pronunciations  []*Pronunciation
}

type Entry struct {
	Senses []*Sense
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
