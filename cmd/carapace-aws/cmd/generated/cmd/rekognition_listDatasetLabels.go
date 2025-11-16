package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_listDatasetLabelsCmd = &cobra.Command{
	Use:   "list-dataset-labels",
	Short: "This operation applies only to Amazon Rekognition Custom Labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_listDatasetLabelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_listDatasetLabelsCmd).Standalone()

		rekognition_listDatasetLabelsCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the dataset that you want to use.")
		rekognition_listDatasetLabelsCmd.Flags().String("max-results", "", "The maximum number of results to return per paginated call.")
		rekognition_listDatasetLabelsCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more results to retrieve), Amazon Rekognition Custom Labels returns a pagination token in the response.")
		rekognition_listDatasetLabelsCmd.MarkFlagRequired("dataset-arn")
	})
	rekognitionCmd.AddCommand(rekognition_listDatasetLabelsCmd)
}
