package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_updateApplicationMaintenanceConfigurationCmd = &cobra.Command{
	Use:   "update-application-maintenance-configuration",
	Short: "Updates the maintenance configuration of the Managed Service for Apache Flink application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_updateApplicationMaintenanceConfigurationCmd).Standalone()

	kinesisanalyticsv2_updateApplicationMaintenanceConfigurationCmd.Flags().String("application-maintenance-configuration-update", "", "Describes the application maintenance configuration update.")
	kinesisanalyticsv2_updateApplicationMaintenanceConfigurationCmd.Flags().String("application-name", "", "The name of the application for which you want to update the maintenance configuration.")
	kinesisanalyticsv2_updateApplicationMaintenanceConfigurationCmd.MarkFlagRequired("application-maintenance-configuration-update")
	kinesisanalyticsv2_updateApplicationMaintenanceConfigurationCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_updateApplicationMaintenanceConfigurationCmd)
}
