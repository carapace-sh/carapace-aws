package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeAvailabilityMonitorTestCmd = &cobra.Command{
	Use:   "describe-availability-monitor-test",
	Short: "Returns information about the most recent high availability monitoring test that was performed on the host in a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeAvailabilityMonitorTestCmd).Standalone()

	storagegateway_describeAvailabilityMonitorTestCmd.Flags().String("gateway-arn", "", "")
	storagegateway_describeAvailabilityMonitorTestCmd.MarkFlagRequired("gateway-arn")
	storagegatewayCmd.AddCommand(storagegateway_describeAvailabilityMonitorTestCmd)
}
