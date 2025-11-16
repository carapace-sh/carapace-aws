package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_scanProvisionedProductsCmd = &cobra.Command{
	Use:   "scan-provisioned-products",
	Short: "Lists the provisioned products that are available (not terminated).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_scanProvisionedProductsCmd).Standalone()

	servicecatalog_scanProvisionedProductsCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_scanProvisionedProductsCmd.Flags().String("access-level-filter", "", "The access level to use to obtain results.")
	servicecatalog_scanProvisionedProductsCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
	servicecatalog_scanProvisionedProductsCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	servicecatalogCmd.AddCommand(servicecatalog_scanProvisionedProductsCmd)
}
