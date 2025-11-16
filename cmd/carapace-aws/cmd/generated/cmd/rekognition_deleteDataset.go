package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_deleteDatasetCmd = &cobra.Command{
	Use:   "delete-dataset",
	Short: "This operation applies only to Amazon Rekognition Custom Labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_deleteDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_deleteDatasetCmd).Standalone()

		rekognition_deleteDatasetCmd.Flags().String("dataset-arn", "", "The ARN of the Amazon Rekognition Custom Labels dataset that you want to delete.")
		rekognition_deleteDatasetCmd.MarkFlagRequired("dataset-arn")
	})
	rekognitionCmd.AddCommand(rekognition_deleteDatasetCmd)
}
