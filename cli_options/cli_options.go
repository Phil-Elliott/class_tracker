package cli_options

func GetOptions() map[string]cli_option {
	return map[string]cli_option{
		"getstudents": {
			Name:        "getstudents",
			Description: "Get a list of current students.",
			Callback:    getStudentsCallback,
		},
		"getstudent": {
			Name:        "getstudent",
			Description: "Get all of the info on a student",
			Callback:    getStudentCallback,
		},
		"addstudent": {
			Name:        "addstudent",
			Description: "Add a new student",
			Callback:    addStudentCallback,
		},
		"deletestudent": {
			Name:        "deletestudent",
			Description: "Deletes a student",
			Callback:    getStudentsCallback,
		},
		"updatestudent": {
			Name:        "updatestudent",
			Description: "Updates the students info",
			Callback:    getStudentsCallback,
		},
		"updateClass": {
			Name:        "updateClass",
			Description: "Removes a class from the students package and updates the current lesson number",
			Callback:    getStudentsCallback,
		},
		"help": {
			Name:        "help",
			Description: "Get a list of possible commands.",
			Callback:    helpCallback,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the program.",
			Callback:    exitCallback,
		},
	}
}
