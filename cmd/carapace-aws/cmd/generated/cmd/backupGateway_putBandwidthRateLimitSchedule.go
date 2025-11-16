package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_putBandwidthRateLimitScheduleCmd = &cobra.Command{
	Use:   "put-bandwidth-rate-limit-schedule",
	Short: "This action sets the bandwidth rate limit schedule for a specified gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_putBandwidthRateLimitScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backupGateway_putBandwidthRateLimitScheduleCmd).Standalone()

		backupGateway_putBandwidthRateLimitScheduleCmd.Flags().String("bandwidth-rate-limit-intervals", "", "An array containing bandwidth rate limit schedule intervals for a gateway.")
		backupGateway_putBandwidthRateLimitScheduleCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the gateway.")
		backupGateway_putBandwidthRateLimitScheduleCmd.MarkFlagRequired("bandwidth-rate-limit-intervals")
		backupGateway_putBandwidthRateLimitScheduleCmd.MarkFlagRequired("gateway-arn")
	})
	backupGatewayCmd.AddCommand(backupGateway_putBandwidthRateLimitScheduleCmd)
}
