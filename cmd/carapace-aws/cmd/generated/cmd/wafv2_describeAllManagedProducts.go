package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_describeAllManagedProductsCmd = &cobra.Command{
	Use:   "describe-all-managed-products",
	Short: "Provides high-level information for the Amazon Web Services Managed Rules rule groups and Amazon Web Services Marketplace managed rule groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_describeAllManagedProductsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_describeAllManagedProductsCmd).Standalone()

		wafv2_describeAllManagedProductsCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
		wafv2_describeAllManagedProductsCmd.MarkFlagRequired("scope")
	})
	wafv2Cmd.AddCommand(wafv2_describeAllManagedProductsCmd)
}
