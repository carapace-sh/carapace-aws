package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteColumnStatisticsTaskSettingsCmd = &cobra.Command{
	Use:   "delete-column-statistics-task-settings",
	Short: "Deletes settings for a column statistics task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteColumnStatisticsTaskSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_deleteColumnStatisticsTaskSettingsCmd).Standalone()

		glue_deleteColumnStatisticsTaskSettingsCmd.Flags().String("database-name", "", "The name of the database where the table resides.")
		glue_deleteColumnStatisticsTaskSettingsCmd.Flags().String("table-name", "", "The name of the table for which to delete column statistics.")
		glue_deleteColumnStatisticsTaskSettingsCmd.MarkFlagRequired("database-name")
		glue_deleteColumnStatisticsTaskSettingsCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_deleteColumnStatisticsTaskSettingsCmd)
}
