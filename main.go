package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use: "hello-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World")
	},
}

func main() {
	rootCmd.Version = fmt.Sprintf("%s, commit %s, built at %s", version, commit, date)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
