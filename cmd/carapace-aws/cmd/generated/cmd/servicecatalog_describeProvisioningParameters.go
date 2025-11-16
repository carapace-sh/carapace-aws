package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describeProvisioningParametersCmd = &cobra.Command{
	Use:   "describe-provisioning-parameters",
	Short: "Gets information about the configuration required to provision the specified product using the specified provisioning artifact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describeProvisioningParametersCmd).Standalone()

	servicecatalog_describeProvisioningParametersCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_describeProvisioningParametersCmd.Flags().String("path-id", "", "The path identifier of the product.")
	servicecatalog_describeProvisioningParametersCmd.Flags().String("path-name", "", "The name of the path.")
	servicecatalog_describeProvisioningParametersCmd.Flags().String("product-id", "", "The product identifier.")
	servicecatalog_describeProvisioningParametersCmd.Flags().String("product-name", "", "The name of the product.")
	servicecatalog_describeProvisioningParametersCmd.Flags().String("provisioning-artifact-id", "", "The identifier of the provisioning artifact.")
	servicecatalog_describeProvisioningParametersCmd.Flags().String("provisioning-artifact-name", "", "The name of the provisioning artifact.")
	servicecatalogCmd.AddCommand(servicecatalog_describeProvisioningParametersCmd)
}
