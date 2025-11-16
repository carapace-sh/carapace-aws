package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteExternalModelCmd = &cobra.Command{
	Use:   "delete-external-model",
	Short: "Removes a SageMaker model from Amazon Fraud Detector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteExternalModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_deleteExternalModelCmd).Standalone()

		frauddetector_deleteExternalModelCmd.Flags().String("model-endpoint", "", "The endpoint of the Amazon Sagemaker model to delete.")
		frauddetector_deleteExternalModelCmd.MarkFlagRequired("model-endpoint")
	})
	frauddetectorCmd.AddCommand(frauddetector_deleteExternalModelCmd)
}
