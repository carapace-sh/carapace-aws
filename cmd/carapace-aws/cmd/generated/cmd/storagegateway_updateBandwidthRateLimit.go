package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_updateBandwidthRateLimitCmd = &cobra.Command{
	Use:   "update-bandwidth-rate-limit",
	Short: "Updates the bandwidth rate limits of a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_updateBandwidthRateLimitCmd).Standalone()

	storagegateway_updateBandwidthRateLimitCmd.Flags().String("average-download-rate-limit-in-bits-per-sec", "", "The average download bandwidth rate limit in bits per second.")
	storagegateway_updateBandwidthRateLimitCmd.Flags().String("average-upload-rate-limit-in-bits-per-sec", "", "The average upload bandwidth rate limit in bits per second.")
	storagegateway_updateBandwidthRateLimitCmd.Flags().String("gateway-arn", "", "")
	storagegateway_updateBandwidthRateLimitCmd.MarkFlagRequired("gateway-arn")
	storagegatewayCmd.AddCommand(storagegateway_updateBandwidthRateLimitCmd)
}
