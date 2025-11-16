package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbsnapshotsCmd = &cobra.Command{
	Use:   "describe-dbsnapshots",
	Short: "Returns information about DB snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbsnapshotsCmd).Standalone()

	rds_describeDbsnapshotsCmd.Flags().String("dbi-resource-id", "", "A specific DB resource ID to describe.")
	rds_describeDbsnapshotsCmd.Flags().String("dbinstance-identifier", "", "The ID of the DB instance to retrieve the list of DB snapshots for.")
	rds_describeDbsnapshotsCmd.Flags().String("dbsnapshot-identifier", "", "A specific DB snapshot identifier to describe.")
	rds_describeDbsnapshotsCmd.Flags().String("filters", "", "A filter that specifies one or more DB snapshots to describe.")
	rds_describeDbsnapshotsCmd.Flags().Bool("include-public", false, "Specifies whether to include manual DB cluster snapshots that are public and can be copied or restored by any Amazon Web Services account.")
	rds_describeDbsnapshotsCmd.Flags().Bool("include-shared", false, "Specifies whether to include shared manual DB cluster snapshots from other Amazon Web Services accounts that this Amazon Web Services account has been given permission to copy or restore.")
	rds_describeDbsnapshotsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBSnapshots` request.")
	rds_describeDbsnapshotsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rds_describeDbsnapshotsCmd.Flags().Bool("no-include-public", false, "Specifies whether to include manual DB cluster snapshots that are public and can be copied or restored by any Amazon Web Services account.")
	rds_describeDbsnapshotsCmd.Flags().Bool("no-include-shared", false, "Specifies whether to include shared manual DB cluster snapshots from other Amazon Web Services accounts that this Amazon Web Services account has been given permission to copy or restore.")
	rds_describeDbsnapshotsCmd.Flags().String("snapshot-type", "", "The type of snapshots to be returned.")
	rds_describeDbsnapshotsCmd.Flag("no-include-public").Hidden = true
	rds_describeDbsnapshotsCmd.Flag("no-include-shared").Hidden = true
	rdsCmd.AddCommand(rds_describeDbsnapshotsCmd)
}
