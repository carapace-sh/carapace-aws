package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_lockSnapshotCmd = &cobra.Command{
	Use:   "lock-snapshot",
	Short: "Locks an Amazon EBS snapshot in either *governance* or *compliance* mode to protect it against accidental or malicious deletions for a specific duration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_lockSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_lockSnapshotCmd).Standalone()

		ec2_lockSnapshotCmd.Flags().String("cool-off-period", "", "The cooling-off period during which you can unlock the snapshot or modify the lock settings after locking the snapshot in compliance mode, in hours.")
		ec2_lockSnapshotCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_lockSnapshotCmd.Flags().String("expiration-date", "", "The date and time at which the snapshot lock is to automatically expire, in the UTC time zone (`YYYY-MM-DDThh:mm:ss.sssZ`).")
		ec2_lockSnapshotCmd.Flags().String("lock-duration", "", "The period of time for which to lock the snapshot, in days.")
		ec2_lockSnapshotCmd.Flags().String("lock-mode", "", "The mode in which to lock the snapshot.")
		ec2_lockSnapshotCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_lockSnapshotCmd.Flags().String("snapshot-id", "", "The ID of the snapshot to lock.")
		ec2_lockSnapshotCmd.MarkFlagRequired("lock-mode")
		ec2_lockSnapshotCmd.Flag("no-dry-run").Hidden = true
		ec2_lockSnapshotCmd.MarkFlagRequired("snapshot-id")
	})
	ec2Cmd.AddCommand(ec2_lockSnapshotCmd)
}
