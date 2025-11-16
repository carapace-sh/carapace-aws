package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_describeDatasetCmd = &cobra.Command{
	Use:   "describe-dataset",
	Short: "This operation applies only to Amazon Rekognition Custom Labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_describeDatasetCmd).Standalone()

	rekognition_describeDatasetCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the dataset that you want to describe.")
	rekognition_describeDatasetCmd.MarkFlagRequired("dataset-arn")
	rekognitionCmd.AddCommand(rekognition_describeDatasetCmd)
}
