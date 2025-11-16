package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_startColumnStatisticsTaskRunScheduleCmd = &cobra.Command{
	Use:   "start-column-statistics-task-run-schedule",
	Short: "Starts a column statistics task run schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_startColumnStatisticsTaskRunScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_startColumnStatisticsTaskRunScheduleCmd).Standalone()

		glue_startColumnStatisticsTaskRunScheduleCmd.Flags().String("database-name", "", "The name of the database where the table resides.")
		glue_startColumnStatisticsTaskRunScheduleCmd.Flags().String("table-name", "", "The name of the table for which to start a column statistic task run schedule.")
		glue_startColumnStatisticsTaskRunScheduleCmd.MarkFlagRequired("database-name")
		glue_startColumnStatisticsTaskRunScheduleCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_startColumnStatisticsTaskRunScheduleCmd)
}
