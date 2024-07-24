package db

import "database/sql"

type DatabaseStructure interface {
	GetDatabasePoll() (*sql.DB, error)
}
