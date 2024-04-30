package cli_options

import (
	"database/sql"
	"os"
)

func exitCallback(db *sql.DB) error {
	os.Exit(0)
	return nil
}
