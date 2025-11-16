package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_createManagedThingCmd = &cobra.Command{
	Use:   "create-managed-thing",
	Short: "Creates a managed thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_createManagedThingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_createManagedThingCmd).Standalone()

		iotManagedIntegrations_createManagedThingCmd.Flags().String("authentication-material", "", "The authentication material defining the device connectivity setup requests.")
		iotManagedIntegrations_createManagedThingCmd.Flags().String("authentication-material-type", "", "The type of authentication material used for device connectivity setup requests.")
		iotManagedIntegrations_createManagedThingCmd.Flags().String("brand", "", "The brand of the device.")
		iotManagedIntegrations_createManagedThingCmd.Flags().String("capabilities", "", "The capabilities of the device such as light bulb.")
		iotManagedIntegrations_createManagedThingCmd.Flags().String("capability-report", "", "A report of the capabilities for the managed thing.")
		iotManagedIntegrations_createManagedThingCmd.Flags().String("capability-schemas", "", "The capability schemas that define the functionality and features supported by the managed thing, including device capabilities and their associated properties.")
		iotManagedIntegrations_createManagedThingCmd.Flags().String("classification", "", "The classification of the managed thing such as light bulb or thermostat.")
		iotManagedIntegrations_createManagedThingCmd.Flags().String("client-token", "", "An idempotency token.")
		iotManagedIntegrations_createManagedThingCmd.Flags().String("credential-locker-id", "", "The identifier of the credential for the managed thing.")
		iotManagedIntegrations_createManagedThingCmd.Flags().String("meta-data", "", "The metadata for the managed thing.")
		iotManagedIntegrations_createManagedThingCmd.Flags().String("model", "", "The model of the device.")
		iotManagedIntegrations_createManagedThingCmd.Flags().String("name", "", "The name of the managed thing representing the physical device.")
		iotManagedIntegrations_createManagedThingCmd.Flags().String("owner", "", "Owner of the device, usually an indication of whom the device belongs to.")
		iotManagedIntegrations_createManagedThingCmd.Flags().String("role", "", "The type of device used.")
		iotManagedIntegrations_createManagedThingCmd.Flags().String("serial-number", "", "The serial number of the device.")
		iotManagedIntegrations_createManagedThingCmd.Flags().String("tags", "", "A set of key/value pairs that are used to manage the managed thing.")
		iotManagedIntegrations_createManagedThingCmd.MarkFlagRequired("authentication-material")
		iotManagedIntegrations_createManagedThingCmd.MarkFlagRequired("authentication-material-type")
		iotManagedIntegrations_createManagedThingCmd.MarkFlagRequired("role")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_createManagedThingCmd)
}
