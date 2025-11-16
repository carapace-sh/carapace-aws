package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeBandwidthRateLimitScheduleCmd = &cobra.Command{
	Use:   "describe-bandwidth-rate-limit-schedule",
	Short: "Returns information about the bandwidth rate limit schedule of a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeBandwidthRateLimitScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_describeBandwidthRateLimitScheduleCmd).Standalone()

		storagegateway_describeBandwidthRateLimitScheduleCmd.Flags().String("gateway-arn", "", "")
		storagegateway_describeBandwidthRateLimitScheduleCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_describeBandwidthRateLimitScheduleCmd)
}
