package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_restoreSnapshotFromRecycleBinCmd = &cobra.Command{
	Use:   "restore-snapshot-from-recycle-bin",
	Short: "Restores a snapshot from the Recycle Bin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_restoreSnapshotFromRecycleBinCmd).Standalone()

	ec2_restoreSnapshotFromRecycleBinCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_restoreSnapshotFromRecycleBinCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_restoreSnapshotFromRecycleBinCmd.Flags().String("snapshot-id", "", "The ID of the snapshot to restore.")
	ec2_restoreSnapshotFromRecycleBinCmd.Flag("no-dry-run").Hidden = true
	ec2_restoreSnapshotFromRecycleBinCmd.MarkFlagRequired("snapshot-id")
	ec2Cmd.AddCommand(ec2_restoreSnapshotFromRecycleBinCmd)
}
