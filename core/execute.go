package core

import (
	"dead-simple/compiler"

	"github.com/rs/zerolog/log"
)

func ExecuteStatement(statement compiler.Statement) {
	switch statement.Type {
	case compiler.Insert:
		log.Info().Msg("This is where we would do an insert.")
	case compiler.Select:
		log.Info().Msg("This is where we would do an select")

	}
}
