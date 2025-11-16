package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_terminateProvisionedProductCmd = &cobra.Command{
	Use:   "terminate-provisioned-product",
	Short: "Terminates the specified provisioned product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_terminateProvisionedProductCmd).Standalone()

	servicecatalog_terminateProvisionedProductCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_terminateProvisionedProductCmd.Flags().String("ignore-errors", "", "If set to true, Service Catalog stops managing the specified provisioned product even if it cannot delete the underlying resources.")
	servicecatalog_terminateProvisionedProductCmd.Flags().String("provisioned-product-id", "", "The identifier of the provisioned product.")
	servicecatalog_terminateProvisionedProductCmd.Flags().String("provisioned-product-name", "", "The name of the provisioned product.")
	servicecatalog_terminateProvisionedProductCmd.Flags().String("retain-physical-resources", "", "When this boolean parameter is set to true, the `TerminateProvisionedProduct` API deletes the Service Catalog provisioned product.")
	servicecatalog_terminateProvisionedProductCmd.Flags().String("terminate-token", "", "An idempotency token that uniquely identifies the termination request.")
	servicecatalog_terminateProvisionedProductCmd.MarkFlagRequired("terminate-token")
	servicecatalogCmd.AddCommand(servicecatalog_terminateProvisionedProductCmd)
}
