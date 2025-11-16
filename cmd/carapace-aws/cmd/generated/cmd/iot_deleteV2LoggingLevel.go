package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteV2LoggingLevelCmd = &cobra.Command{
	Use:   "delete-v2-logging-level",
	Short: "Deletes a logging level.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteV2LoggingLevelCmd).Standalone()

	iot_deleteV2LoggingLevelCmd.Flags().String("target-name", "", "The name of the resource for which you are configuring logging.")
	iot_deleteV2LoggingLevelCmd.Flags().String("target-type", "", "The type of resource for which you are configuring logging.")
	iot_deleteV2LoggingLevelCmd.MarkFlagRequired("target-name")
	iot_deleteV2LoggingLevelCmd.MarkFlagRequired("target-type")
	iotCmd.AddCommand(iot_deleteV2LoggingLevelCmd)
}
