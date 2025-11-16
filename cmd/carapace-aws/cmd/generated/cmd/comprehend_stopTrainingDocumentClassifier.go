package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_stopTrainingDocumentClassifierCmd = &cobra.Command{
	Use:   "stop-training-document-classifier",
	Short: "Stops a document classifier training job while in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_stopTrainingDocumentClassifierCmd).Standalone()

	comprehend_stopTrainingDocumentClassifierCmd.Flags().String("document-classifier-arn", "", "The Amazon Resource Name (ARN) that identifies the document classifier currently being trained.")
	comprehend_stopTrainingDocumentClassifierCmd.MarkFlagRequired("document-classifier-arn")
	comprehendCmd.AddCommand(comprehend_stopTrainingDocumentClassifierCmd)
}
