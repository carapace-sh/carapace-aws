package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_deleteDocumentClassifierCmd = &cobra.Command{
	Use:   "delete-document-classifier",
	Short: "Deletes a previously created document classifier\n\nOnly those classifiers that are in terminated states (IN\\_ERROR, TRAINED) will be deleted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_deleteDocumentClassifierCmd).Standalone()

	comprehend_deleteDocumentClassifierCmd.Flags().String("document-classifier-arn", "", "The Amazon Resource Name (ARN) that identifies the document classifier.")
	comprehend_deleteDocumentClassifierCmd.MarkFlagRequired("document-classifier-arn")
	comprehendCmd.AddCommand(comprehend_deleteDocumentClassifierCmd)
}
