package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_getTextDetectionCmd = &cobra.Command{
	Use:   "get-text-detection",
	Short: "Gets the text detection results of a Amazon Rekognition Video analysis started by [StartTextDetection]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_getTextDetectionCmd).Standalone()

	rekognition_getTextDetectionCmd.Flags().String("job-id", "", "Job identifier for the text detection operation for which you want results returned.")
	rekognition_getTextDetectionCmd.Flags().String("max-results", "", "Maximum number of results to return per paginated call.")
	rekognition_getTextDetectionCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there are more labels to retrieve), Amazon Rekognition Video returns a pagination token in the response.")
	rekognition_getTextDetectionCmd.MarkFlagRequired("job-id")
	rekognitionCmd.AddCommand(rekognition_getTextDetectionCmd)
}
