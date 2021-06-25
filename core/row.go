package core

import "github.com/marianogappa/sqlparser/query"

type Row struct {
	id       string
	username string
	password string
}

func RowFromQuery(q query.Query) Row {
	r := Row{}
	for i, field := range q.Fields {
		if field == "id" {
			r.id = q.Inserts[0][i]
		} else if field == "username" {
			r.username = q.Inserts[0][i]
		} else if field == "password" {
			r.password = q.Inserts[0][i]
		}
	}

	return r
}
