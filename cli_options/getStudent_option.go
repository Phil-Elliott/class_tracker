package cli_options

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Phil-Elliott/class_tracker/student"
)

func getStudentCallback(db *sql.DB) error {
	fmt.Println("What is your students name?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	rows, err := db.Query("Select * from students where name is ?", text)
	if err != nil {
		log.Fatal("There was an error getting the student with the name you chose: ", err)
	}

	students := []student.Student{}

	for rows.Next() {
		var s student.Student
		rows.Scan(&s.Name, &s.ClassesLeft, &s.CurrentLesson, &s.Platform)
		students = append(students, s)
	}

	if len(students) == 0 {
		fmt.Println("You don't have any students with that name")
	} else if len(students) == 1 {
		student.PrintStudentsData(students[0])
	} else {
		fmt.Println("You have multiple students with that name")
		fmt.Println("Here is all of the info on each student below.")
		for _, s := range students {
			student.PrintStudentsData(s)
		}
	}

	return nil
}
