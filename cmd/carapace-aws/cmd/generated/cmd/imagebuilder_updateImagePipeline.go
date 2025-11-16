package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_updateImagePipelineCmd = &cobra.Command{
	Use:   "update-image-pipeline",
	Short: "Updates an image pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_updateImagePipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_updateImagePipelineCmd).Standalone()

		imagebuilder_updateImagePipelineCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
		imagebuilder_updateImagePipelineCmd.Flags().String("container-recipe-arn", "", "The Amazon Resource Name (ARN) of the container pipeline to update.")
		imagebuilder_updateImagePipelineCmd.Flags().String("description", "", "The description of the image pipeline.")
		imagebuilder_updateImagePipelineCmd.Flags().String("distribution-configuration-arn", "", "The Amazon Resource Name (ARN) of the distribution configuration that Image Builder uses to configure and distribute images that this image pipeline has updated.")
		imagebuilder_updateImagePipelineCmd.Flags().String("enhanced-image-metadata-enabled", "", "Collects additional information about the image being created, including the operating system (OS) version and package list.")
		imagebuilder_updateImagePipelineCmd.Flags().String("execution-role", "", "The name or Amazon Resource Name (ARN) for the IAM role you create that grants Image Builder access to perform workflow actions.")
		imagebuilder_updateImagePipelineCmd.Flags().String("image-pipeline-arn", "", "The Amazon Resource Name (ARN) of the image pipeline that you want to update.")
		imagebuilder_updateImagePipelineCmd.Flags().String("image-recipe-arn", "", "The Amazon Resource Name (ARN) of the image recipe that will be used to configure images updated by this image pipeline.")
		imagebuilder_updateImagePipelineCmd.Flags().String("image-scanning-configuration", "", "Contains settings for vulnerability scans.")
		imagebuilder_updateImagePipelineCmd.Flags().String("image-tests-configuration", "", "The image test configuration of the image pipeline.")
		imagebuilder_updateImagePipelineCmd.Flags().String("infrastructure-configuration-arn", "", "The Amazon Resource Name (ARN) of the infrastructure configuration that Image Builder uses to build images that this image pipeline has updated.")
		imagebuilder_updateImagePipelineCmd.Flags().String("logging-configuration", "", "Update logging configuration for the output image that's created when the pipeline runs.")
		imagebuilder_updateImagePipelineCmd.Flags().String("schedule", "", "The schedule of the image pipeline.")
		imagebuilder_updateImagePipelineCmd.Flags().String("status", "", "The status of the image pipeline.")
		imagebuilder_updateImagePipelineCmd.Flags().String("workflows", "", "Contains the workflows to run for the pipeline.")
		imagebuilder_updateImagePipelineCmd.MarkFlagRequired("client-token")
		imagebuilder_updateImagePipelineCmd.MarkFlagRequired("image-pipeline-arn")
		imagebuilder_updateImagePipelineCmd.MarkFlagRequired("infrastructure-configuration-arn")
	})
	imagebuilderCmd.AddCommand(imagebuilder_updateImagePipelineCmd)
}
