package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getImagePipelineCmd = &cobra.Command{
	Use:   "get-image-pipeline",
	Short: "Gets an image pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getImagePipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_getImagePipelineCmd).Standalone()

		imagebuilder_getImagePipelineCmd.Flags().String("image-pipeline-arn", "", "The Amazon Resource Name (ARN) of the image pipeline that you want to retrieve.")
		imagebuilder_getImagePipelineCmd.MarkFlagRequired("image-pipeline-arn")
	})
	imagebuilderCmd.AddCommand(imagebuilder_getImagePipelineCmd)
}
