package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listStackInstancesForProvisionedProductCmd = &cobra.Command{
	Use:   "list-stack-instances-for-provisioned-product",
	Short: "Returns summary information about stack instances that are associated with the specified `CFN_STACKSET` type provisioned product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listStackInstancesForProvisionedProductCmd).Standalone()

	servicecatalog_listStackInstancesForProvisionedProductCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_listStackInstancesForProvisionedProductCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
	servicecatalog_listStackInstancesForProvisionedProductCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	servicecatalog_listStackInstancesForProvisionedProductCmd.Flags().String("provisioned-product-id", "", "The identifier of the provisioned product.")
	servicecatalog_listStackInstancesForProvisionedProductCmd.MarkFlagRequired("provisioned-product-id")
	servicecatalogCmd.AddCommand(servicecatalog_listStackInstancesForProvisionedProductCmd)
}
