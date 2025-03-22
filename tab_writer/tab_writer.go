package tabwriter

import (
	"fmt"
	"io"
	"text/tabwriter"
)

const minWidth int = 0
const tabWidth int = 0
const padding int = 1
const padChar byte = ' '
const flags uint = 0

func PrettyWrite(writable string, output io.Writer) {
	w := tabwriter.NewWriter(
		output,
		minWidth,
		tabWidth,
		padding,
		padChar,
		flags,
	)

	fmt.Fprintln(w, writable)
	w.Flush()
}
