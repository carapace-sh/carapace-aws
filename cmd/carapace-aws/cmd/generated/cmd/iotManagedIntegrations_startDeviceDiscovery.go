package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_startDeviceDiscoveryCmd = &cobra.Command{
	Use:   "start-device-discovery",
	Short: "This API is used to start device discovery for hub-connected and third-party-connected devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_startDeviceDiscoveryCmd).Standalone()

	iotManagedIntegrations_startDeviceDiscoveryCmd.Flags().String("account-association-id", "", "The identifier of the cloud-to-cloud account association to use for discovery of third-party devices.")
	iotManagedIntegrations_startDeviceDiscoveryCmd.Flags().String("authentication-material", "", "The authentication material required to start the local device discovery job request.")
	iotManagedIntegrations_startDeviceDiscoveryCmd.Flags().String("authentication-material-type", "", "The type of authentication material used for device discovery jobs.")
	iotManagedIntegrations_startDeviceDiscoveryCmd.Flags().String("client-token", "", "An idempotency token.")
	iotManagedIntegrations_startDeviceDiscoveryCmd.Flags().String("connector-association-identifier", "", "The id of the connector association.")
	iotManagedIntegrations_startDeviceDiscoveryCmd.Flags().String("controller-identifier", "", "The id of the end-user's IoT hub.")
	iotManagedIntegrations_startDeviceDiscoveryCmd.Flags().String("custom-protocol-detail", "", "Additional protocol-specific details required for device discovery, which vary based on the discovery type.")
	iotManagedIntegrations_startDeviceDiscoveryCmd.Flags().String("discovery-type", "", "The discovery type supporting the type of device to be discovered in the device discovery task request.")
	iotManagedIntegrations_startDeviceDiscoveryCmd.Flags().String("tags", "", "A set of key/value pairs that are used to manage the device discovery request.")
	iotManagedIntegrations_startDeviceDiscoveryCmd.MarkFlagRequired("discovery-type")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_startDeviceDiscoveryCmd)
}
