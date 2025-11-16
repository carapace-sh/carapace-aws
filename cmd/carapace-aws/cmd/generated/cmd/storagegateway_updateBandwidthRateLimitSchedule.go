package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_updateBandwidthRateLimitScheduleCmd = &cobra.Command{
	Use:   "update-bandwidth-rate-limit-schedule",
	Short: "Updates the bandwidth rate limit schedule for a specified gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_updateBandwidthRateLimitScheduleCmd).Standalone()

	storagegateway_updateBandwidthRateLimitScheduleCmd.Flags().String("bandwidth-rate-limit-intervals", "", "An array containing bandwidth rate limit schedule intervals for a gateway.")
	storagegateway_updateBandwidthRateLimitScheduleCmd.Flags().String("gateway-arn", "", "")
	storagegateway_updateBandwidthRateLimitScheduleCmd.MarkFlagRequired("bandwidth-rate-limit-intervals")
	storagegateway_updateBandwidthRateLimitScheduleCmd.MarkFlagRequired("gateway-arn")
	storagegatewayCmd.AddCommand(storagegateway_updateBandwidthRateLimitScheduleCmd)
}
