package cli

import (
	"dead-simple/core"
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/marianogappa/sqlparser"
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
	table := &core.Table{}
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
			if strings.HasPrefix(s, ".") {
				log.Error().Msg(fmt.Sprintf("Unrecognized command '%s'", s))
				break
			}
			query, err := sqlparser.Parse(s)
			if err == nil {
				core.ExecuteStatement(query, table)
			} else {
				log.Error().Msg(fmt.Sprintf("Unrecognized keyword at start of '%s'", s))
			}
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
