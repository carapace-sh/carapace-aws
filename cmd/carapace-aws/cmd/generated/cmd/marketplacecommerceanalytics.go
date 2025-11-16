package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplacecommerceanalyticsCmd = &cobra.Command{
	Use:   "marketplacecommerceanalytics",
	Short: "Provides AWS Marketplace business intelligence data on-demand.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplacecommerceanalyticsCmd).Standalone()

	rootCmd.AddCommand(marketplacecommerceanalyticsCmd)
}
