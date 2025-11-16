package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbclusterSnapshotsCmd = &cobra.Command{
	Use:   "describe-dbcluster-snapshots",
	Short: "Returns information about DB cluster snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbclusterSnapshotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeDbclusterSnapshotsCmd).Standalone()

		rds_describeDbclusterSnapshotsCmd.Flags().String("db-cluster-resource-id", "", "A specific DB cluster resource ID to describe.")
		rds_describeDbclusterSnapshotsCmd.Flags().String("dbcluster-identifier", "", "The ID of the DB cluster to retrieve the list of DB cluster snapshots for.")
		rds_describeDbclusterSnapshotsCmd.Flags().String("dbcluster-snapshot-identifier", "", "A specific DB cluster snapshot identifier to describe.")
		rds_describeDbclusterSnapshotsCmd.Flags().String("filters", "", "A filter that specifies one or more DB cluster snapshots to describe.")
		rds_describeDbclusterSnapshotsCmd.Flags().Bool("include-public", false, "Specifies whether to include manual DB cluster snapshots that are public and can be copied or restored by any Amazon Web Services account.")
		rds_describeDbclusterSnapshotsCmd.Flags().Bool("include-shared", false, "Specifies whether to include shared manual DB cluster snapshots from other Amazon Web Services accounts that this Amazon Web Services account has been given permission to copy or restore.")
		rds_describeDbclusterSnapshotsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBClusterSnapshots` request.")
		rds_describeDbclusterSnapshotsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		rds_describeDbclusterSnapshotsCmd.Flags().Bool("no-include-public", false, "Specifies whether to include manual DB cluster snapshots that are public and can be copied or restored by any Amazon Web Services account.")
		rds_describeDbclusterSnapshotsCmd.Flags().Bool("no-include-shared", false, "Specifies whether to include shared manual DB cluster snapshots from other Amazon Web Services accounts that this Amazon Web Services account has been given permission to copy or restore.")
		rds_describeDbclusterSnapshotsCmd.Flags().String("snapshot-type", "", "The type of DB cluster snapshots to be returned.")
		rds_describeDbclusterSnapshotsCmd.Flag("no-include-public").Hidden = true
		rds_describeDbclusterSnapshotsCmd.Flag("no-include-shared").Hidden = true
	})
	rdsCmd.AddCommand(rds_describeDbclusterSnapshotsCmd)
}
