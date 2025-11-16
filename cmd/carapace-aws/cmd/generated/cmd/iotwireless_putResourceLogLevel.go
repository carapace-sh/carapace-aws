package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_putResourceLogLevelCmd = &cobra.Command{
	Use:   "put-resource-log-level",
	Short: "Sets the log-level override for a resource ID and resource type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_putResourceLogLevelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_putResourceLogLevelCmd).Standalone()

		iotwireless_putResourceLogLevelCmd.Flags().String("log-level", "", "")
		iotwireless_putResourceLogLevelCmd.Flags().String("resource-identifier", "", "")
		iotwireless_putResourceLogLevelCmd.Flags().String("resource-type", "", "The type of resource, which can be `WirelessDevice`, `WirelessGateway`, or `FuotaTask`.")
		iotwireless_putResourceLogLevelCmd.MarkFlagRequired("log-level")
		iotwireless_putResourceLogLevelCmd.MarkFlagRequired("resource-identifier")
		iotwireless_putResourceLogLevelCmd.MarkFlagRequired("resource-type")
	})
	iotwirelessCmd.AddCommand(iotwireless_putResourceLogLevelCmd)
}
