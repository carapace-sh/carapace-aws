package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceReporting_getBuyerDashboardCmd = &cobra.Command{
	Use:   "get-buyer-dashboard",
	Short: "Generates an embedding URL for an Amazon QuickSight dashboard for an anonymous user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceReporting_getBuyerDashboardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(marketplaceReporting_getBuyerDashboardCmd).Standalone()

		marketplaceReporting_getBuyerDashboardCmd.Flags().String("dashboard-identifier", "", "The ARN of the requested dashboard.")
		marketplaceReporting_getBuyerDashboardCmd.Flags().String("embedding-domains", "", "Fully qualified domains that you add to the allow list for access to the generated URL that is then embedded.")
		marketplaceReporting_getBuyerDashboardCmd.MarkFlagRequired("dashboard-identifier")
		marketplaceReporting_getBuyerDashboardCmd.MarkFlagRequired("embedding-domains")
	})
	marketplaceReportingCmd.AddCommand(marketplaceReporting_getBuyerDashboardCmd)
}
