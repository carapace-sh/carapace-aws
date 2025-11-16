package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_startImagePipelineExecutionCmd = &cobra.Command{
	Use:   "start-image-pipeline-execution",
	Short: "Manually triggers a pipeline to create an image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_startImagePipelineExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_startImagePipelineExecutionCmd).Standalone()

		imagebuilder_startImagePipelineExecutionCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
		imagebuilder_startImagePipelineExecutionCmd.Flags().String("image-pipeline-arn", "", "The Amazon Resource Name (ARN) of the image pipeline that you want to manually invoke.")
		imagebuilder_startImagePipelineExecutionCmd.Flags().String("tags", "", "Specify tags for Image Builder to apply to the image resource that's created When it starts pipeline execution.")
		imagebuilder_startImagePipelineExecutionCmd.MarkFlagRequired("client-token")
		imagebuilder_startImagePipelineExecutionCmd.MarkFlagRequired("image-pipeline-arn")
	})
	imagebuilderCmd.AddCommand(imagebuilder_startImagePipelineExecutionCmd)
}
