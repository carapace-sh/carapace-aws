package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_createDatasetCmd = &cobra.Command{
	Use:   "create-dataset",
	Short: "This operation applies only to Amazon Rekognition Custom Labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_createDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_createDatasetCmd).Standalone()

		rekognition_createDatasetCmd.Flags().String("dataset-source", "", "The source files for the dataset.")
		rekognition_createDatasetCmd.Flags().String("dataset-type", "", "The type of the dataset.")
		rekognition_createDatasetCmd.Flags().String("project-arn", "", "The ARN of the Amazon Rekognition Custom Labels project to which you want to asssign the dataset.")
		rekognition_createDatasetCmd.Flags().String("tags", "", "A set of tags (key-value pairs) that you want to attach to the dataset.")
		rekognition_createDatasetCmd.MarkFlagRequired("dataset-type")
		rekognition_createDatasetCmd.MarkFlagRequired("project-arn")
	})
	rekognitionCmd.AddCommand(rekognition_createDatasetCmd)
}
