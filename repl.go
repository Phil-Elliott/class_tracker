package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Phil-Elliott/class_tracker/cli_options"
	"github.com/Phil-Elliott/class_tracker/utils"
)

func startRepl() {
	db, err := sql.Open("sqlite3", "students.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter a command.")
		scanner.Scan()
		text := scanner.Text()

		cleanedText := utils.CleanInput(text)
		if len(cleanedText) == 0 {
			continue
		}
		commands := cli_options.GetOptions()
		command, ok := commands[cleanedText]

		if !ok {
			fmt.Println("invalid command", cleanedText)
			continue
		}

		command.Callback(db)

	}
}
