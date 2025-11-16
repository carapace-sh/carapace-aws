package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_createImagePipelineCmd = &cobra.Command{
	Use:   "create-image-pipeline",
	Short: "Creates a new image pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_createImagePipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_createImagePipelineCmd).Standalone()

		imagebuilder_createImagePipelineCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
		imagebuilder_createImagePipelineCmd.Flags().String("container-recipe-arn", "", "The Amazon Resource Name (ARN) of the container recipe that is used to configure images created by this container pipeline.")
		imagebuilder_createImagePipelineCmd.Flags().String("description", "", "The description of the image pipeline.")
		imagebuilder_createImagePipelineCmd.Flags().String("distribution-configuration-arn", "", "The Amazon Resource Name (ARN) of the distribution configuration that will be used to configure and distribute images created by this image pipeline.")
		imagebuilder_createImagePipelineCmd.Flags().String("enhanced-image-metadata-enabled", "", "Collects additional information about the image being created, including the operating system (OS) version and package list.")
		imagebuilder_createImagePipelineCmd.Flags().String("execution-role", "", "The name or Amazon Resource Name (ARN) for the IAM role you create that grants Image Builder access to perform workflow actions.")
		imagebuilder_createImagePipelineCmd.Flags().String("image-recipe-arn", "", "The Amazon Resource Name (ARN) of the image recipe that will be used to configure images created by this image pipeline.")
		imagebuilder_createImagePipelineCmd.Flags().String("image-scanning-configuration", "", "Contains settings for vulnerability scans.")
		imagebuilder_createImagePipelineCmd.Flags().String("image-tests-configuration", "", "The image test configuration of the image pipeline.")
		imagebuilder_createImagePipelineCmd.Flags().String("infrastructure-configuration-arn", "", "The Amazon Resource Name (ARN) of the infrastructure configuration that will be used to build images created by this image pipeline.")
		imagebuilder_createImagePipelineCmd.Flags().String("logging-configuration", "", "Define logging configuration for the image build process.")
		imagebuilder_createImagePipelineCmd.Flags().String("name", "", "The name of the image pipeline.")
		imagebuilder_createImagePipelineCmd.Flags().String("schedule", "", "The schedule of the image pipeline.")
		imagebuilder_createImagePipelineCmd.Flags().String("status", "", "The status of the image pipeline.")
		imagebuilder_createImagePipelineCmd.Flags().String("tags", "", "The tags of the image pipeline.")
		imagebuilder_createImagePipelineCmd.Flags().String("workflows", "", "Contains an array of workflow configuration objects.")
		imagebuilder_createImagePipelineCmd.MarkFlagRequired("client-token")
		imagebuilder_createImagePipelineCmd.MarkFlagRequired("infrastructure-configuration-arn")
		imagebuilder_createImagePipelineCmd.MarkFlagRequired("name")
	})
	imagebuilderCmd.AddCommand(imagebuilder_createImagePipelineCmd)
}
