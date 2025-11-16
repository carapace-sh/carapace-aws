package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexRuntimeCmd = &cobra.Command{
	Use:   "lex-runtime",
	Short: "Amazon Lex provides both build and runtime endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexRuntimeCmd).Standalone()

	rootCmd.AddCommand(lexRuntimeCmd)
}
