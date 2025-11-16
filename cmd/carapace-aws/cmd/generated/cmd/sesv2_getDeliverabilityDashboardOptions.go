package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getDeliverabilityDashboardOptionsCmd = &cobra.Command{
	Use:   "get-deliverability-dashboard-options",
	Short: "Retrieve information about the status of the Deliverability dashboard for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getDeliverabilityDashboardOptionsCmd).Standalone()

	sesv2Cmd.AddCommand(sesv2_getDeliverabilityDashboardOptionsCmd)
}
