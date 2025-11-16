package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listManagedThingAccountAssociationsCmd = &cobra.Command{
	Use:   "list-managed-thing-account-associations",
	Short: "Lists all account associations for a specific managed thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listManagedThingAccountAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_listManagedThingAccountAssociationsCmd).Standalone()

		iotManagedIntegrations_listManagedThingAccountAssociationsCmd.Flags().String("account-association-id", "", "The identifier of the account association to filter results by.")
		iotManagedIntegrations_listManagedThingAccountAssociationsCmd.Flags().String("managed-thing-id", "", "The identifier of the managed thing to list account associations for.")
		iotManagedIntegrations_listManagedThingAccountAssociationsCmd.Flags().String("max-results", "", "The maximum number of account associations to return in a single response.")
		iotManagedIntegrations_listManagedThingAccountAssociationsCmd.Flags().String("next-token", "", "A token used for pagination of results.")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listManagedThingAccountAssociationsCmd)
}
