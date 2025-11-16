package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_executeProvisionedProductServiceActionCmd = &cobra.Command{
	Use:   "execute-provisioned-product-service-action",
	Short: "Executes a self-service action against a provisioned product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_executeProvisionedProductServiceActionCmd).Standalone()

	servicecatalog_executeProvisionedProductServiceActionCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_executeProvisionedProductServiceActionCmd.Flags().String("execute-token", "", "An idempotency token that uniquely identifies the execute request.")
	servicecatalog_executeProvisionedProductServiceActionCmd.Flags().String("parameters", "", "A map of all self-service action parameters and their values.")
	servicecatalog_executeProvisionedProductServiceActionCmd.Flags().String("provisioned-product-id", "", "The identifier of the provisioned product.")
	servicecatalog_executeProvisionedProductServiceActionCmd.Flags().String("service-action-id", "", "The self-service action identifier.")
	servicecatalog_executeProvisionedProductServiceActionCmd.MarkFlagRequired("execute-token")
	servicecatalog_executeProvisionedProductServiceActionCmd.MarkFlagRequired("provisioned-product-id")
	servicecatalog_executeProvisionedProductServiceActionCmd.MarkFlagRequired("service-action-id")
	servicecatalogCmd.AddCommand(servicecatalog_executeProvisionedProductServiceActionCmd)
}
