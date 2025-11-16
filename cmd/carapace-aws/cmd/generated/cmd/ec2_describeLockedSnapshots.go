package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeLockedSnapshotsCmd = &cobra.Command{
	Use:   "describe-locked-snapshots",
	Short: "Describes the lock status for a snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeLockedSnapshotsCmd).Standalone()

	ec2_describeLockedSnapshotsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeLockedSnapshotsCmd.Flags().String("filters", "", "The filters.")
	ec2_describeLockedSnapshotsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeLockedSnapshotsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeLockedSnapshotsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeLockedSnapshotsCmd.Flags().String("snapshot-ids", "", "The IDs of the snapshots for which to view the lock status.")
	ec2_describeLockedSnapshotsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeLockedSnapshotsCmd)
}
