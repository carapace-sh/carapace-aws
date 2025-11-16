package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_listDataIntegrationsCmd = &cobra.Command{
	Use:   "list-data-integrations",
	Short: "Returns a paginated list of DataIntegrations in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_listDataIntegrationsCmd).Standalone()

	appintegrations_listDataIntegrationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	appintegrations_listDataIntegrationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	appintegrationsCmd.AddCommand(appintegrations_listDataIntegrationsCmd)
}
