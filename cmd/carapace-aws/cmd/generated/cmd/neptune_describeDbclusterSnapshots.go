package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeDbclusterSnapshotsCmd = &cobra.Command{
	Use:   "describe-dbcluster-snapshots",
	Short: "Returns information about DB cluster snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeDbclusterSnapshotsCmd).Standalone()

	neptune_describeDbclusterSnapshotsCmd.Flags().String("dbcluster-identifier", "", "The ID of the DB cluster to retrieve the list of DB cluster snapshots for.")
	neptune_describeDbclusterSnapshotsCmd.Flags().String("dbcluster-snapshot-identifier", "", "A specific DB cluster snapshot identifier to describe.")
	neptune_describeDbclusterSnapshotsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
	neptune_describeDbclusterSnapshotsCmd.Flags().Bool("include-public", false, "True to include manual DB cluster snapshots that are public and can be copied or restored by any Amazon account, and otherwise false.")
	neptune_describeDbclusterSnapshotsCmd.Flags().Bool("include-shared", false, "True to include shared manual DB cluster snapshots from other Amazon accounts that this Amazon account has been given permission to copy or restore, and otherwise false.")
	neptune_describeDbclusterSnapshotsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBClusterSnapshots` request.")
	neptune_describeDbclusterSnapshotsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	neptune_describeDbclusterSnapshotsCmd.Flags().Bool("no-include-public", false, "True to include manual DB cluster snapshots that are public and can be copied or restored by any Amazon account, and otherwise false.")
	neptune_describeDbclusterSnapshotsCmd.Flags().Bool("no-include-shared", false, "True to include shared manual DB cluster snapshots from other Amazon accounts that this Amazon account has been given permission to copy or restore, and otherwise false.")
	neptune_describeDbclusterSnapshotsCmd.Flags().String("snapshot-type", "", "The type of DB cluster snapshots to be returned.")
	neptune_describeDbclusterSnapshotsCmd.Flag("no-include-public").Hidden = true
	neptune_describeDbclusterSnapshotsCmd.Flag("no-include-shared").Hidden = true
	neptuneCmd.AddCommand(neptune_describeDbclusterSnapshotsCmd)
}
