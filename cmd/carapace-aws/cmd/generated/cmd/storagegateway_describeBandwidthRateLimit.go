package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeBandwidthRateLimitCmd = &cobra.Command{
	Use:   "describe-bandwidth-rate-limit",
	Short: "Returns the bandwidth rate limits of a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeBandwidthRateLimitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_describeBandwidthRateLimitCmd).Standalone()

		storagegateway_describeBandwidthRateLimitCmd.Flags().String("gateway-arn", "", "")
		storagegateway_describeBandwidthRateLimitCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_describeBandwidthRateLimitCmd)
}
