package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_importImageCmd = &cobra.Command{
	Use:   "import-image",
	Short: "To import your virtual machines (VMs) with a console-based experience, you can use the *Import virtual machine images to Amazon Web Services* template in the [Migration Hub Orchestrator console](https://console.aws.amazon.com/migrationhub/orchestrator). For more information, see the [*Migration Hub Orchestrator User Guide*](https://docs.aws.amazon.com/migrationhub-orchestrator/latest/userguide/import-vm-images.html) .",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_importImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_importImageCmd).Standalone()

		ec2_importImageCmd.Flags().String("architecture", "", "The architecture of the virtual machine.")
		ec2_importImageCmd.Flags().String("boot-mode", "", "The boot mode of the virtual machine.")
		ec2_importImageCmd.Flags().String("client-data", "", "The client-specific data.")
		ec2_importImageCmd.Flags().String("client-token", "", "The token to enable idempotency for VM import requests.")
		ec2_importImageCmd.Flags().String("description", "", "A description string for the import image task.")
		ec2_importImageCmd.Flags().String("disk-containers", "", "Information about the disk containers.")
		ec2_importImageCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_importImageCmd.Flags().Bool("encrypted", false, "Specifies whether the destination AMI of the imported image should be encrypted.")
		ec2_importImageCmd.Flags().String("hypervisor", "", "The target hypervisor platform.")
		ec2_importImageCmd.Flags().String("kms-key-id", "", "An identifier for the symmetric KMS key to use when creating the encrypted AMI.")
		ec2_importImageCmd.Flags().String("license-specifications", "", "The ARNs of the license configurations.")
		ec2_importImageCmd.Flags().String("license-type", "", "The license type to be used for the Amazon Machine Image (AMI) after importing.")
		ec2_importImageCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_importImageCmd.Flags().Bool("no-encrypted", false, "Specifies whether the destination AMI of the imported image should be encrypted.")
		ec2_importImageCmd.Flags().String("platform", "", "The operating system of the virtual machine.")
		ec2_importImageCmd.Flags().String("role-name", "", "The name of the role to use when not using the default role, 'vmimport'.")
		ec2_importImageCmd.Flags().String("tag-specifications", "", "The tags to apply to the import image task during creation.")
		ec2_importImageCmd.Flags().String("usage-operation", "", "The usage operation value.")
		ec2_importImageCmd.Flag("no-dry-run").Hidden = true
		ec2_importImageCmd.Flag("no-encrypted").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_importImageCmd)
}
