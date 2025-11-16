package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableFastSnapshotRestoresCmd = &cobra.Command{
	Use:   "enable-fast-snapshot-restores",
	Short: "Enables fast snapshot restores for the specified snapshots in the specified Availability Zones.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableFastSnapshotRestoresCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_enableFastSnapshotRestoresCmd).Standalone()

		ec2_enableFastSnapshotRestoresCmd.Flags().String("availability-zone-ids", "", "One or more Availability Zone IDs.")
		ec2_enableFastSnapshotRestoresCmd.Flags().String("availability-zones", "", "One or more Availability Zones.")
		ec2_enableFastSnapshotRestoresCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableFastSnapshotRestoresCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_enableFastSnapshotRestoresCmd.Flags().String("source-snapshot-ids", "", "The IDs of one or more snapshots.")
		ec2_enableFastSnapshotRestoresCmd.Flag("no-dry-run").Hidden = true
		ec2_enableFastSnapshotRestoresCmd.MarkFlagRequired("source-snapshot-ids")
	})
	ec2Cmd.AddCommand(ec2_enableFastSnapshotRestoresCmd)
}
