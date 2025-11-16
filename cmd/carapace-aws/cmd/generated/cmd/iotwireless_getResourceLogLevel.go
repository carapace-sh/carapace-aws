package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getResourceLogLevelCmd = &cobra.Command{
	Use:   "get-resource-log-level",
	Short: "Fetches the log-level override, if any, for a given resource ID and resource type..",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getResourceLogLevelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_getResourceLogLevelCmd).Standalone()

		iotwireless_getResourceLogLevelCmd.Flags().String("resource-identifier", "", "")
		iotwireless_getResourceLogLevelCmd.Flags().String("resource-type", "", "The type of resource, which can be `WirelessDevice`, `WirelessGateway`, or `FuotaTask`.")
		iotwireless_getResourceLogLevelCmd.MarkFlagRequired("resource-identifier")
		iotwireless_getResourceLogLevelCmd.MarkFlagRequired("resource-type")
	})
	iotwirelessCmd.AddCommand(iotwireless_getResourceLogLevelCmd)
}
