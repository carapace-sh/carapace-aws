package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_deleteProvisionedProductPlanCmd = &cobra.Command{
	Use:   "delete-provisioned-product-plan",
	Short: "Deletes the specified plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_deleteProvisionedProductPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_deleteProvisionedProductPlanCmd).Standalone()

		servicecatalog_deleteProvisionedProductPlanCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_deleteProvisionedProductPlanCmd.Flags().String("ignore-errors", "", "If set to true, Service Catalog stops managing the specified provisioned product even if it cannot delete the underlying resources.")
		servicecatalog_deleteProvisionedProductPlanCmd.Flags().String("plan-id", "", "The plan identifier.")
		servicecatalog_deleteProvisionedProductPlanCmd.MarkFlagRequired("plan-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_deleteProvisionedProductPlanCmd)
}
