package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listConnectorDestinationsCmd = &cobra.Command{
	Use:   "list-connector-destinations",
	Short: "Lists all connector destinations, with optional filtering by cloud connector ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listConnectorDestinationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_listConnectorDestinationsCmd).Standalone()

		iotManagedIntegrations_listConnectorDestinationsCmd.Flags().String("cloud-connector-id", "", "The identifier of the cloud connector to filter connector destinations by.")
		iotManagedIntegrations_listConnectorDestinationsCmd.Flags().String("max-results", "", "The maximum number of connector destinations to return in a single response.")
		iotManagedIntegrations_listConnectorDestinationsCmd.Flags().String("next-token", "", "A token used for pagination of results.")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listConnectorDestinationsCmd)
}
