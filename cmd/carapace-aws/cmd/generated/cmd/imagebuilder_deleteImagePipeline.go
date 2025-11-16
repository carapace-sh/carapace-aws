package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_deleteImagePipelineCmd = &cobra.Command{
	Use:   "delete-image-pipeline",
	Short: "Deletes an image pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_deleteImagePipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_deleteImagePipelineCmd).Standalone()

		imagebuilder_deleteImagePipelineCmd.Flags().String("image-pipeline-arn", "", "The Amazon Resource Name (ARN) of the image pipeline to delete.")
		imagebuilder_deleteImagePipelineCmd.MarkFlagRequired("image-pipeline-arn")
	})
	imagebuilderCmd.AddCommand(imagebuilder_deleteImagePipelineCmd)
}
