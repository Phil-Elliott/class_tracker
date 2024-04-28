package cli_options

import "fmt"

func helpCallback() error {
	fmt.Println("Please see the list of options below.")
	commands := GetOptions()
	for _, command := range commands {
		fmt.Printf("- %s: %s\n", command.Name, command.Description)
	}

	return nil
}
