package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createMonitoringSubscription2020_05_31Cmd = &cobra.Command{
	Use:   "create-monitoring-subscription2020_05_31",
	Short: "Enables or disables additional Amazon CloudWatch metrics for the specified CloudFront distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createMonitoringSubscription2020_05_31Cmd).Standalone()

	cloudfront_createMonitoringSubscription2020_05_31Cmd.Flags().String("distribution-id", "", "The ID of the distribution that you are enabling metrics for.")
	cloudfront_createMonitoringSubscription2020_05_31Cmd.Flags().String("monitoring-subscription", "", "A monitoring subscription.")
	cloudfront_createMonitoringSubscription2020_05_31Cmd.MarkFlagRequired("distribution-id")
	cloudfront_createMonitoringSubscription2020_05_31Cmd.MarkFlagRequired("monitoring-subscription")
	cloudfrontCmd.AddCommand(cloudfront_createMonitoringSubscription2020_05_31Cmd)
}
