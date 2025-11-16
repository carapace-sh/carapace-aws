package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getMonitoringSubscription2020_05_31Cmd = &cobra.Command{
	Use:   "get-monitoring-subscription2020_05_31",
	Short: "Gets information about whether additional CloudWatch metrics are enabled for the specified CloudFront distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getMonitoringSubscription2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_getMonitoringSubscription2020_05_31Cmd).Standalone()

		cloudfront_getMonitoringSubscription2020_05_31Cmd.Flags().String("distribution-id", "", "The ID of the distribution that you are getting metrics information for.")
		cloudfront_getMonitoringSubscription2020_05_31Cmd.MarkFlagRequired("distribution-id")
	})
	cloudfrontCmd.AddCommand(cloudfront_getMonitoringSubscription2020_05_31Cmd)
}
