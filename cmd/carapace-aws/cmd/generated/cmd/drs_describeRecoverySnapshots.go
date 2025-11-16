package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_describeRecoverySnapshotsCmd = &cobra.Command{
	Use:   "describe-recovery-snapshots",
	Short: "Lists all Recovery Snapshots for a single Source Server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_describeRecoverySnapshotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_describeRecoverySnapshotsCmd).Standalone()

		drs_describeRecoverySnapshotsCmd.Flags().String("filters", "", "A set of filters by which to return Recovery Snapshots.")
		drs_describeRecoverySnapshotsCmd.Flags().String("max-results", "", "Maximum number of Recovery Snapshots to retrieve.")
		drs_describeRecoverySnapshotsCmd.Flags().String("next-token", "", "The token of the next Recovery Snapshot to retrieve.")
		drs_describeRecoverySnapshotsCmd.Flags().String("order", "", "The sorted ordering by which to return Recovery Snapshots.")
		drs_describeRecoverySnapshotsCmd.Flags().String("source-server-id", "", "Filter Recovery Snapshots by Source Server ID.")
		drs_describeRecoverySnapshotsCmd.MarkFlagRequired("source-server-id")
	})
	drsCmd.AddCommand(drs_describeRecoverySnapshotsCmd)
}
