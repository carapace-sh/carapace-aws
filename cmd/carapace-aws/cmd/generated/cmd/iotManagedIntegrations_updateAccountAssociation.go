package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_updateAccountAssociationCmd = &cobra.Command{
	Use:   "update-account-association",
	Short: "Updates the properties of an existing account association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_updateAccountAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_updateAccountAssociationCmd).Standalone()

		iotManagedIntegrations_updateAccountAssociationCmd.Flags().String("account-association-id", "", "The unique identifier of the account association to update.")
		iotManagedIntegrations_updateAccountAssociationCmd.Flags().String("description", "", "The new description to assign to the account association.")
		iotManagedIntegrations_updateAccountAssociationCmd.Flags().String("name", "", "The new name to assign to the account association.")
		iotManagedIntegrations_updateAccountAssociationCmd.MarkFlagRequired("account-association-id")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_updateAccountAssociationCmd)
}
