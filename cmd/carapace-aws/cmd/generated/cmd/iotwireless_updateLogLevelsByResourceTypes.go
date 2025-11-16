package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_updateLogLevelsByResourceTypesCmd = &cobra.Command{
	Use:   "update-log-levels-by-resource-types",
	Short: "Set default log level, or log levels by resource types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_updateLogLevelsByResourceTypesCmd).Standalone()

	iotwireless_updateLogLevelsByResourceTypesCmd.Flags().String("default-log-level", "", "")
	iotwireless_updateLogLevelsByResourceTypesCmd.Flags().String("fuota-task-log-options", "", "")
	iotwireless_updateLogLevelsByResourceTypesCmd.Flags().String("wireless-device-log-options", "", "")
	iotwireless_updateLogLevelsByResourceTypesCmd.Flags().String("wireless-gateway-log-options", "", "")
	iotwirelessCmd.AddCommand(iotwireless_updateLogLevelsByResourceTypesCmd)
}
