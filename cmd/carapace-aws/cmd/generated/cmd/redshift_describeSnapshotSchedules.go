package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeSnapshotSchedulesCmd = &cobra.Command{
	Use:   "describe-snapshot-schedules",
	Short: "Returns a list of snapshot schedules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeSnapshotSchedulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeSnapshotSchedulesCmd).Standalone()

		redshift_describeSnapshotSchedulesCmd.Flags().String("cluster-identifier", "", "The unique identifier for the cluster whose snapshot schedules you want to view.")
		redshift_describeSnapshotSchedulesCmd.Flags().String("marker", "", "A value that indicates the starting point for the next set of response records in a subsequent request.")
		redshift_describeSnapshotSchedulesCmd.Flags().String("max-records", "", "The maximum number or response records to return in each call.")
		redshift_describeSnapshotSchedulesCmd.Flags().String("schedule-identifier", "", "A unique identifier for a snapshot schedule.")
		redshift_describeSnapshotSchedulesCmd.Flags().String("tag-keys", "", "The key value for a snapshot schedule tag.")
		redshift_describeSnapshotSchedulesCmd.Flags().String("tag-values", "", "The value corresponding to the key of the snapshot schedule tag.")
	})
	redshiftCmd.AddCommand(redshift_describeSnapshotSchedulesCmd)
}
