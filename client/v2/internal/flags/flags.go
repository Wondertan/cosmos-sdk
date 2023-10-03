package flags

// This defines flags names that can be used in autocli.
const (
	// FlagFrom is the flag to set the from address from which to sign the transaction.
	FlagFrom = "from"

	// FlagOutput is the flag to set the output format.
	FlagOutput = "output"

	// FlagNoIndent is the flag to not indent the output.
	FlagNoIndent = "no-indent"
)

// List of supported output formats
const (
	OutputFormatJSON = "json"
	OutputFormatText = "text"
)
