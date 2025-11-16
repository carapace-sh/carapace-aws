package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_getProvisionedProductOutputsCmd = &cobra.Command{
	Use:   "get-provisioned-product-outputs",
	Short: "This API takes either a `ProvisonedProductId` or a `ProvisionedProductName`, along with a list of one or more output keys, and responds with the key/value pairs of those outputs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_getProvisionedProductOutputsCmd).Standalone()

	servicecatalog_getProvisionedProductOutputsCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_getProvisionedProductOutputsCmd.Flags().String("output-keys", "", "The list of keys that the API should return with their values.")
	servicecatalog_getProvisionedProductOutputsCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
	servicecatalog_getProvisionedProductOutputsCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	servicecatalog_getProvisionedProductOutputsCmd.Flags().String("provisioned-product-id", "", "The identifier of the provisioned product that you want the outputs from.")
	servicecatalog_getProvisionedProductOutputsCmd.Flags().String("provisioned-product-name", "", "The name of the provisioned product that you want the outputs from.")
	servicecatalogCmd.AddCommand(servicecatalog_getProvisionedProductOutputsCmd)
}
