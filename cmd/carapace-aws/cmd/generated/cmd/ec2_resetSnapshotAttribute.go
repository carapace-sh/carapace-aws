package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_resetSnapshotAttributeCmd = &cobra.Command{
	Use:   "reset-snapshot-attribute",
	Short: "Resets permission settings for the specified snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_resetSnapshotAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_resetSnapshotAttributeCmd).Standalone()

		ec2_resetSnapshotAttributeCmd.Flags().String("attribute", "", "The attribute to reset.")
		ec2_resetSnapshotAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_resetSnapshotAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_resetSnapshotAttributeCmd.Flags().String("snapshot-id", "", "The ID of the snapshot.")
		ec2_resetSnapshotAttributeCmd.MarkFlagRequired("attribute")
		ec2_resetSnapshotAttributeCmd.Flag("no-dry-run").Hidden = true
		ec2_resetSnapshotAttributeCmd.MarkFlagRequired("snapshot-id")
	})
	ec2Cmd.AddCommand(ec2_resetSnapshotAttributeCmd)
}
