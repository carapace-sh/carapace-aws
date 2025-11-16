package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_analyzeDocumentCmd = &cobra.Command{
	Use:   "analyze-document",
	Short: "Analyzes an input document for relationships between detected items.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_analyzeDocumentCmd).Standalone()

	textract_analyzeDocumentCmd.Flags().String("adapters-config", "", "Specifies the adapter to be used when analyzing a document.")
	textract_analyzeDocumentCmd.Flags().String("document", "", "The input document as base64-encoded bytes or an Amazon S3 object.")
	textract_analyzeDocumentCmd.Flags().String("feature-types", "", "A list of the types of analysis to perform.")
	textract_analyzeDocumentCmd.Flags().String("human-loop-config", "", "Sets the configuration for the human in the loop workflow for analyzing documents.")
	textract_analyzeDocumentCmd.Flags().String("queries-config", "", "Contains Queries and the alias for those Queries, as determined by the input.")
	textract_analyzeDocumentCmd.MarkFlagRequired("document")
	textract_analyzeDocumentCmd.MarkFlagRequired("feature-types")
	textractCmd.AddCommand(textract_analyzeDocumentCmd)
}
