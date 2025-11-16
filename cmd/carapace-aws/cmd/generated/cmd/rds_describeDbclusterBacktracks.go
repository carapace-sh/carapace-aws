package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbclusterBacktracksCmd = &cobra.Command{
	Use:   "describe-dbcluster-backtracks",
	Short: "Returns information about backtracks for a DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbclusterBacktracksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeDbclusterBacktracksCmd).Standalone()

		rds_describeDbclusterBacktracksCmd.Flags().String("backtrack-identifier", "", "If specified, this value is the backtrack identifier of the backtrack to be described.")
		rds_describeDbclusterBacktracksCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier of the DB cluster to be described.")
		rds_describeDbclusterBacktracksCmd.Flags().String("filters", "", "A filter that specifies one or more DB clusters to describe.")
		rds_describeDbclusterBacktracksCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBClusterBacktracks` request.")
		rds_describeDbclusterBacktracksCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		rds_describeDbclusterBacktracksCmd.MarkFlagRequired("dbcluster-identifier")
	})
	rdsCmd.AddCommand(rds_describeDbclusterBacktracksCmd)
}
