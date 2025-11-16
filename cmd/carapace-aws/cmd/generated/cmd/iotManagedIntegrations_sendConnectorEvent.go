package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_sendConnectorEventCmd = &cobra.Command{
	Use:   "send-connector-event",
	Short: "Relays third-party device events for a connector such as a new device or a device state change event.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_sendConnectorEventCmd).Standalone()

	iotManagedIntegrations_sendConnectorEventCmd.Flags().String("connector-device-id", "", "The third-party device id as defined by the connector.")
	iotManagedIntegrations_sendConnectorEventCmd.Flags().String("connector-id", "", "The id of the connector between the third-party cloud provider and IoT managed integrations.")
	iotManagedIntegrations_sendConnectorEventCmd.Flags().String("device-discovery-id", "", "The id for the device discovery job.")
	iotManagedIntegrations_sendConnectorEventCmd.Flags().String("devices", "", "The list of devices.")
	iotManagedIntegrations_sendConnectorEventCmd.Flags().String("matter-endpoint", "", "The device endpoint.")
	iotManagedIntegrations_sendConnectorEventCmd.Flags().String("message", "", "The device state change event payload.")
	iotManagedIntegrations_sendConnectorEventCmd.Flags().String("operation", "", "The Open Connectivity Foundation (OCF) operation requested to be performed on the managed thing.")
	iotManagedIntegrations_sendConnectorEventCmd.Flags().String("operation-version", "", "The Open Connectivity Foundation (OCF) security specification version for the operation being requested on the managed thing.")
	iotManagedIntegrations_sendConnectorEventCmd.Flags().String("status-code", "", "The status code of the Open Connectivity Foundation (OCF) operation being performed on the managed thing.")
	iotManagedIntegrations_sendConnectorEventCmd.Flags().String("trace-id", "", "The trace request identifier.")
	iotManagedIntegrations_sendConnectorEventCmd.Flags().String("user-id", "", "The id of the third-party cloud provider.")
	iotManagedIntegrations_sendConnectorEventCmd.MarkFlagRequired("connector-id")
	iotManagedIntegrations_sendConnectorEventCmd.MarkFlagRequired("operation")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_sendConnectorEventCmd)
}
