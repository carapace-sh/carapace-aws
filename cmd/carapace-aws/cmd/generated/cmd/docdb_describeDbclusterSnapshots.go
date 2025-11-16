package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describeDbclusterSnapshotsCmd = &cobra.Command{
	Use:   "describe-dbcluster-snapshots",
	Short: "Returns information about cluster snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describeDbclusterSnapshotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_describeDbclusterSnapshotsCmd).Standalone()

		docdb_describeDbclusterSnapshotsCmd.Flags().String("dbcluster-identifier", "", "The ID of the cluster to retrieve the list of cluster snapshots for.")
		docdb_describeDbclusterSnapshotsCmd.Flags().String("dbcluster-snapshot-identifier", "", "A specific cluster snapshot identifier to describe.")
		docdb_describeDbclusterSnapshotsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
		docdb_describeDbclusterSnapshotsCmd.Flags().Bool("include-public", false, "Set to `true` to include manual cluster snapshots that are public and can be copied or restored by any Amazon Web Services account, and otherwise `false`.")
		docdb_describeDbclusterSnapshotsCmd.Flags().Bool("include-shared", false, "Set to `true` to include shared manual cluster snapshots from other Amazon Web Services accounts that this Amazon Web Services account has been given permission to copy or restore, and otherwise `false`.")
		docdb_describeDbclusterSnapshotsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		docdb_describeDbclusterSnapshotsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		docdb_describeDbclusterSnapshotsCmd.Flags().Bool("no-include-public", false, "Set to `true` to include manual cluster snapshots that are public and can be copied or restored by any Amazon Web Services account, and otherwise `false`.")
		docdb_describeDbclusterSnapshotsCmd.Flags().Bool("no-include-shared", false, "Set to `true` to include shared manual cluster snapshots from other Amazon Web Services accounts that this Amazon Web Services account has been given permission to copy or restore, and otherwise `false`.")
		docdb_describeDbclusterSnapshotsCmd.Flags().String("snapshot-type", "", "The type of cluster snapshots to be returned.")
		docdb_describeDbclusterSnapshotsCmd.Flag("no-include-public").Hidden = true
		docdb_describeDbclusterSnapshotsCmd.Flag("no-include-shared").Hidden = true
	})
	docdbCmd.AddCommand(docdb_describeDbclusterSnapshotsCmd)
}
