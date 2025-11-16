package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_provisionProductCmd = &cobra.Command{
	Use:   "provision-product",
	Short: "Provisions the specified product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_provisionProductCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_provisionProductCmd).Standalone()

		servicecatalog_provisionProductCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_provisionProductCmd.Flags().String("notification-arns", "", "Passed to CloudFormation.")
		servicecatalog_provisionProductCmd.Flags().String("path-id", "", "The path identifier of the product.")
		servicecatalog_provisionProductCmd.Flags().String("path-name", "", "The name of the path.")
		servicecatalog_provisionProductCmd.Flags().String("product-id", "", "The product identifier.")
		servicecatalog_provisionProductCmd.Flags().String("product-name", "", "The name of the product.")
		servicecatalog_provisionProductCmd.Flags().String("provision-token", "", "An idempotency token that uniquely identifies the provisioning request.")
		servicecatalog_provisionProductCmd.Flags().String("provisioned-product-name", "", "A user-friendly name for the provisioned product.")
		servicecatalog_provisionProductCmd.Flags().String("provisioning-artifact-id", "", "The identifier of the provisioning artifact.")
		servicecatalog_provisionProductCmd.Flags().String("provisioning-artifact-name", "", "The name of the provisioning artifact.")
		servicecatalog_provisionProductCmd.Flags().String("provisioning-parameters", "", "Parameters specified by the administrator that are required for provisioning the product.")
		servicecatalog_provisionProductCmd.Flags().String("provisioning-preferences", "", "An object that contains information about the provisioning preferences for a stack set.")
		servicecatalog_provisionProductCmd.Flags().String("tags", "", "One or more tags.")
		servicecatalog_provisionProductCmd.MarkFlagRequired("provision-token")
		servicecatalog_provisionProductCmd.MarkFlagRequired("provisioned-product-name")
	})
	servicecatalogCmd.AddCommand(servicecatalog_provisionProductCmd)
}
