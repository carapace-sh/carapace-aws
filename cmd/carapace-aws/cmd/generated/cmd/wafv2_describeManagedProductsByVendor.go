package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_describeManagedProductsByVendorCmd = &cobra.Command{
	Use:   "describe-managed-products-by-vendor",
	Short: "Provides high-level information for the managed rule groups owned by a specific vendor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_describeManagedProductsByVendorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_describeManagedProductsByVendorCmd).Standalone()

		wafv2_describeManagedProductsByVendorCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
		wafv2_describeManagedProductsByVendorCmd.Flags().String("vendor-name", "", "The name of the managed rule group vendor.")
		wafv2_describeManagedProductsByVendorCmd.MarkFlagRequired("scope")
		wafv2_describeManagedProductsByVendorCmd.MarkFlagRequired("vendor-name")
	})
	wafv2Cmd.AddCommand(wafv2_describeManagedProductsByVendorCmd)
}
