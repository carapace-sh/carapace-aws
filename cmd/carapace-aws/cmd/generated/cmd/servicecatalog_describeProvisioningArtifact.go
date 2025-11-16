package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describeProvisioningArtifactCmd = &cobra.Command{
	Use:   "describe-provisioning-artifact",
	Short: "Gets information about the specified provisioning artifact (also known as a version) for the specified product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describeProvisioningArtifactCmd).Standalone()

	servicecatalog_describeProvisioningArtifactCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_describeProvisioningArtifactCmd.Flags().Bool("include-provisioning-artifact-parameters", false, "Indicates if the API call response does or does not include additional details about the provisioning parameters.")
	servicecatalog_describeProvisioningArtifactCmd.Flags().Bool("no-include-provisioning-artifact-parameters", false, "Indicates if the API call response does or does not include additional details about the provisioning parameters.")
	servicecatalog_describeProvisioningArtifactCmd.Flags().String("product-id", "", "The product identifier.")
	servicecatalog_describeProvisioningArtifactCmd.Flags().String("product-name", "", "The product name.")
	servicecatalog_describeProvisioningArtifactCmd.Flags().String("provisioning-artifact-id", "", "The identifier of the provisioning artifact.")
	servicecatalog_describeProvisioningArtifactCmd.Flags().String("provisioning-artifact-name", "", "The provisioning artifact name.")
	servicecatalog_describeProvisioningArtifactCmd.Flags().String("verbose", "", "Indicates whether a verbose level of detail is enabled.")
	servicecatalog_describeProvisioningArtifactCmd.Flag("no-include-provisioning-artifact-parameters").Hidden = true
	servicecatalogCmd.AddCommand(servicecatalog_describeProvisioningArtifactCmd)
}
