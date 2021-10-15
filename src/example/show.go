package example

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	cmdShow = "show"
)

// SubCliExampleWithArgs to represent command - ./netools example show --message="<input text>"
var SubCliExampleWithArgs = &cobra.Command{
	Use:   cmdShow,
	Short: "An example of sub CLI of a CLI with args",
	Long:  `An example of sub CLI of a CLI with args`,
	Run: func(cmd *cobra.Command, args []string) {
		msg, _ := cmd.Flags().GetString("message")

		fmt.Println(msg)
	},
}
