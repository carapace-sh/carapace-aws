package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_startInputDeviceMaintenanceWindowCmd = &cobra.Command{
	Use:   "start-input-device-maintenance-window",
	Short: "Start a maintenance window for the specified input device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_startInputDeviceMaintenanceWindowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_startInputDeviceMaintenanceWindowCmd).Standalone()

		medialive_startInputDeviceMaintenanceWindowCmd.Flags().String("input-device-id", "", "The unique ID of the input device to start a maintenance window for.")
		medialive_startInputDeviceMaintenanceWindowCmd.MarkFlagRequired("input-device-id")
	})
	medialiveCmd.AddCommand(medialive_startInputDeviceMaintenanceWindowCmd)
}
