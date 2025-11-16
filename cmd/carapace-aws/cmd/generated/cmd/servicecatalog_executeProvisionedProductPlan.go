package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_executeProvisionedProductPlanCmd = &cobra.Command{
	Use:   "execute-provisioned-product-plan",
	Short: "Provisions or modifies a product based on the resource changes for the specified plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_executeProvisionedProductPlanCmd).Standalone()

	servicecatalog_executeProvisionedProductPlanCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_executeProvisionedProductPlanCmd.Flags().String("idempotency-token", "", "A unique identifier that you provide to ensure idempotency.")
	servicecatalog_executeProvisionedProductPlanCmd.Flags().String("plan-id", "", "The plan identifier.")
	servicecatalog_executeProvisionedProductPlanCmd.MarkFlagRequired("idempotency-token")
	servicecatalog_executeProvisionedProductPlanCmd.MarkFlagRequired("plan-id")
	servicecatalogCmd.AddCommand(servicecatalog_executeProvisionedProductPlanCmd)
}
