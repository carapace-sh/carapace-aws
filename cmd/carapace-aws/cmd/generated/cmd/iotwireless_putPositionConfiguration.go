package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_putPositionConfigurationCmd = &cobra.Command{
	Use:   "put-position-configuration",
	Short: "Put position configuration for a given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_putPositionConfigurationCmd).Standalone()

	iotwireless_putPositionConfigurationCmd.Flags().String("destination", "", "The position data destination that describes the AWS IoT rule that processes the device's position data for use by AWS IoT Core for LoRaWAN.")
	iotwireless_putPositionConfigurationCmd.Flags().String("resource-identifier", "", "Resource identifier used to update the position configuration.")
	iotwireless_putPositionConfigurationCmd.Flags().String("resource-type", "", "Resource type of the resource for which you want to update the position configuration.")
	iotwireless_putPositionConfigurationCmd.Flags().String("solvers", "", "The positioning solvers used to update the position configuration of the resource.")
	iotwireless_putPositionConfigurationCmd.MarkFlagRequired("resource-identifier")
	iotwireless_putPositionConfigurationCmd.MarkFlagRequired("resource-type")
	iotwirelessCmd.AddCommand(iotwireless_putPositionConfigurationCmd)
}
