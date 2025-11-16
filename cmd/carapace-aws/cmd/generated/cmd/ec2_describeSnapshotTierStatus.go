package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeSnapshotTierStatusCmd = &cobra.Command{
	Use:   "describe-snapshot-tier-status",
	Short: "Describes the storage tier status of one or more Amazon EBS snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeSnapshotTierStatusCmd).Standalone()

	ec2_describeSnapshotTierStatusCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeSnapshotTierStatusCmd.Flags().String("filters", "", "The filters.")
	ec2_describeSnapshotTierStatusCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeSnapshotTierStatusCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeSnapshotTierStatusCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeSnapshotTierStatusCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeSnapshotTierStatusCmd)
}
