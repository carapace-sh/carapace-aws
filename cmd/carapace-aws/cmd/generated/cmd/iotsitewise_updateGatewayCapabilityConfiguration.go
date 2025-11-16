package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_updateGatewayCapabilityConfigurationCmd = &cobra.Command{
	Use:   "update-gateway-capability-configuration",
	Short: "Updates a gateway capability configuration or defines a new capability configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_updateGatewayCapabilityConfigurationCmd).Standalone()

	iotsitewise_updateGatewayCapabilityConfigurationCmd.Flags().String("capability-configuration", "", "The JSON document that defines the configuration for the gateway capability.")
	iotsitewise_updateGatewayCapabilityConfigurationCmd.Flags().String("capability-namespace", "", "The namespace of the gateway capability configuration to be updated.")
	iotsitewise_updateGatewayCapabilityConfigurationCmd.Flags().String("gateway-id", "", "The ID of the gateway to be updated.")
	iotsitewise_updateGatewayCapabilityConfigurationCmd.MarkFlagRequired("capability-configuration")
	iotsitewise_updateGatewayCapabilityConfigurationCmd.MarkFlagRequired("capability-namespace")
	iotsitewise_updateGatewayCapabilityConfigurationCmd.MarkFlagRequired("gateway-id")
	iotsitewiseCmd.AddCommand(iotsitewise_updateGatewayCapabilityConfigurationCmd)
}
