package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_updateOtaTaskCmd = &cobra.Command{
	Use:   "update-ota-task",
	Short: "Update an over-the-air (OTA) task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_updateOtaTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_updateOtaTaskCmd).Standalone()

		iotManagedIntegrations_updateOtaTaskCmd.Flags().String("description", "", "The description of the over-the-air (OTA) task.")
		iotManagedIntegrations_updateOtaTaskCmd.Flags().String("identifier", "", "The over-the-air (OTA) task id.")
		iotManagedIntegrations_updateOtaTaskCmd.Flags().String("task-configuration-id", "", "The identifier for the over-the-air (OTA) task configuration.")
		iotManagedIntegrations_updateOtaTaskCmd.MarkFlagRequired("identifier")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_updateOtaTaskCmd)
}
