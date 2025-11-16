package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_startAvailabilityMonitorTestCmd = &cobra.Command{
	Use:   "start-availability-monitor-test",
	Short: "Start a test that verifies that the specified gateway is configured for High Availability monitoring in your host environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_startAvailabilityMonitorTestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_startAvailabilityMonitorTestCmd).Standalone()

		storagegateway_startAvailabilityMonitorTestCmd.Flags().String("gateway-arn", "", "")
		storagegateway_startAvailabilityMonitorTestCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_startAvailabilityMonitorTestCmd)
}
