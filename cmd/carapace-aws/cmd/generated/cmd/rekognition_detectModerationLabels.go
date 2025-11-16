package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_detectModerationLabelsCmd = &cobra.Command{
	Use:   "detect-moderation-labels",
	Short: "Detects unsafe content in a specified JPEG or PNG format image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_detectModerationLabelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_detectModerationLabelsCmd).Standalone()

		rekognition_detectModerationLabelsCmd.Flags().String("human-loop-config", "", "Sets up the configuration for human evaluation, including the FlowDefinition the image will be sent to.")
		rekognition_detectModerationLabelsCmd.Flags().String("image", "", "The input image as base64-encoded bytes or an S3 object.")
		rekognition_detectModerationLabelsCmd.Flags().String("min-confidence", "", "Specifies the minimum confidence level for the labels to return.")
		rekognition_detectModerationLabelsCmd.Flags().String("project-version", "", "Identifier for the custom adapter.")
		rekognition_detectModerationLabelsCmd.MarkFlagRequired("image")
	})
	rekognitionCmd.AddCommand(rekognition_detectModerationLabelsCmd)
}
