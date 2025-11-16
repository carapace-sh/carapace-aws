package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateEventConfigurationsCmd = &cobra.Command{
	Use:   "update-event-configurations",
	Short: "Updates the event configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateEventConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateEventConfigurationsCmd).Standalone()

		iot_updateEventConfigurationsCmd.Flags().String("event-configurations", "", "The new event configuration values.")
	})
	iotCmd.AddCommand(iot_updateEventConfigurationsCmd)
}
