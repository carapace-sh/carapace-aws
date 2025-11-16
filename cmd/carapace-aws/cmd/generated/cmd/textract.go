package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textractCmd = &cobra.Command{
	Use:   "textract",
	Short: "Amazon Textract detects and analyzes text in documents and converts it into machine-readable text.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textractCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(textractCmd).Standalone()

	})
	rootCmd.AddCommand(textractCmd)
}
