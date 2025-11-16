package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "Provides translation of the input content from the source language to the target language.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translateCmd).Standalone()

	rootCmd.AddCommand(translateCmd)
}
