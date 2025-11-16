package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_stopColumnStatisticsTaskRunScheduleCmd = &cobra.Command{
	Use:   "stop-column-statistics-task-run-schedule",
	Short: "Stops a column statistics task run schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_stopColumnStatisticsTaskRunScheduleCmd).Standalone()

	glue_stopColumnStatisticsTaskRunScheduleCmd.Flags().String("database-name", "", "The name of the database where the table resides.")
	glue_stopColumnStatisticsTaskRunScheduleCmd.Flags().String("table-name", "", "The name of the table for which to stop a column statistic task run schedule.")
	glue_stopColumnStatisticsTaskRunScheduleCmd.MarkFlagRequired("database-name")
	glue_stopColumnStatisticsTaskRunScheduleCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_stopColumnStatisticsTaskRunScheduleCmd)
}
