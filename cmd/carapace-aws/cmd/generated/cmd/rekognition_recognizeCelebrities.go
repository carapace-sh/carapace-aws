package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_recognizeCelebritiesCmd = &cobra.Command{
	Use:   "recognize-celebrities",
	Short: "Returns an array of celebrities recognized in the input image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_recognizeCelebritiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_recognizeCelebritiesCmd).Standalone()

		rekognition_recognizeCelebritiesCmd.Flags().String("image", "", "The input image as base64-encoded bytes or an S3 object.")
		rekognition_recognizeCelebritiesCmd.MarkFlagRequired("image")
	})
	rekognitionCmd.AddCommand(rekognition_recognizeCelebritiesCmd)
}
