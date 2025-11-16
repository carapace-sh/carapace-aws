package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_distributeDatasetEntriesCmd = &cobra.Command{
	Use:   "distribute-dataset-entries",
	Short: "This operation applies only to Amazon Rekognition Custom Labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_distributeDatasetEntriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_distributeDatasetEntriesCmd).Standalone()

		rekognition_distributeDatasetEntriesCmd.Flags().String("datasets", "", "The ARNS for the training dataset and test dataset that you want to use.")
		rekognition_distributeDatasetEntriesCmd.MarkFlagRequired("datasets")
	})
	rekognitionCmd.AddCommand(rekognition_distributeDatasetEntriesCmd)
}
