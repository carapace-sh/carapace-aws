package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getColumnStatisticsTaskRunCmd = &cobra.Command{
	Use:   "get-column-statistics-task-run",
	Short: "Get the associated metadata/information for a task run, given a task run ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getColumnStatisticsTaskRunCmd).Standalone()

	glue_getColumnStatisticsTaskRunCmd.Flags().String("column-statistics-task-run-id", "", "The identifier for the particular column statistics task run.")
	glue_getColumnStatisticsTaskRunCmd.MarkFlagRequired("column-statistics-task-run-id")
	glueCmd.AddCommand(glue_getColumnStatisticsTaskRunCmd)
}
