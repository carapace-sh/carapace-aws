package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createSnapshotCmd = &cobra.Command{
	Use:   "create-snapshot",
	Short: "Creates a snapshot of an EBS volume and stores it in Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createSnapshotCmd).Standalone()

		ec2_createSnapshotCmd.Flags().String("description", "", "A description for the snapshot.")
		ec2_createSnapshotCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createSnapshotCmd.Flags().String("location", "", "Only supported for volumes in Local Zones.")
		ec2_createSnapshotCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createSnapshotCmd.Flags().String("outpost-arn", "", "Only supported for volumes on Outposts.")
		ec2_createSnapshotCmd.Flags().String("tag-specifications", "", "The tags to apply to the snapshot during creation.")
		ec2_createSnapshotCmd.Flags().String("volume-id", "", "The ID of the Amazon EBS volume.")
		ec2_createSnapshotCmd.Flag("no-dry-run").Hidden = true
		ec2_createSnapshotCmd.MarkFlagRequired("volume-id")
	})
	ec2Cmd.AddCommand(ec2_createSnapshotCmd)
}
