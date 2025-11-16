package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_getSegmentDetectionCmd = &cobra.Command{
	Use:   "get-segment-detection",
	Short: "Gets the segment detection results of a Amazon Rekognition Video analysis started by [StartSegmentDetection]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_getSegmentDetectionCmd).Standalone()

	rekognition_getSegmentDetectionCmd.Flags().String("job-id", "", "Job identifier for the text detection operation for which you want results returned.")
	rekognition_getSegmentDetectionCmd.Flags().String("max-results", "", "Maximum number of results to return per paginated call.")
	rekognition_getSegmentDetectionCmd.Flags().String("next-token", "", "If the response is truncated, Amazon Rekognition Video returns this token that you can use in the subsequent request to retrieve the next set of text.")
	rekognition_getSegmentDetectionCmd.MarkFlagRequired("job-id")
	rekognitionCmd.AddCommand(rekognition_getSegmentDetectionCmd)
}
