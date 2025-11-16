package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeSnapshotsCmd = &cobra.Command{
	Use:   "describe-snapshots",
	Short: "Obtains information about the directory snapshots that belong to this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeSnapshotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_describeSnapshotsCmd).Standalone()

		ds_describeSnapshotsCmd.Flags().String("directory-id", "", "The identifier of the directory for which to retrieve snapshot information.")
		ds_describeSnapshotsCmd.Flags().String("limit", "", "The maximum number of objects to return.")
		ds_describeSnapshotsCmd.Flags().String("next-token", "", "The *DescribeSnapshotsResult.NextToken* value from a previous call to [DescribeSnapshots]().")
		ds_describeSnapshotsCmd.Flags().String("snapshot-ids", "", "A list of identifiers of the snapshots to obtain the information for.")
	})
	dsCmd.AddCommand(ds_describeSnapshotsCmd)
}
