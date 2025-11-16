package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceEntitlement_getEntitlementsCmd = &cobra.Command{
	Use:   "get-entitlements",
	Short: "GetEntitlements retrieves entitlement values for a given product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceEntitlement_getEntitlementsCmd).Standalone()

	marketplaceEntitlement_getEntitlementsCmd.Flags().String("filter", "", "Filter is used to return entitlements for a specific customer or for a specific dimension.")
	marketplaceEntitlement_getEntitlementsCmd.Flags().String("max-results", "", "The maximum number of items to retrieve from the GetEntitlements operation.")
	marketplaceEntitlement_getEntitlementsCmd.Flags().String("next-token", "", "For paginated calls to GetEntitlements, pass the NextToken from the previous GetEntitlementsResult.")
	marketplaceEntitlement_getEntitlementsCmd.Flags().String("product-code", "", "Product code is used to uniquely identify a product in AWS Marketplace.")
	marketplaceEntitlement_getEntitlementsCmd.MarkFlagRequired("product-code")
	marketplaceEntitlementCmd.AddCommand(marketplaceEntitlement_getEntitlementsCmd)
}
