package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationAutoscaling_getPredictiveScalingForecastCmd = &cobra.Command{
	Use:   "get-predictive-scaling-forecast",
	Short: "Retrieves the forecast data for a predictive scaling policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationAutoscaling_getPredictiveScalingForecastCmd).Standalone()

	applicationAutoscaling_getPredictiveScalingForecastCmd.Flags().String("end-time", "", "The exclusive end time of the time range for the forecast data to get.")
	applicationAutoscaling_getPredictiveScalingForecastCmd.Flags().String("policy-name", "", "The name of the policy.")
	applicationAutoscaling_getPredictiveScalingForecastCmd.Flags().String("resource-id", "", "The identifier of the resource.")
	applicationAutoscaling_getPredictiveScalingForecastCmd.Flags().String("scalable-dimension", "", "The scalable dimension.")
	applicationAutoscaling_getPredictiveScalingForecastCmd.Flags().String("service-namespace", "", "The namespace of the Amazon Web Services service that provides the resource.")
	applicationAutoscaling_getPredictiveScalingForecastCmd.Flags().String("start-time", "", "The inclusive start time of the time range for the forecast data to get.")
	applicationAutoscaling_getPredictiveScalingForecastCmd.MarkFlagRequired("end-time")
	applicationAutoscaling_getPredictiveScalingForecastCmd.MarkFlagRequired("policy-name")
	applicationAutoscaling_getPredictiveScalingForecastCmd.MarkFlagRequired("resource-id")
	applicationAutoscaling_getPredictiveScalingForecastCmd.MarkFlagRequired("scalable-dimension")
	applicationAutoscaling_getPredictiveScalingForecastCmd.MarkFlagRequired("service-namespace")
	applicationAutoscaling_getPredictiveScalingForecastCmd.MarkFlagRequired("start-time")
	applicationAutoscalingCmd.AddCommand(applicationAutoscaling_getPredictiveScalingForecastCmd)
}
