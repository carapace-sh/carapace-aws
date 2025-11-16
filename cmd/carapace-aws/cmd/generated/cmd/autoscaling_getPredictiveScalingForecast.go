package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_getPredictiveScalingForecastCmd = &cobra.Command{
	Use:   "get-predictive-scaling-forecast",
	Short: "Retrieves the forecast data for a predictive scaling policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_getPredictiveScalingForecastCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_getPredictiveScalingForecastCmd).Standalone()

		autoscaling_getPredictiveScalingForecastCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_getPredictiveScalingForecastCmd.Flags().String("end-time", "", "The exclusive end time of the time range for the forecast data to get.")
		autoscaling_getPredictiveScalingForecastCmd.Flags().String("policy-name", "", "The name of the policy.")
		autoscaling_getPredictiveScalingForecastCmd.Flags().String("start-time", "", "The inclusive start time of the time range for the forecast data to get.")
		autoscaling_getPredictiveScalingForecastCmd.MarkFlagRequired("auto-scaling-group-name")
		autoscaling_getPredictiveScalingForecastCmd.MarkFlagRequired("end-time")
		autoscaling_getPredictiveScalingForecastCmd.MarkFlagRequired("policy-name")
		autoscaling_getPredictiveScalingForecastCmd.MarkFlagRequired("start-time")
	})
	autoscalingCmd.AddCommand(autoscaling_getPredictiveScalingForecastCmd)
}
