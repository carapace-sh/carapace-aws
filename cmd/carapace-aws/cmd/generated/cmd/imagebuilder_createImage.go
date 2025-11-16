package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_createImageCmd = &cobra.Command{
	Use:   "create-image",
	Short: "Creates a new image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_createImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_createImageCmd).Standalone()

		imagebuilder_createImageCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
		imagebuilder_createImageCmd.Flags().String("container-recipe-arn", "", "The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.")
		imagebuilder_createImageCmd.Flags().String("distribution-configuration-arn", "", "The Amazon Resource Name (ARN) of the distribution configuration that defines and configures the outputs of your pipeline.")
		imagebuilder_createImageCmd.Flags().String("enhanced-image-metadata-enabled", "", "Collects additional information about the image being created, including the operating system (OS) version and package list.")
		imagebuilder_createImageCmd.Flags().String("execution-role", "", "The name or Amazon Resource Name (ARN) for the IAM role you create that grants Image Builder access to perform workflow actions.")
		imagebuilder_createImageCmd.Flags().String("image-recipe-arn", "", "The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.")
		imagebuilder_createImageCmd.Flags().String("image-scanning-configuration", "", "Contains settings for vulnerability scans.")
		imagebuilder_createImageCmd.Flags().String("image-tests-configuration", "", "The image tests configuration of the image.")
		imagebuilder_createImageCmd.Flags().String("infrastructure-configuration-arn", "", "The Amazon Resource Name (ARN) of the infrastructure configuration that defines the environment in which your image will be built and tested.")
		imagebuilder_createImageCmd.Flags().String("logging-configuration", "", "Define logging configuration for the image build process.")
		imagebuilder_createImageCmd.Flags().String("tags", "", "The tags of the image.")
		imagebuilder_createImageCmd.Flags().String("workflows", "", "Contains an array of workflow configuration objects.")
		imagebuilder_createImageCmd.MarkFlagRequired("client-token")
		imagebuilder_createImageCmd.MarkFlagRequired("infrastructure-configuration-arn")
	})
	imagebuilderCmd.AddCommand(imagebuilder_createImageCmd)
}
