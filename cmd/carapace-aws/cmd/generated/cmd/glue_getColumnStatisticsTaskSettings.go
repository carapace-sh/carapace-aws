package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getColumnStatisticsTaskSettingsCmd = &cobra.Command{
	Use:   "get-column-statistics-task-settings",
	Short: "Gets settings for a column statistics task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getColumnStatisticsTaskSettingsCmd).Standalone()

	glue_getColumnStatisticsTaskSettingsCmd.Flags().String("database-name", "", "The name of the database where the table resides.")
	glue_getColumnStatisticsTaskSettingsCmd.Flags().String("table-name", "", "The name of the table for which to retrieve column statistics.")
	glue_getColumnStatisticsTaskSettingsCmd.MarkFlagRequired("database-name")
	glue_getColumnStatisticsTaskSettingsCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_getColumnStatisticsTaskSettingsCmd)
}
