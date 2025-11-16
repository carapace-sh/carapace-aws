package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getColumnStatisticsTaskRunsCmd = &cobra.Command{
	Use:   "get-column-statistics-task-runs",
	Short: "Retrieves information about all runs associated with the specified table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getColumnStatisticsTaskRunsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getColumnStatisticsTaskRunsCmd).Standalone()

		glue_getColumnStatisticsTaskRunsCmd.Flags().String("database-name", "", "The name of the database where the table resides.")
		glue_getColumnStatisticsTaskRunsCmd.Flags().String("max-results", "", "The maximum size of the response.")
		glue_getColumnStatisticsTaskRunsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
		glue_getColumnStatisticsTaskRunsCmd.Flags().String("table-name", "", "The name of the table.")
		glue_getColumnStatisticsTaskRunsCmd.MarkFlagRequired("database-name")
		glue_getColumnStatisticsTaskRunsCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_getColumnStatisticsTaskRunsCmd)
}
