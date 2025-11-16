package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getLogLevelsByResourceTypesCmd = &cobra.Command{
	Use:   "get-log-levels-by-resource-types",
	Short: "Returns current default log levels or log levels by resource types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getLogLevelsByResourceTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_getLogLevelsByResourceTypesCmd).Standalone()

	})
	iotwirelessCmd.AddCommand(iotwireless_getLogLevelsByResourceTypesCmd)
}
