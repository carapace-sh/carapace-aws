package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_getTableBucketMaintenanceConfigurationCmd = &cobra.Command{
	Use:   "get-table-bucket-maintenance-configuration",
	Short: "Gets details about a maintenance configuration for a given table bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_getTableBucketMaintenanceConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_getTableBucketMaintenanceConfigurationCmd).Standalone()

		s3tables_getTableBucketMaintenanceConfigurationCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket associated with the maintenance configuration.")
		s3tables_getTableBucketMaintenanceConfigurationCmd.MarkFlagRequired("table-bucket-arn")
	})
	s3tablesCmd.AddCommand(s3tables_getTableBucketMaintenanceConfigurationCmd)
}
