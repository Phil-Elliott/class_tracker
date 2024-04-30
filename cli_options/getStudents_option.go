package cli_options

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Phil-Elliott/class_tracker/student"
)

func getStudentsCallback(db *sql.DB) error {
	var students []student.Student

	rows, err := db.Query("select name from students")
	if err != nil {
		log.Fatal("An error occured when trying to get the students: ", err)
	}

	defer rows.Close()

	for rows.Next() {
		var s student.Student
		rows.Scan(&s.Name)
		students = append(students, s)
	}

	if len(students) == 0 {
		fmt.Println("You don't currently have any students")
		return nil
	}

	fmt.Println("Here are your current students")

	for _, s := range students {
		fmt.Println(s.Name)
	}

	return nil
}
