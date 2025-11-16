package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupGateway_getBandwidthRateLimitScheduleCmd = &cobra.Command{
	Use:   "get-bandwidth-rate-limit-schedule",
	Short: "Retrieves the bandwidth rate limit schedule for a specified gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupGateway_getBandwidthRateLimitScheduleCmd).Standalone()

	backupGateway_getBandwidthRateLimitScheduleCmd.Flags().String("gateway-arn", "", "The Amazon Resource Name (ARN) of the gateway.")
	backupGateway_getBandwidthRateLimitScheduleCmd.MarkFlagRequired("gateway-arn")
	backupGatewayCmd.AddCommand(backupGateway_getBandwidthRateLimitScheduleCmd)
}
