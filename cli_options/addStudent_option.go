package cli_options

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
)

func addStudentCallback(db *sql.DB) error {
	fmt.Println("What is the students name?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("What platform will you be teaching the student on?")
	scanner.Scan()
	platform := scanner.Text()

	fmt.Println("What lesson will the student start on?")
	scanner.Scan()
	currentLesson := scanner.Text()

	fmt.Println("How many lessons does the student have paid for")
	scanner.Scan()
	lessonCount := scanner.Text()

	stmt, err := db.Prepare("insert into students (name, platform, current_lesson, classes_left) values(?, ?, ?, ?)")

	if err != nil {
		log.Fatal("Prepare insert statement failed: ", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(name, platform, currentLesson, lessonCount)
	if err != nil {
		log.Fatal("Insert statement failed: ", err)
	}

	fmt.Println("Here is the students info: Naem: ", name, " Platform: ", platform, " Current lesson: ", currentLesson, " Lesson Count ", lessonCount)
	return nil
}

// Need to clean input to accept lower and upper case answers
