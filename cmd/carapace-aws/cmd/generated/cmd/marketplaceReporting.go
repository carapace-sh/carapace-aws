package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceReportingCmd = &cobra.Command{
	Use:   "marketplace-reporting",
	Short: "The Amazon Web Services Marketplace `GetBuyerDashboard` API enables you to get a procurement insights dashboard programmatically.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceReportingCmd).Standalone()

	rootCmd.AddCommand(marketplaceReportingCmd)
}
