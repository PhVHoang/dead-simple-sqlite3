package cli

import (
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/rs/zerolog/log"
)

func StartPrompt(file string) {
	p := prompt.New(
		getExecutor(file),
		completer,
		prompt.OptionTitle("Dead simple sqlite3"),
		prompt.OptionPrefix("sql > "),
	)
	p.Run()
}

func getExecutor(file string) func(string) {
	return func(s string) {
		s = strings.TrimSpace(s)
		s = strings.ToLower(s)
		switch s {
		case "":
			return
		case ".quit", ".exit":
			log.Info().Msg("Goodbye!")
			os.Exit(0)
		default:
			// TODO: Proccess commands here
		}
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{}
	for _, s := range suggestionsMap {
		suggestions = append(suggestions, s...)
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}
