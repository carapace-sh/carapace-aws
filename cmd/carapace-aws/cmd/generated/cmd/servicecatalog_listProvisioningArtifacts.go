package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listProvisioningArtifactsCmd = &cobra.Command{
	Use:   "list-provisioning-artifacts",
	Short: "Lists all provisioning artifacts (also known as versions) for the specified product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listProvisioningArtifactsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_listProvisioningArtifactsCmd).Standalone()

		servicecatalog_listProvisioningArtifactsCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_listProvisioningArtifactsCmd.Flags().String("product-id", "", "The product identifier.")
		servicecatalog_listProvisioningArtifactsCmd.MarkFlagRequired("product-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_listProvisioningArtifactsCmd)
}
