package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_putTableMaintenanceConfigurationCmd = &cobra.Command{
	Use:   "put-table-maintenance-configuration",
	Short: "Creates a new maintenance configuration or replaces an existing maintenance configuration for a table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_putTableMaintenanceConfigurationCmd).Standalone()

	s3tables_putTableMaintenanceConfigurationCmd.Flags().String("name", "", "The name of the table.")
	s3tables_putTableMaintenanceConfigurationCmd.Flags().String("namespace", "", "The namespace of the table.")
	s3tables_putTableMaintenanceConfigurationCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table associated with the maintenance configuration.")
	s3tables_putTableMaintenanceConfigurationCmd.Flags().String("type", "", "The type of the maintenance configuration.")
	s3tables_putTableMaintenanceConfigurationCmd.Flags().String("value", "", "Defines the values of the maintenance configuration for the table.")
	s3tables_putTableMaintenanceConfigurationCmd.MarkFlagRequired("name")
	s3tables_putTableMaintenanceConfigurationCmd.MarkFlagRequired("namespace")
	s3tables_putTableMaintenanceConfigurationCmd.MarkFlagRequired("table-bucket-arn")
	s3tables_putTableMaintenanceConfigurationCmd.MarkFlagRequired("type")
	s3tables_putTableMaintenanceConfigurationCmd.MarkFlagRequired("value")
	s3tablesCmd.AddCommand(s3tables_putTableMaintenanceConfigurationCmd)
}
