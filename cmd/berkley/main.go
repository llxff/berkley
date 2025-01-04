package main

import (
	"log"

	"berkley/cmd/berkley/commands"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "berkley",
		Short: "Berkeley DB tools",
	}

	rootCmd.AddCommand(
		commands.GetGenerateCmd(),
		commands.GetMetaCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
