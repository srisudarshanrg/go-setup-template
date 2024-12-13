package functions

import "database/sql"

var db *sql.DB

// DBAccess provides the handlers with access to the database
func DBAccessFunctions(dbAccess *sql.DB) {
	db = dbAccess
}
