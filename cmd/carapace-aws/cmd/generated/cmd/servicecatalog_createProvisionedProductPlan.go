package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_createProvisionedProductPlanCmd = &cobra.Command{
	Use:   "create-provisioned-product-plan",
	Short: "Creates a plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_createProvisionedProductPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_createProvisionedProductPlanCmd).Standalone()

		servicecatalog_createProvisionedProductPlanCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_createProvisionedProductPlanCmd.Flags().String("idempotency-token", "", "A unique identifier that you provide to ensure idempotency.")
		servicecatalog_createProvisionedProductPlanCmd.Flags().String("notification-arns", "", "Passed to CloudFormation.")
		servicecatalog_createProvisionedProductPlanCmd.Flags().String("path-id", "", "The path identifier of the product.")
		servicecatalog_createProvisionedProductPlanCmd.Flags().String("plan-name", "", "The name of the plan.")
		servicecatalog_createProvisionedProductPlanCmd.Flags().String("plan-type", "", "The plan type.")
		servicecatalog_createProvisionedProductPlanCmd.Flags().String("product-id", "", "The product identifier.")
		servicecatalog_createProvisionedProductPlanCmd.Flags().String("provisioned-product-name", "", "A user-friendly name for the provisioned product.")
		servicecatalog_createProvisionedProductPlanCmd.Flags().String("provisioning-artifact-id", "", "The identifier of the provisioning artifact.")
		servicecatalog_createProvisionedProductPlanCmd.Flags().String("provisioning-parameters", "", "Parameters specified by the administrator that are required for provisioning the product.")
		servicecatalog_createProvisionedProductPlanCmd.Flags().String("tags", "", "One or more tags.")
		servicecatalog_createProvisionedProductPlanCmd.MarkFlagRequired("idempotency-token")
		servicecatalog_createProvisionedProductPlanCmd.MarkFlagRequired("plan-name")
		servicecatalog_createProvisionedProductPlanCmd.MarkFlagRequired("plan-type")
		servicecatalog_createProvisionedProductPlanCmd.MarkFlagRequired("product-id")
		servicecatalog_createProvisionedProductPlanCmd.MarkFlagRequired("provisioned-product-name")
		servicecatalog_createProvisionedProductPlanCmd.MarkFlagRequired("provisioning-artifact-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_createProvisionedProductPlanCmd)
}
