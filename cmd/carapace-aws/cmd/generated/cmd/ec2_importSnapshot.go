package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_importSnapshotCmd = &cobra.Command{
	Use:   "import-snapshot",
	Short: "Imports a disk into an EBS snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_importSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_importSnapshotCmd).Standalone()

		ec2_importSnapshotCmd.Flags().String("client-data", "", "The client-specific data.")
		ec2_importSnapshotCmd.Flags().String("client-token", "", "Token to enable idempotency for VM import requests.")
		ec2_importSnapshotCmd.Flags().String("description", "", "The description string for the import snapshot task.")
		ec2_importSnapshotCmd.Flags().String("disk-container", "", "Information about the disk container.")
		ec2_importSnapshotCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_importSnapshotCmd.Flags().Bool("encrypted", false, "Specifies whether the destination snapshot of the imported image should be encrypted.")
		ec2_importSnapshotCmd.Flags().String("kms-key-id", "", "An identifier for the symmetric KMS key to use when creating the encrypted snapshot.")
		ec2_importSnapshotCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_importSnapshotCmd.Flags().Bool("no-encrypted", false, "Specifies whether the destination snapshot of the imported image should be encrypted.")
		ec2_importSnapshotCmd.Flags().String("role-name", "", "The name of the role to use when not using the default role, 'vmimport'.")
		ec2_importSnapshotCmd.Flags().String("tag-specifications", "", "The tags to apply to the import snapshot task during creation.")
		ec2_importSnapshotCmd.Flag("no-dry-run").Hidden = true
		ec2_importSnapshotCmd.Flag("no-encrypted").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_importSnapshotCmd)
}
