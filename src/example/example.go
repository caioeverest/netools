package example

import (
	"github.com/spf13/cobra"
)

const (
	cmdExample = "example"
)

// CliExample to represent command - ./netools example
var CliExample = &cobra.Command{
	Use:   cmdExample,
	Short: "An example CLI",
	Long:  `An example CLI`,
}
