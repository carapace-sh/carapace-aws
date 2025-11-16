package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_importVmImageCmd = &cobra.Command{
	Use:   "import-vm-image",
	Short: "When you export your virtual machine (VM) from its virtualization environment, that process creates a set of one or more disk container files that act as snapshots of your VM’s environment, settings, and data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_importVmImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_importVmImageCmd).Standalone()

		imagebuilder_importVmImageCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
		imagebuilder_importVmImageCmd.Flags().String("description", "", "The description for the base image that is created by the import process.")
		imagebuilder_importVmImageCmd.Flags().String("logging-configuration", "", "Define logging configuration for the image build process.")
		imagebuilder_importVmImageCmd.Flags().String("name", "", "The name of the base image that is created by the import process.")
		imagebuilder_importVmImageCmd.Flags().String("os-version", "", "The operating system version for the imported VM.")
		imagebuilder_importVmImageCmd.Flags().String("platform", "", "The operating system platform for the imported VM.")
		imagebuilder_importVmImageCmd.Flags().String("semantic-version", "", "The semantic version to attach to the base image that was created during the import process.")
		imagebuilder_importVmImageCmd.Flags().String("tags", "", "Tags that are attached to the import resources.")
		imagebuilder_importVmImageCmd.Flags().String("vm-import-task-id", "", "The `importTaskId` (API) or `ImportTaskId` (CLI) from the Amazon EC2 VM import process.")
		imagebuilder_importVmImageCmd.MarkFlagRequired("client-token")
		imagebuilder_importVmImageCmd.MarkFlagRequired("name")
		imagebuilder_importVmImageCmd.MarkFlagRequired("platform")
		imagebuilder_importVmImageCmd.MarkFlagRequired("semantic-version")
		imagebuilder_importVmImageCmd.MarkFlagRequired("vm-import-task-id")
	})
	imagebuilderCmd.AddCommand(imagebuilder_importVmImageCmd)
}
