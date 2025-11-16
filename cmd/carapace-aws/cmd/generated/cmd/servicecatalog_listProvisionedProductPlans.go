package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listProvisionedProductPlansCmd = &cobra.Command{
	Use:   "list-provisioned-product-plans",
	Short: "Lists the plans for the specified provisioned product or all plans to which the user has access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listProvisionedProductPlansCmd).Standalone()

	servicecatalog_listProvisionedProductPlansCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_listProvisionedProductPlansCmd.Flags().String("access-level-filter", "", "The access level to use to obtain results.")
	servicecatalog_listProvisionedProductPlansCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
	servicecatalog_listProvisionedProductPlansCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	servicecatalog_listProvisionedProductPlansCmd.Flags().String("provision-product-id", "", "The product identifier.")
	servicecatalogCmd.AddCommand(servicecatalog_listProvisionedProductPlansCmd)
}
