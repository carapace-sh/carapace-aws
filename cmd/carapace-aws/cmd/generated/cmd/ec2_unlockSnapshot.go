package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_unlockSnapshotCmd = &cobra.Command{
	Use:   "unlock-snapshot",
	Short: "Unlocks a snapshot that is locked in governance mode or that is locked in compliance mode but still in the cooling-off period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_unlockSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_unlockSnapshotCmd).Standalone()

		ec2_unlockSnapshotCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_unlockSnapshotCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_unlockSnapshotCmd.Flags().String("snapshot-id", "", "The ID of the snapshot to unlock.")
		ec2_unlockSnapshotCmd.Flag("no-dry-run").Hidden = true
		ec2_unlockSnapshotCmd.MarkFlagRequired("snapshot-id")
	})
	ec2Cmd.AddCommand(ec2_unlockSnapshotCmd)
}
