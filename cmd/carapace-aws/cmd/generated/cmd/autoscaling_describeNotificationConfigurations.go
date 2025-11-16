package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeNotificationConfigurationsCmd = &cobra.Command{
	Use:   "describe-notification-configurations",
	Short: "Gets information about the Amazon SNS notifications that are configured for one or more Auto Scaling groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeNotificationConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_describeNotificationConfigurationsCmd).Standalone()

		autoscaling_describeNotificationConfigurationsCmd.Flags().String("auto-scaling-group-names", "", "The name of the Auto Scaling group.")
		autoscaling_describeNotificationConfigurationsCmd.Flags().String("max-records", "", "The maximum number of items to return with this call.")
		autoscaling_describeNotificationConfigurationsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	})
	autoscalingCmd.AddCommand(autoscaling_describeNotificationConfigurationsCmd)
}
