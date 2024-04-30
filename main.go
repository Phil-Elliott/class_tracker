package main

import (
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	startRepl()

}

// - create students command (shows a list of students)
// - create addStudent command (creates a student)
// - create delStudent command (deletes a student)
// - create updateStudent command (updates a students info)
// - create taughtClass command - decrements class cound by 1 and increments class by 1

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
