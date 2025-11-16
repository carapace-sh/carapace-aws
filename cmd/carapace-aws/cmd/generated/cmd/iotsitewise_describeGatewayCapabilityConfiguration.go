package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeGatewayCapabilityConfigurationCmd = &cobra.Command{
	Use:   "describe-gateway-capability-configuration",
	Short: "Each gateway capability defines data sources for a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeGatewayCapabilityConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_describeGatewayCapabilityConfigurationCmd).Standalone()

		iotsitewise_describeGatewayCapabilityConfigurationCmd.Flags().String("capability-namespace", "", "The namespace of the capability configuration.")
		iotsitewise_describeGatewayCapabilityConfigurationCmd.Flags().String("gateway-id", "", "The ID of the gateway that defines the capability configuration.")
		iotsitewise_describeGatewayCapabilityConfigurationCmd.MarkFlagRequired("capability-namespace")
		iotsitewise_describeGatewayCapabilityConfigurationCmd.MarkFlagRequired("gateway-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_describeGatewayCapabilityConfigurationCmd)
}
