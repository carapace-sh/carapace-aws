package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_putAlarmCmd = &cobra.Command{
	Use:   "put-alarm",
	Short: "Creates or updates an alarm, and associates it with the specified metric.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_putAlarmCmd).Standalone()

	lightsail_putAlarmCmd.Flags().String("alarm-name", "", "The name for the alarm.")
	lightsail_putAlarmCmd.Flags().String("comparison-operator", "", "The arithmetic operation to use when comparing the specified statistic to the threshold.")
	lightsail_putAlarmCmd.Flags().String("contact-protocols", "", "The contact protocols to use for the alarm, such as `Email`, `SMS` (text messaging), or both.")
	lightsail_putAlarmCmd.Flags().String("datapoints-to-alarm", "", "The number of data points that must be not within the specified threshold to trigger the alarm.")
	lightsail_putAlarmCmd.Flags().String("evaluation-periods", "", "The number of most recent periods over which data is compared to the specified threshold.")
	lightsail_putAlarmCmd.Flags().String("metric-name", "", "The name of the metric to associate with the alarm.")
	lightsail_putAlarmCmd.Flags().String("monitored-resource-name", "", "The name of the Lightsail resource that will be monitored.")
	lightsail_putAlarmCmd.Flags().String("notification-enabled", "", "Indicates whether the alarm is enabled.")
	lightsail_putAlarmCmd.Flags().String("notification-triggers", "", "The alarm states that trigger a notification.")
	lightsail_putAlarmCmd.Flags().String("threshold", "", "The value against which the specified statistic is compared.")
	lightsail_putAlarmCmd.Flags().String("treat-missing-data", "", "Sets how this alarm will handle missing data points.")
	lightsail_putAlarmCmd.MarkFlagRequired("alarm-name")
	lightsail_putAlarmCmd.MarkFlagRequired("comparison-operator")
	lightsail_putAlarmCmd.MarkFlagRequired("evaluation-periods")
	lightsail_putAlarmCmd.MarkFlagRequired("metric-name")
	lightsail_putAlarmCmd.MarkFlagRequired("monitored-resource-name")
	lightsail_putAlarmCmd.MarkFlagRequired("threshold")
	lightsailCmd.AddCommand(lightsail_putAlarmCmd)
}
