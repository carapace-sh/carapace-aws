package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_deleteOtaTaskCmd = &cobra.Command{
	Use:   "delete-ota-task",
	Short: "Delete the over-the-air (OTA) task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_deleteOtaTaskCmd).Standalone()

	iotManagedIntegrations_deleteOtaTaskCmd.Flags().String("identifier", "", "The identifier of the over-the-air (OTA) task.")
	iotManagedIntegrations_deleteOtaTaskCmd.MarkFlagRequired("identifier")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_deleteOtaTaskCmd)
}
