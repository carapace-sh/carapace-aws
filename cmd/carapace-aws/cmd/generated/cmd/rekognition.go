package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognitionCmd = &cobra.Command{
	Use:   "rekognition",
	Short: "This is the API Reference for [Amazon Rekognition Image](https://docs.aws.amazon.com/rekognition/latest/dg/images.html), [Amazon Rekognition Custom Labels](https://docs.aws.amazon.com/rekognition/latest/customlabels-dg/what-is.html), [Amazon Rekognition Stored Video](https://docs.aws.amazon.com/rekognition/latest/dg/video.html), [Amazon Rekognition Streaming Video](https://docs.aws.amazon.com/rekognition/latest/dg/streaming-video.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognitionCmd).Standalone()

	})
	rootCmd.AddCommand(rekognitionCmd)
}
