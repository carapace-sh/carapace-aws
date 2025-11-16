package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_detectDocumentTextCmd = &cobra.Command{
	Use:   "detect-document-text",
	Short: "Detects text in the input document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_detectDocumentTextCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(textract_detectDocumentTextCmd).Standalone()

		textract_detectDocumentTextCmd.Flags().String("document", "", "The input document as base64-encoded bytes or an Amazon S3 object.")
		textract_detectDocumentTextCmd.MarkFlagRequired("document")
	})
	textractCmd.AddCommand(textract_detectDocumentTextCmd)
}
