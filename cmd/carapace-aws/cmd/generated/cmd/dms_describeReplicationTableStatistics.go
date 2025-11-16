package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeReplicationTableStatisticsCmd = &cobra.Command{
	Use:   "describe-replication-table-statistics",
	Short: "Returns table and schema statistics for one or more provisioned replications that use a given DMS Serverless replication configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeReplicationTableStatisticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeReplicationTableStatisticsCmd).Standalone()

		dms_describeReplicationTableStatisticsCmd.Flags().String("filters", "", "Filters applied to the replication table statistics.")
		dms_describeReplicationTableStatisticsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		dms_describeReplicationTableStatisticsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		dms_describeReplicationTableStatisticsCmd.Flags().String("replication-config-arn", "", "The replication config to describe.")
		dms_describeReplicationTableStatisticsCmd.MarkFlagRequired("replication-config-arn")
	})
	dmsCmd.AddCommand(dms_describeReplicationTableStatisticsCmd)
}
