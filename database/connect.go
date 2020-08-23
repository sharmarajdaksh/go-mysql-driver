package database

import (
	"database/sql"

	// Documentation
	_ "github.com/go-sql-driver/mysql"
)

// DBConnection is the shared DB connection used by all modules
var DBConnection *sql.DB
