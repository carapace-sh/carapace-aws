package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listAccountAssociationsCmd = &cobra.Command{
	Use:   "list-account-associations",
	Short: "Lists all account associations, with optional filtering by connector destination ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listAccountAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_listAccountAssociationsCmd).Standalone()

		iotManagedIntegrations_listAccountAssociationsCmd.Flags().String("connector-destination-id", "", "The identifier of the connector destination to filter account associations by.")
		iotManagedIntegrations_listAccountAssociationsCmd.Flags().String("max-results", "", "The maximum number of account associations to return in a single response.")
		iotManagedIntegrations_listAccountAssociationsCmd.Flags().String("next-token", "", "A token used for pagination of results.")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listAccountAssociationsCmd)
}
