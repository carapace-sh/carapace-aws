package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_getFaceDetectionCmd = &cobra.Command{
	Use:   "get-face-detection",
	Short: "Gets face detection results for a Amazon Rekognition Video analysis started by [StartFaceDetection]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_getFaceDetectionCmd).Standalone()

	rekognition_getFaceDetectionCmd.Flags().String("job-id", "", "Unique identifier for the face detection job.")
	rekognition_getFaceDetectionCmd.Flags().String("max-results", "", "Maximum number of results to return per paginated call.")
	rekognition_getFaceDetectionCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there are more faces to retrieve), Amazon Rekognition Video returns a pagination token in the response.")
	rekognition_getFaceDetectionCmd.MarkFlagRequired("job-id")
	rekognitionCmd.AddCommand(rekognition_getFaceDetectionCmd)
}
