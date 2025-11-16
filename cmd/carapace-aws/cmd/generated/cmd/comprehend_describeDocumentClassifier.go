package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describeDocumentClassifierCmd = &cobra.Command{
	Use:   "describe-document-classifier",
	Short: "Gets the properties associated with a document classifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describeDocumentClassifierCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_describeDocumentClassifierCmd).Standalone()

		comprehend_describeDocumentClassifierCmd.Flags().String("document-classifier-arn", "", "The Amazon Resource Name (ARN) that identifies the document classifier.")
		comprehend_describeDocumentClassifierCmd.MarkFlagRequired("document-classifier-arn")
	})
	comprehendCmd.AddCommand(comprehend_describeDocumentClassifierCmd)
}
