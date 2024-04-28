package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Phil-Elliott/class_tracker/cli_options"
	"github.com/Phil-Elliott/class_tracker/utils"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
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
			fmt.Println("invalid command")
		}

		command.Callback()

	}

	// db, err := sql.Open("sqlite3", "students.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal("Failed to connect to the database", err)
	// }

	// log.Println("It connected")
	// os.Exit(0)
}

// - clean the input
// - maybe clean up the code a bit

// - type something to the userconst
// - create a help and exit commands
// - start creating the other commands

// Purpose
// - Track how many classes a student has had out of their current package
// - Track what lesson they were taught
// - Get alert when it was their last class
// - Can provide a text to send to their parent
// - Have option to restart at a new total

// Data structure
// - students table
// - name stirng,
// - classes left
// - current_lesson string,
// - platform string

// - read the current students and maybe also get a list of current students
// - make it so you can add students to the data
// - be able to input class taught and it will decrement by 1 and move to the next lesson (6 in a unit - should be easy to follow)
