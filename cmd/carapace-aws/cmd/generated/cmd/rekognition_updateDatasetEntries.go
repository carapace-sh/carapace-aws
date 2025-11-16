package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_updateDatasetEntriesCmd = &cobra.Command{
	Use:   "update-dataset-entries",
	Short: "This operation applies only to Amazon Rekognition Custom Labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_updateDatasetEntriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_updateDatasetEntriesCmd).Standalone()

		rekognition_updateDatasetEntriesCmd.Flags().String("changes", "", "The changes that you want to make to the dataset.")
		rekognition_updateDatasetEntriesCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the dataset that you want to update.")
		rekognition_updateDatasetEntriesCmd.MarkFlagRequired("changes")
		rekognition_updateDatasetEntriesCmd.MarkFlagRequired("dataset-arn")
	})
	rekognitionCmd.AddCommand(rekognition_updateDatasetEntriesCmd)
}
