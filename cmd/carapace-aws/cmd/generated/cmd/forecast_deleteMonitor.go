package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_deleteMonitorCmd = &cobra.Command{
	Use:   "delete-monitor",
	Short: "Deletes a monitor resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_deleteMonitorCmd).Standalone()

	forecast_deleteMonitorCmd.Flags().String("monitor-arn", "", "The Amazon Resource Name (ARN) of the monitor resource to delete.")
	forecast_deleteMonitorCmd.MarkFlagRequired("monitor-arn")
	forecastCmd.AddCommand(forecast_deleteMonitorCmd)
}
