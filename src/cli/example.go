package cli

import (
	"netools/src/example"
)

func init() {
	rootCmd.AddCommand(example.Cli)

	example.Cli.PersistentFlags().StringP("message", "m", "", "example of cli")

	_ = example.Cli.MarkPersistentFlagRequired("message")

}
