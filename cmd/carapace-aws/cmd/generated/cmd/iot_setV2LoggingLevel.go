package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_setV2LoggingLevelCmd = &cobra.Command{
	Use:   "set-v2-logging-level",
	Short: "Sets the logging level.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_setV2LoggingLevelCmd).Standalone()

	iot_setV2LoggingLevelCmd.Flags().String("log-level", "", "The log level.")
	iot_setV2LoggingLevelCmd.Flags().String("log-target", "", "The log target.")
	iot_setV2LoggingLevelCmd.MarkFlagRequired("log-level")
	iot_setV2LoggingLevelCmd.MarkFlagRequired("log-target")
	iotCmd.AddCommand(iot_setV2LoggingLevelCmd)
}
