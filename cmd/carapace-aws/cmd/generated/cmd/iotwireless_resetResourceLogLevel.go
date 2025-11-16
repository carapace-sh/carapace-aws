package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_resetResourceLogLevelCmd = &cobra.Command{
	Use:   "reset-resource-log-level",
	Short: "Removes the log-level override, if any, for a specific resource ID and resource type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_resetResourceLogLevelCmd).Standalone()

	iotwireless_resetResourceLogLevelCmd.Flags().String("resource-identifier", "", "")
	iotwireless_resetResourceLogLevelCmd.Flags().String("resource-type", "", "The type of resource, which can be `WirelessDevice`, `WirelessGateway`, or `FuotaTask`.")
	iotwireless_resetResourceLogLevelCmd.MarkFlagRequired("resource-identifier")
	iotwireless_resetResourceLogLevelCmd.MarkFlagRequired("resource-type")
	iotwirelessCmd.AddCommand(iotwireless_resetResourceLogLevelCmd)
}
