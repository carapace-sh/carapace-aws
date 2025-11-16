package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_listDatasetEntriesCmd = &cobra.Command{
	Use:   "list-dataset-entries",
	Short: "This operation applies only to Amazon Rekognition Custom Labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_listDatasetEntriesCmd).Standalone()

	rekognition_listDatasetEntriesCmd.Flags().String("contains-labels", "", "Specifies a label filter for the response.")
	rekognition_listDatasetEntriesCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) for the dataset that you want to use.")
	rekognition_listDatasetEntriesCmd.Flags().String("has-errors", "", "Specifies an error filter for the response.")
	rekognition_listDatasetEntriesCmd.Flags().String("labeled", "", "Specify `true` to get only the JSON Lines where the image is labeled.")
	rekognition_listDatasetEntriesCmd.Flags().String("max-results", "", "The maximum number of results to return per paginated call.")
	rekognition_listDatasetEntriesCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more results to retrieve), Amazon Rekognition Custom Labels returns a pagination token in the response.")
	rekognition_listDatasetEntriesCmd.Flags().String("source-ref-contains", "", "If specified, `ListDatasetEntries` only returns JSON Lines where the value of `SourceRefContains` is part of the `source-ref` field.")
	rekognition_listDatasetEntriesCmd.MarkFlagRequired("dataset-arn")
	rekognitionCmd.AddCommand(rekognition_listDatasetEntriesCmd)
}
