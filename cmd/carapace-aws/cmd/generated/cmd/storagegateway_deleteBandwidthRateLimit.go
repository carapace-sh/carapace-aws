package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_deleteBandwidthRateLimitCmd = &cobra.Command{
	Use:   "delete-bandwidth-rate-limit",
	Short: "Deletes the bandwidth rate limits of a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_deleteBandwidthRateLimitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_deleteBandwidthRateLimitCmd).Standalone()

		storagegateway_deleteBandwidthRateLimitCmd.Flags().String("bandwidth-type", "", "One of the BandwidthType values that indicates the gateway bandwidth rate limit to delete.")
		storagegateway_deleteBandwidthRateLimitCmd.Flags().String("gateway-arn", "", "")
		storagegateway_deleteBandwidthRateLimitCmd.MarkFlagRequired("bandwidth-type")
		storagegateway_deleteBandwidthRateLimitCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_deleteBandwidthRateLimitCmd)
}
