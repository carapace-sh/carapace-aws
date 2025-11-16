package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_describeSnapshotsCmd = &cobra.Command{
	Use:   "describe-snapshots",
	Short: "Returns the description of specific Amazon FSx for OpenZFS snapshots, if a `SnapshotIds` value is provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_describeSnapshotsCmd).Standalone()

	fsx_describeSnapshotsCmd.Flags().String("filters", "", "The filters structure.")
	fsx_describeSnapshotsCmd.Flags().String("include-shared", "", "Set to `false` (default) if you want to only see the snapshots owned by your Amazon Web Services account.")
	fsx_describeSnapshotsCmd.Flags().String("max-results", "", "")
	fsx_describeSnapshotsCmd.Flags().String("next-token", "", "")
	fsx_describeSnapshotsCmd.Flags().String("snapshot-ids", "", "The IDs of the snapshots that you want to retrieve.")
	fsxCmd.AddCommand(fsx_describeSnapshotsCmd)
}
