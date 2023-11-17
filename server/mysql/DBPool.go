package mysql

import "github.com/jmoiron/sqlx"

type DBPool struct {
	*sqlx.DB
}
