package cli

import "github.com/c-bata/go-prompt"

type suggestionType int

const (
	Commands suggestionType = iota
	Keywords
)

var commandSuggestions = []prompt.Suggest{
	{Text: ".quit", Description: "Quit/Exit application"},
	{Text: ".exit", Description: "Quit/Exit application"},
}

var keywordSuggestions = []prompt.Suggest{
	{Text: "insert", Description: "Insert statement to add data to a table"},
	{Text: "select", Description: "select statement to read data from a table"},
}

var suggestionsMap = map[suggestionType][]prompt.Suggest{
	Commands: commandSuggestions,
	Keywords: keywordSuggestions,
}
