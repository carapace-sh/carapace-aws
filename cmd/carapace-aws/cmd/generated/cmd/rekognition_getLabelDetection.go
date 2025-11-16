package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_getLabelDetectionCmd = &cobra.Command{
	Use:   "get-label-detection",
	Short: "Gets the label detection results of a Amazon Rekognition Video analysis started by [StartLabelDetection]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_getLabelDetectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_getLabelDetectionCmd).Standalone()

		rekognition_getLabelDetectionCmd.Flags().String("aggregate-by", "", "Defines how to aggregate the returned results.")
		rekognition_getLabelDetectionCmd.Flags().String("job-id", "", "Job identifier for the label detection operation for which you want results returned.")
		rekognition_getLabelDetectionCmd.Flags().String("max-results", "", "Maximum number of results to return per paginated call.")
		rekognition_getLabelDetectionCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there are more labels to retrieve), Amazon Rekognition Video returns a pagination token in the response.")
		rekognition_getLabelDetectionCmd.Flags().String("sort-by", "", "Sort to use for elements in the `Labels` array.")
		rekognition_getLabelDetectionCmd.MarkFlagRequired("job-id")
	})
	rekognitionCmd.AddCommand(rekognition_getLabelDetectionCmd)
}
