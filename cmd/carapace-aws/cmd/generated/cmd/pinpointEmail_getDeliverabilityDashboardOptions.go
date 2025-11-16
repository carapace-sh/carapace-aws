package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_getDeliverabilityDashboardOptionsCmd = &cobra.Command{
	Use:   "get-deliverability-dashboard-options",
	Short: "Retrieve information about the status of the Deliverability dashboard for your Amazon Pinpoint account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_getDeliverabilityDashboardOptionsCmd).Standalone()

	pinpointEmailCmd.AddCommand(pinpointEmail_getDeliverabilityDashboardOptionsCmd)
}
