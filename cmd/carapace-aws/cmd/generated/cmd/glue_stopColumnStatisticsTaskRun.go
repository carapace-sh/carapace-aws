package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_stopColumnStatisticsTaskRunCmd = &cobra.Command{
	Use:   "stop-column-statistics-task-run",
	Short: "Stops a task run for the specified table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_stopColumnStatisticsTaskRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_stopColumnStatisticsTaskRunCmd).Standalone()

		glue_stopColumnStatisticsTaskRunCmd.Flags().String("database-name", "", "The name of the database where the table resides.")
		glue_stopColumnStatisticsTaskRunCmd.Flags().String("table-name", "", "The name of the table.")
		glue_stopColumnStatisticsTaskRunCmd.MarkFlagRequired("database-name")
		glue_stopColumnStatisticsTaskRunCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_stopColumnStatisticsTaskRunCmd)
}
