package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getOtaTaskConfigurationCmd = &cobra.Command{
	Use:   "get-ota-task-configuration",
	Short: "Get a configuraiton for the over-the-air (OTA) task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getOtaTaskConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_getOtaTaskConfigurationCmd).Standalone()

		iotManagedIntegrations_getOtaTaskConfigurationCmd.Flags().String("identifier", "", "The over-the-air (OTA) task configuration id.")
		iotManagedIntegrations_getOtaTaskConfigurationCmd.MarkFlagRequired("identifier")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getOtaTaskConfigurationCmd)
}
