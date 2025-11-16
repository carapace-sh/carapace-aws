package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_analyzeIdCmd = &cobra.Command{
	Use:   "analyze-id",
	Short: "Analyzes identity documents for relevant information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_analyzeIdCmd).Standalone()

	textract_analyzeIdCmd.Flags().String("document-pages", "", "The document being passed to AnalyzeID.")
	textract_analyzeIdCmd.MarkFlagRequired("document-pages")
	textractCmd.AddCommand(textract_analyzeIdCmd)
}
