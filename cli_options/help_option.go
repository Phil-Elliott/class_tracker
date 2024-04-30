package cli_options

import (
	"database/sql"
	"fmt"
)

func helpCallback(db *sql.DB) error {
	fmt.Println("Please see the list of options below.")
	commands := GetOptions()
	for _, command := range commands {
		fmt.Printf("- %s: %s\n", command.Name, command.Description)
	}

	return nil
}
