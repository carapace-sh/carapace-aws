package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describeProvisionedProductPlanCmd = &cobra.Command{
	Use:   "describe-provisioned-product-plan",
	Short: "Gets information about the resource changes for the specified plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describeProvisionedProductPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_describeProvisionedProductPlanCmd).Standalone()

		servicecatalog_describeProvisionedProductPlanCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_describeProvisionedProductPlanCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
		servicecatalog_describeProvisionedProductPlanCmd.Flags().String("page-token", "", "The page token for the next set of results.")
		servicecatalog_describeProvisionedProductPlanCmd.Flags().String("plan-id", "", "The plan identifier.")
		servicecatalog_describeProvisionedProductPlanCmd.MarkFlagRequired("plan-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_describeProvisionedProductPlanCmd)
}
