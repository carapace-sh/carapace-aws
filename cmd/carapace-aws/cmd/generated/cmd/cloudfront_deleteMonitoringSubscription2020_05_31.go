package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteMonitoringSubscription2020_05_31Cmd = &cobra.Command{
	Use:   "delete-monitoring-subscription2020_05_31",
	Short: "Disables additional CloudWatch metrics for the specified CloudFront distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteMonitoringSubscription2020_05_31Cmd).Standalone()

	cloudfront_deleteMonitoringSubscription2020_05_31Cmd.Flags().String("distribution-id", "", "The ID of the distribution that you are disabling metrics for.")
	cloudfront_deleteMonitoringSubscription2020_05_31Cmd.MarkFlagRequired("distribution-id")
	cloudfrontCmd.AddCommand(cloudfront_deleteMonitoringSubscription2020_05_31Cmd)
}
