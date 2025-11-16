package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_createAccountAssociationCmd = &cobra.Command{
	Use:   "create-account-association",
	Short: "Creates a new account association via the destination id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_createAccountAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_createAccountAssociationCmd).Standalone()

		iotManagedIntegrations_createAccountAssociationCmd.Flags().String("client-token", "", "An idempotency token.")
		iotManagedIntegrations_createAccountAssociationCmd.Flags().String("connector-destination-id", "", "The identifier of the connector destination.")
		iotManagedIntegrations_createAccountAssociationCmd.Flags().String("description", "", "A description of the account association request.")
		iotManagedIntegrations_createAccountAssociationCmd.Flags().String("name", "", "The name of the destination for the new account association.")
		iotManagedIntegrations_createAccountAssociationCmd.Flags().String("tags", "", "A set of key/value pairs that are used to manage the account association.")
		iotManagedIntegrations_createAccountAssociationCmd.MarkFlagRequired("connector-destination-id")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_createAccountAssociationCmd)
}
