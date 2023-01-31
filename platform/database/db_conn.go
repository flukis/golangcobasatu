package database

import (
	"github.com/fahmilukis/expense-tracker/app/queries"
	"github.com/jmoiron/sqlx"
)

type Queries struct {
	*queries.ExpenseCategoryQuery
	*queries.UserAccountQuery
}

func OpenDBConn() (*Queries, error) {
	// define database connection
	var (
		db  *sqlx.DB
		err error
	)

	// define connection
	db, err = PsqlConn()
	if err != nil {
		return nil, err
	}

	// return query
	return &Queries{
		&queries.ExpenseCategoryQuery{DB: db},
		&queries.UserAccountQuery{DB: db},
	}, err
}
