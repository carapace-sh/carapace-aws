package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscalingPlans_getScalingPlanResourceForecastDataCmd = &cobra.Command{
	Use:   "get-scaling-plan-resource-forecast-data",
	Short: "Retrieves the forecast data for a scalable resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscalingPlans_getScalingPlanResourceForecastDataCmd).Standalone()

	autoscalingPlans_getScalingPlanResourceForecastDataCmd.Flags().String("end-time", "", "The exclusive end time of the time range for the forecast data to get.")
	autoscalingPlans_getScalingPlanResourceForecastDataCmd.Flags().String("forecast-data-type", "", "The type of forecast data to get.")
	autoscalingPlans_getScalingPlanResourceForecastDataCmd.Flags().String("resource-id", "", "The ID of the resource.")
	autoscalingPlans_getScalingPlanResourceForecastDataCmd.Flags().String("scalable-dimension", "", "The scalable dimension for the resource.")
	autoscalingPlans_getScalingPlanResourceForecastDataCmd.Flags().String("scaling-plan-name", "", "The name of the scaling plan.")
	autoscalingPlans_getScalingPlanResourceForecastDataCmd.Flags().String("scaling-plan-version", "", "The version number of the scaling plan.")
	autoscalingPlans_getScalingPlanResourceForecastDataCmd.Flags().String("service-namespace", "", "The namespace of the AWS service.")
	autoscalingPlans_getScalingPlanResourceForecastDataCmd.Flags().String("start-time", "", "The inclusive start time of the time range for the forecast data to get.")
	autoscalingPlans_getScalingPlanResourceForecastDataCmd.MarkFlagRequired("end-time")
	autoscalingPlans_getScalingPlanResourceForecastDataCmd.MarkFlagRequired("forecast-data-type")
	autoscalingPlans_getScalingPlanResourceForecastDataCmd.MarkFlagRequired("resource-id")
	autoscalingPlans_getScalingPlanResourceForecastDataCmd.MarkFlagRequired("scalable-dimension")
	autoscalingPlans_getScalingPlanResourceForecastDataCmd.MarkFlagRequired("scaling-plan-name")
	autoscalingPlans_getScalingPlanResourceForecastDataCmd.MarkFlagRequired("scaling-plan-version")
	autoscalingPlans_getScalingPlanResourceForecastDataCmd.MarkFlagRequired("service-namespace")
	autoscalingPlans_getScalingPlanResourceForecastDataCmd.MarkFlagRequired("start-time")
	autoscalingPlansCmd.AddCommand(autoscalingPlans_getScalingPlanResourceForecastDataCmd)
}
