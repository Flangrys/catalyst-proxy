package cli

// Represents the arguments passed into the command line executable.
type Flags struct {
	Verbose    *bool
	ConfigPath *string
}
