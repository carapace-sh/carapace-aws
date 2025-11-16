package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_deleteOtaTaskConfigurationCmd = &cobra.Command{
	Use:   "delete-ota-task-configuration",
	Short: "Delete the over-the-air (OTA) task configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_deleteOtaTaskConfigurationCmd).Standalone()

	iotManagedIntegrations_deleteOtaTaskConfigurationCmd.Flags().String("identifier", "", "The identifier of the over-the-air (OTA) task configuration.")
	iotManagedIntegrations_deleteOtaTaskConfigurationCmd.MarkFlagRequired("identifier")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_deleteOtaTaskConfigurationCmd)
}
