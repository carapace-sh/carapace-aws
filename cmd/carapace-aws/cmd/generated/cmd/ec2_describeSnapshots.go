package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeSnapshotsCmd = &cobra.Command{
	Use:   "describe-snapshots",
	Short: "Describes the specified EBS snapshots available to you or all of the EBS snapshots available to you.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeSnapshotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeSnapshotsCmd).Standalone()

		ec2_describeSnapshotsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSnapshotsCmd.Flags().String("filters", "", "The filters.")
		ec2_describeSnapshotsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeSnapshotsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeSnapshotsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeSnapshotsCmd.Flags().String("owner-ids", "", "Scopes the results to snapshots with the specified owners.")
		ec2_describeSnapshotsCmd.Flags().String("restorable-by-user-ids", "", "The IDs of the Amazon Web Services accounts that can create volumes from the snapshot.")
		ec2_describeSnapshotsCmd.Flags().String("snapshot-ids", "", "The snapshot IDs.")
		ec2_describeSnapshotsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeSnapshotsCmd)
}
