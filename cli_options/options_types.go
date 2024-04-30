package cli_options

import "database/sql"

type cli_option struct {
	Name        string
	Description string
	Callback    func(db *sql.DB) error
}
