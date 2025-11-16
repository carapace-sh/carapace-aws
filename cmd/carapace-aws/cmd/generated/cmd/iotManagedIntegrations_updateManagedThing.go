package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_updateManagedThingCmd = &cobra.Command{
	Use:   "update-managed-thing",
	Short: "Update the attributes and capabilities associated with a managed thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_updateManagedThingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_updateManagedThingCmd).Standalone()

		iotManagedIntegrations_updateManagedThingCmd.Flags().String("brand", "", "The brand of the device.")
		iotManagedIntegrations_updateManagedThingCmd.Flags().String("capabilities", "", "The capabilities of the device such as light bulb.")
		iotManagedIntegrations_updateManagedThingCmd.Flags().String("capability-report", "", "A report of the capabilities for the managed thing.")
		iotManagedIntegrations_updateManagedThingCmd.Flags().String("capability-schemas", "", "The updated capability schemas that define the functionality and features supported by the managed thing.")
		iotManagedIntegrations_updateManagedThingCmd.Flags().String("classification", "", "The classification of the managed thing such as light bulb or thermostat.")
		iotManagedIntegrations_updateManagedThingCmd.Flags().String("credential-locker-id", "", "The identifier of the credential for the managed thing.")
		iotManagedIntegrations_updateManagedThingCmd.Flags().String("hub-network-mode", "", "The network mode for the hub-connected device.")
		iotManagedIntegrations_updateManagedThingCmd.Flags().String("identifier", "", "The id of the managed thing.")
		iotManagedIntegrations_updateManagedThingCmd.Flags().String("meta-data", "", "The metadata for the managed thing.")
		iotManagedIntegrations_updateManagedThingCmd.Flags().String("model", "", "The model of the device.")
		iotManagedIntegrations_updateManagedThingCmd.Flags().String("name", "", "The name of the managed thing representing the physical device.")
		iotManagedIntegrations_updateManagedThingCmd.Flags().String("owner", "", "Owner of the device, usually an indication of whom the device belongs to.")
		iotManagedIntegrations_updateManagedThingCmd.Flags().String("serial-number", "", "The serial number of the device.")
		iotManagedIntegrations_updateManagedThingCmd.MarkFlagRequired("identifier")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_updateManagedThingCmd)
}
