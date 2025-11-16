package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_resetAllResourceLogLevelsCmd = &cobra.Command{
	Use:   "reset-all-resource-log-levels",
	Short: "Removes the log-level overrides for all resources; wireless devices, wireless gateways, and FUOTA tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_resetAllResourceLogLevelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_resetAllResourceLogLevelsCmd).Standalone()

	})
	iotwirelessCmd.AddCommand(iotwireless_resetAllResourceLogLevelsCmd)
}
