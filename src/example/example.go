package example

import (
	"github.com/spf13/cobra"
)

// CliExample to represent command - ./netools example
var CliExample = &cobra.Command{
	Use:   "example",
	Short: "An example CLI",
	Long:  `An example CLI`,
}
