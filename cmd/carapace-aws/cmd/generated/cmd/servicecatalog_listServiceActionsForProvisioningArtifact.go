package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listServiceActionsForProvisioningArtifactCmd = &cobra.Command{
	Use:   "list-service-actions-for-provisioning-artifact",
	Short: "Returns a paginated list of self-service actions associated with the specified Product ID and Provisioning Artifact ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listServiceActionsForProvisioningArtifactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_listServiceActionsForProvisioningArtifactCmd).Standalone()

		servicecatalog_listServiceActionsForProvisioningArtifactCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_listServiceActionsForProvisioningArtifactCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
		servicecatalog_listServiceActionsForProvisioningArtifactCmd.Flags().String("page-token", "", "The page token for the next set of results.")
		servicecatalog_listServiceActionsForProvisioningArtifactCmd.Flags().String("product-id", "", "The product identifier.")
		servicecatalog_listServiceActionsForProvisioningArtifactCmd.Flags().String("provisioning-artifact-id", "", "The identifier of the provisioning artifact.")
		servicecatalog_listServiceActionsForProvisioningArtifactCmd.MarkFlagRequired("product-id")
		servicecatalog_listServiceActionsForProvisioningArtifactCmd.MarkFlagRequired("provisioning-artifact-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_listServiceActionsForProvisioningArtifactCmd)
}
