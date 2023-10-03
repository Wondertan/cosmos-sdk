package tx

// This defines flags names that can be used in autocli.
const (
	// FlagFrom is the flag to set the from address from which to sign the transaction.
	FlagFrom = "from"

	// FlagAccountNumber is the flag to set the account number of the transaction.
	FlagAccountNumber = "account-number"

	// FlagSequence is the flag to set the sequence number of the transaction.
	FlagSequence = "sequence"

	// FlagGas defines the gas units to set per-transaction.
	FlagGas = "gas"

	// FlagGasAdjustment precises how to scale the maximum amount of gas.
	FlagGasAdjustment = "gas-adjustment"

	// FlagGasPrices defines the gas prices per gas unit.
	FlagGasPrices = "gas-prices"

	// FlagFees defines the fees to pay along with the transaction.
	FlagFees = "fees"

	// FlagNote is the flag to set a note (memo) in the transaction.
	FlagNote = "note"

	// FlagTimeoutHeight is the flag to set the timeout height of the transaction.
	FlagTimeoutHeight = "timeout-height"

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

// List of supported sign modes
const (
	// SignModeDirect is the value of the --sign-mode flag for SIGN_MODE_DIRECT
	SignModeDirect = "direct"
	// SignModeTextual is the value of the --sign-mode flag for SIGN_MODE_TEXTUAL.
	SignModeTextual = "textual"
)
