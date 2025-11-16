package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_restoreSnapshotTierCmd = &cobra.Command{
	Use:   "restore-snapshot-tier",
	Short: "Restores an archived Amazon EBS snapshot for use temporarily or permanently, or modifies the restore period or restore type for a snapshot that was previously temporarily restored.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_restoreSnapshotTierCmd).Standalone()

	ec2_restoreSnapshotTierCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_restoreSnapshotTierCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_restoreSnapshotTierCmd.Flags().Bool("no-permanent-restore", false, "Indicates whether to permanently restore an archived snapshot.")
	ec2_restoreSnapshotTierCmd.Flags().Bool("permanent-restore", false, "Indicates whether to permanently restore an archived snapshot.")
	ec2_restoreSnapshotTierCmd.Flags().String("snapshot-id", "", "The ID of the snapshot to restore.")
	ec2_restoreSnapshotTierCmd.Flags().String("temporary-restore-days", "", "Specifies the number of days for which to temporarily restore an archived snapshot.")
	ec2_restoreSnapshotTierCmd.Flag("no-dry-run").Hidden = true
	ec2_restoreSnapshotTierCmd.Flag("no-permanent-restore").Hidden = true
	ec2_restoreSnapshotTierCmd.MarkFlagRequired("snapshot-id")
	ec2Cmd.AddCommand(ec2_restoreSnapshotTierCmd)
}
