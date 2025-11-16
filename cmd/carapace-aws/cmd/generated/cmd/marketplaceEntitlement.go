package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceEntitlementCmd = &cobra.Command{
	Use:   "marketplace-entitlement",
	Short: "AWS Marketplace Entitlement Service\n\nThis reference provides descriptions of the AWS Marketplace Entitlement Service API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceEntitlementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(marketplaceEntitlementCmd).Standalone()

	})
	rootCmd.AddCommand(marketplaceEntitlementCmd)
}
