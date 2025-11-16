package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createImageCmd = &cobra.Command{
	Use:   "create-image",
	Short: "Creates an Amazon EBS-backed AMI from an Amazon EBS-backed instance that is either running or stopped.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createImageCmd).Standalone()

	ec2_createImageCmd.Flags().String("block-device-mappings", "", "The block device mappings.")
	ec2_createImageCmd.Flags().String("description", "", "A description for the new image.")
	ec2_createImageCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createImageCmd.Flags().String("instance-id", "", "The ID of the instance.")
	ec2_createImageCmd.Flags().String("name", "", "A name for the new image.")
	ec2_createImageCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createImageCmd.Flags().Bool("no-no-reboot", false, "Indicates whether or not the instance should be automatically rebooted before creating the image.")
	ec2_createImageCmd.Flags().Bool("no-reboot", false, "Indicates whether or not the instance should be automatically rebooted before creating the image.")
	ec2_createImageCmd.Flags().String("snapshot-location", "", "Only supported for instances in Local Zones.")
	ec2_createImageCmd.Flags().String("tag-specifications", "", "The tags to apply to the AMI and snapshots on creation.")
	ec2_createImageCmd.MarkFlagRequired("instance-id")
	ec2_createImageCmd.MarkFlagRequired("name")
	ec2_createImageCmd.Flag("no-dry-run").Hidden = true
	ec2_createImageCmd.Flag("no-no-reboot").Hidden = true
	ec2Cmd.AddCommand(ec2_createImageCmd)
}
