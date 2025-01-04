package commands

import (
	"berkley/internal/generator"
	"fmt"

	"github.com/spf13/cobra"
)

func GetGenerateCmd() *cobra.Command {
	var numPairs int

	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate random key-value pairs with realistic data",
		Run: func(cmd *cobra.Command, args []string) {
			generateSequence(numPairs)
		},
	}

	generateCmd.Flags().IntVarP(&numPairs, "num", "n", 10, "number of key-value pairs to generate")

	return generateCmd
}

func generateSequence(numPairs int) {
	for i := 1; i <= numPairs; i++ {
		pair := generator.GeneratePair(i)
		fmt.Println(pair.Key)
		fmt.Println(pair.Value)
	}
}
