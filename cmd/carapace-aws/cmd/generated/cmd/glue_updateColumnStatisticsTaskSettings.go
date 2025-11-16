package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateColumnStatisticsTaskSettingsCmd = &cobra.Command{
	Use:   "update-column-statistics-task-settings",
	Short: "Updates settings for a column statistics task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateColumnStatisticsTaskSettingsCmd).Standalone()

	glue_updateColumnStatisticsTaskSettingsCmd.Flags().String("catalog-id", "", "The ID of the Data Catalog in which the database resides.")
	glue_updateColumnStatisticsTaskSettingsCmd.Flags().String("column-name-list", "", "A list of column names for which to run statistics.")
	glue_updateColumnStatisticsTaskSettingsCmd.Flags().String("database-name", "", "The name of the database where the table resides.")
	glue_updateColumnStatisticsTaskSettingsCmd.Flags().String("role", "", "The role used for running the column statistics.")
	glue_updateColumnStatisticsTaskSettingsCmd.Flags().String("sample-size", "", "The percentage of data to sample.")
	glue_updateColumnStatisticsTaskSettingsCmd.Flags().String("schedule", "", "A schedule for running the column statistics, specified in CRON syntax.")
	glue_updateColumnStatisticsTaskSettingsCmd.Flags().String("security-configuration", "", "Name of the security configuration that is used to encrypt CloudWatch logs.")
	glue_updateColumnStatisticsTaskSettingsCmd.Flags().String("table-name", "", "The name of the table for which to generate column statistics.")
	glue_updateColumnStatisticsTaskSettingsCmd.MarkFlagRequired("database-name")
	glue_updateColumnStatisticsTaskSettingsCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_updateColumnStatisticsTaskSettingsCmd)
}
