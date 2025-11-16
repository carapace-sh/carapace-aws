package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_detectLabelsCmd = &cobra.Command{
	Use:   "detect-labels",
	Short: "Detects instances of real-world entities within an image (JPEG or PNG) provided as input.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_detectLabelsCmd).Standalone()

	rekognition_detectLabelsCmd.Flags().String("features", "", "A list of the types of analysis to perform.")
	rekognition_detectLabelsCmd.Flags().String("image", "", "The input image as base64-encoded bytes or an S3 object.")
	rekognition_detectLabelsCmd.Flags().String("max-labels", "", "Maximum number of labels you want the service to return in the response.")
	rekognition_detectLabelsCmd.Flags().String("min-confidence", "", "Specifies the minimum confidence level for the labels to return.")
	rekognition_detectLabelsCmd.Flags().String("settings", "", "A list of the filters to be applied to returned detected labels and image properties.")
	rekognition_detectLabelsCmd.MarkFlagRequired("image")
	rekognitionCmd.AddCommand(rekognition_detectLabelsCmd)
}
