package cli_options

func GetOptions() map[string]cli_option {
	return map[string]cli_option{
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
