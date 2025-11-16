package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_listDataIntegrationAssociationsCmd = &cobra.Command{
	Use:   "list-data-integration-associations",
	Short: "Returns a paginated list of DataIntegration associations in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_listDataIntegrationAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appintegrations_listDataIntegrationAssociationsCmd).Standalone()

		appintegrations_listDataIntegrationAssociationsCmd.Flags().String("data-integration-identifier", "", "A unique identifier for the DataIntegration.")
		appintegrations_listDataIntegrationAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		appintegrations_listDataIntegrationAssociationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		appintegrations_listDataIntegrationAssociationsCmd.MarkFlagRequired("data-integration-identifier")
	})
	appintegrationsCmd.AddCommand(appintegrations_listDataIntegrationAssociationsCmd)
}
