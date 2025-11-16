package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_listEventIntegrationsCmd = &cobra.Command{
	Use:   "list-event-integrations",
	Short: "Returns a paginated list of event integrations in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_listEventIntegrationsCmd).Standalone()

	appintegrations_listEventIntegrationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	appintegrations_listEventIntegrationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	appintegrationsCmd.AddCommand(appintegrations_listEventIntegrationsCmd)
}
