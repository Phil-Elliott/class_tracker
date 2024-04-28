package cli_options

import (
	"os"
)

func exitCallback() error {
	os.Exit(0)
	return nil
}
