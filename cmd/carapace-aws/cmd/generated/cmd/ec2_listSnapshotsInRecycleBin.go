package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_listSnapshotsInRecycleBinCmd = &cobra.Command{
	Use:   "list-snapshots-in-recycle-bin",
	Short: "Lists one or more snapshots that are currently in the Recycle Bin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_listSnapshotsInRecycleBinCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_listSnapshotsInRecycleBinCmd).Standalone()

		ec2_listSnapshotsInRecycleBinCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_listSnapshotsInRecycleBinCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_listSnapshotsInRecycleBinCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_listSnapshotsInRecycleBinCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_listSnapshotsInRecycleBinCmd.Flags().String("snapshot-ids", "", "The IDs of the snapshots to list.")
		ec2_listSnapshotsInRecycleBinCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_listSnapshotsInRecycleBinCmd)
}
