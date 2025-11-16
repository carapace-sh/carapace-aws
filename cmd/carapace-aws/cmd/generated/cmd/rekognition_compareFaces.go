package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_compareFacesCmd = &cobra.Command{
	Use:   "compare-faces",
	Short: "Compares a face in the *source* input image with each of the 100 largest faces detected in the *target* input image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_compareFacesCmd).Standalone()

	rekognition_compareFacesCmd.Flags().String("quality-filter", "", "A filter that specifies a quality bar for how much filtering is done to identify faces.")
	rekognition_compareFacesCmd.Flags().String("similarity-threshold", "", "The minimum level of confidence in the face matches that a match must meet to be included in the `FaceMatches` array.")
	rekognition_compareFacesCmd.Flags().String("source-image", "", "The input image as base64-encoded bytes or an S3 object.")
	rekognition_compareFacesCmd.Flags().String("target-image", "", "The target image as base64-encoded bytes or an S3 object.")
	rekognition_compareFacesCmd.MarkFlagRequired("source-image")
	rekognition_compareFacesCmd.MarkFlagRequired("target-image")
	rekognitionCmd.AddCommand(rekognition_compareFacesCmd)
}
