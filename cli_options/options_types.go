package cli_options

type cli_option struct {
	Name        string
	Description string
	Callback    func() error
}
