package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listImagePipelineImagesCmd = &cobra.Command{
	Use:   "list-image-pipeline-images",
	Short: "Returns a list of images created by the specified pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listImagePipelineImagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_listImagePipelineImagesCmd).Standalone()

		imagebuilder_listImagePipelineImagesCmd.Flags().String("filters", "", "Use the following filters to streamline results:")
		imagebuilder_listImagePipelineImagesCmd.Flags().String("image-pipeline-arn", "", "The Amazon Resource Name (ARN) of the image pipeline whose images you want to view.")
		imagebuilder_listImagePipelineImagesCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
		imagebuilder_listImagePipelineImagesCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
		imagebuilder_listImagePipelineImagesCmd.MarkFlagRequired("image-pipeline-arn")
	})
	imagebuilderCmd.AddCommand(imagebuilder_listImagePipelineImagesCmd)
}
