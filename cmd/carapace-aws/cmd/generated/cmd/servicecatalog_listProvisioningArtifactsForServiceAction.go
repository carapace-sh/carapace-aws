package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listProvisioningArtifactsForServiceActionCmd = &cobra.Command{
	Use:   "list-provisioning-artifacts-for-service-action",
	Short: "Lists all provisioning artifacts (also known as versions) for the specified self-service action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listProvisioningArtifactsForServiceActionCmd).Standalone()

	servicecatalog_listProvisioningArtifactsForServiceActionCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_listProvisioningArtifactsForServiceActionCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
	servicecatalog_listProvisioningArtifactsForServiceActionCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	servicecatalog_listProvisioningArtifactsForServiceActionCmd.Flags().String("service-action-id", "", "The self-service action identifier.")
	servicecatalog_listProvisioningArtifactsForServiceActionCmd.MarkFlagRequired("service-action-id")
	servicecatalogCmd.AddCommand(servicecatalog_listProvisioningArtifactsForServiceActionCmd)
}
