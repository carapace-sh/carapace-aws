package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getOtaTaskCmd = &cobra.Command{
	Use:   "get-ota-task",
	Short: "Get details of the over-the-air (OTA) task by its task id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getOtaTaskCmd).Standalone()

	iotManagedIntegrations_getOtaTaskCmd.Flags().String("identifier", "", "The over-the-air (OTA) task id.")
	iotManagedIntegrations_getOtaTaskCmd.MarkFlagRequired("identifier")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getOtaTaskCmd)
}
