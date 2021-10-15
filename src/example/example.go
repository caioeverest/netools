package example

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cli = &cobra.Command{
	Use:   "example",
	Short: "A brief description of your command",
	Long:  `A longer description.`,
	Run: func(cmd *cobra.Command, args []string) {
		msg, _ := cmd.Flags().GetString("message")

		fmt.Println(msg)
	},
}
