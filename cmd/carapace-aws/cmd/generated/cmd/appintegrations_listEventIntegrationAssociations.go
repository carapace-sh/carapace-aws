package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_listEventIntegrationAssociationsCmd = &cobra.Command{
	Use:   "list-event-integration-associations",
	Short: "Returns a paginated list of event integration associations in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_listEventIntegrationAssociationsCmd).Standalone()

	appintegrations_listEventIntegrationAssociationsCmd.Flags().String("event-integration-name", "", "The name of the event integration.")
	appintegrations_listEventIntegrationAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	appintegrations_listEventIntegrationAssociationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	appintegrations_listEventIntegrationAssociationsCmd.MarkFlagRequired("event-integration-name")
	appintegrationsCmd.AddCommand(appintegrations_listEventIntegrationAssociationsCmd)
}
