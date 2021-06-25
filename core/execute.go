package core

import (
	"github.com/marianogappa/sqlparser/query"
	"github.com/rs/zerolog/log"
)

func ExecuteStatement(q query.Query, table *Table) {
	switch q.Type {
	case query.Insert:
		table.ExecuteInsert(q)
		log.Info().Msg(table.GetTableSize())
	case query.Select:
		log.Info().Msg("This is where we would do an select")
	}
}
