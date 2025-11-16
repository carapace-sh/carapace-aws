package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_detectTextCmd = &cobra.Command{
	Use:   "detect-text",
	Short: "Detects text in the input image and converts it into machine-readable text.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_detectTextCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_detectTextCmd).Standalone()

		rekognition_detectTextCmd.Flags().String("filters", "", "Optional parameters that let you set the criteria that the text must meet to be included in your response.")
		rekognition_detectTextCmd.Flags().String("image", "", "The input image as base64-encoded bytes or an Amazon S3 object.")
		rekognition_detectTextCmd.MarkFlagRequired("image")
	})
	rekognitionCmd.AddCommand(rekognition_detectTextCmd)
}
