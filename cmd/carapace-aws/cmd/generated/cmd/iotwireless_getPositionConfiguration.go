package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getPositionConfigurationCmd = &cobra.Command{
	Use:   "get-position-configuration",
	Short: "Get position configuration for a given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getPositionConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_getPositionConfigurationCmd).Standalone()

		iotwireless_getPositionConfigurationCmd.Flags().String("resource-identifier", "", "Resource identifier used in a position configuration.")
		iotwireless_getPositionConfigurationCmd.Flags().String("resource-type", "", "Resource type of the resource for which position configuration is retrieved.")
		iotwireless_getPositionConfigurationCmd.MarkFlagRequired("resource-identifier")
		iotwireless_getPositionConfigurationCmd.MarkFlagRequired("resource-type")
	})
	iotwirelessCmd.AddCommand(iotwireless_getPositionConfigurationCmd)
}
