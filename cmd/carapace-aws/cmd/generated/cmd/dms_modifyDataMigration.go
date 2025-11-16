package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_modifyDataMigrationCmd = &cobra.Command{
	Use:   "modify-data-migration",
	Short: "Modifies an existing DMS data migration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_modifyDataMigrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_modifyDataMigrationCmd).Standalone()

		dms_modifyDataMigrationCmd.Flags().String("data-migration-identifier", "", "The identifier (name or ARN) of the data migration to modify.")
		dms_modifyDataMigrationCmd.Flags().String("data-migration-name", "", "The new name for the data migration.")
		dms_modifyDataMigrationCmd.Flags().String("data-migration-type", "", "The new migration type for the data migration.")
		dms_modifyDataMigrationCmd.Flags().String("enable-cloudwatch-logs", "", "Whether to enable Cloudwatch logs for the data migration.")
		dms_modifyDataMigrationCmd.Flags().String("number-of-jobs", "", "The number of parallel jobs that trigger parallel threads to unload the tables from the source, and then load them to the target.")
		dms_modifyDataMigrationCmd.Flags().String("selection-rules", "", "A JSON-formatted string that defines what objects to include and exclude from the migration.")
		dms_modifyDataMigrationCmd.Flags().String("service-access-role-arn", "", "The new service access role ARN for the data migration.")
		dms_modifyDataMigrationCmd.Flags().String("source-data-settings", "", "The new information about the source data provider for the data migration.")
		dms_modifyDataMigrationCmd.Flags().String("target-data-settings", "", "The new information about the target data provider for the data migration.")
		dms_modifyDataMigrationCmd.MarkFlagRequired("data-migration-identifier")
	})
	dmsCmd.AddCommand(dms_modifyDataMigrationCmd)
}
