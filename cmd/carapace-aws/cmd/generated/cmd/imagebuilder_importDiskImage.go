package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_importDiskImageCmd = &cobra.Command{
	Use:   "import-disk-image",
	Short: "Import a Windows operating system image from a verified Microsoft ISO disk file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_importDiskImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_importDiskImageCmd).Standalone()

		imagebuilder_importDiskImageCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
		imagebuilder_importDiskImageCmd.Flags().String("description", "", "The description for your disk image import.")
		imagebuilder_importDiskImageCmd.Flags().String("execution-role", "", "The name or Amazon Resource Name (ARN) for the IAM role you create that grants Image Builder access to perform workflow actions to import an image from a Microsoft ISO file.")
		imagebuilder_importDiskImageCmd.Flags().String("infrastructure-configuration-arn", "", "The Amazon Resource Name (ARN) of the infrastructure configuration resource that's used for launching the EC2 instance on which the ISO image is built.")
		imagebuilder_importDiskImageCmd.Flags().String("logging-configuration", "", "Define logging configuration for the image build process.")
		imagebuilder_importDiskImageCmd.Flags().String("name", "", "The name of the image resource that's created from the import.")
		imagebuilder_importDiskImageCmd.Flags().String("os-version", "", "The operating system version for the imported image.")
		imagebuilder_importDiskImageCmd.Flags().String("platform", "", "The operating system platform for the imported image.")
		imagebuilder_importDiskImageCmd.Flags().String("semantic-version", "", "The semantic version to attach to the image that's created during the import process.")
		imagebuilder_importDiskImageCmd.Flags().String("tags", "", "Tags that are attached to image resources created from the import.")
		imagebuilder_importDiskImageCmd.Flags().String("uri", "", "The `uri` of the ISO disk file that's stored in Amazon S3.")
		imagebuilder_importDiskImageCmd.MarkFlagRequired("client-token")
		imagebuilder_importDiskImageCmd.MarkFlagRequired("infrastructure-configuration-arn")
		imagebuilder_importDiskImageCmd.MarkFlagRequired("name")
		imagebuilder_importDiskImageCmd.MarkFlagRequired("os-version")
		imagebuilder_importDiskImageCmd.MarkFlagRequired("platform")
		imagebuilder_importDiskImageCmd.MarkFlagRequired("semantic-version")
		imagebuilder_importDiskImageCmd.MarkFlagRequired("uri")
	})
	imagebuilderCmd.AddCommand(imagebuilder_importDiskImageCmd)
}
