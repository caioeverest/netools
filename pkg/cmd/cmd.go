package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"netools/pkg/cmd/route"
)

// cmd represents the base command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "netools",
	Short: "Netools CLI App",
	Long:  `Welcome to netools CLI`,
}

func init() {
	route.Initialize(cmd)
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
