package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeTableStatisticsCmd = &cobra.Command{
	Use:   "describe-table-statistics",
	Short: "Returns table statistics on the database migration task, including table name, rows inserted, rows updated, and rows deleted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeTableStatisticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeTableStatisticsCmd).Standalone()

		dms_describeTableStatisticsCmd.Flags().String("filters", "", "Filters applied to table statistics.")
		dms_describeTableStatisticsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		dms_describeTableStatisticsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		dms_describeTableStatisticsCmd.Flags().String("replication-task-arn", "", "The Amazon Resource Name (ARN) of the replication task.")
		dms_describeTableStatisticsCmd.MarkFlagRequired("replication-task-arn")
	})
	dmsCmd.AddCommand(dms_describeTableStatisticsCmd)
}
