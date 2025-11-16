package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeExportTasksCmd = &cobra.Command{
	Use:   "describe-export-tasks",
	Short: "Returns information about a snapshot or cluster export to Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeExportTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeExportTasksCmd).Standalone()

		rds_describeExportTasksCmd.Flags().String("export-task-identifier", "", "The identifier of the snapshot or cluster export task to be described.")
		rds_describeExportTasksCmd.Flags().String("filters", "", "Filters specify one or more snapshot or cluster exports to describe.")
		rds_describeExportTasksCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeExportTasks` request.")
		rds_describeExportTasksCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		rds_describeExportTasksCmd.Flags().String("source-arn", "", "The Amazon Resource Name (ARN) of the snapshot or cluster exported to Amazon S3.")
		rds_describeExportTasksCmd.Flags().String("source-type", "", "The type of source for the export.")
	})
	rdsCmd.AddCommand(rds_describeExportTasksCmd)
}
