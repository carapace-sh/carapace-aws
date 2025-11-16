package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_describeMonitorCmd = &cobra.Command{
	Use:   "describe-monitor",
	Short: "Describes a monitor resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_describeMonitorCmd).Standalone()

	forecast_describeMonitorCmd.Flags().String("monitor-arn", "", "The Amazon Resource Name (ARN) of the monitor resource to describe.")
	forecast_describeMonitorCmd.MarkFlagRequired("monitor-arn")
	forecastCmd.AddCommand(forecast_describeMonitorCmd)
}
