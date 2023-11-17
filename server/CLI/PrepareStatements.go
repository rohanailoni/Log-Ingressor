package CLI

import "database/sql"

func PrepareGeneralStatement(db *sql.DB, query string) (*sql.Stmt, error) {
	stmt, err := db.Prepare(query)
	return stmt, err
}
