package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_deleteAccountAssociationCmd = &cobra.Command{
	Use:   "delete-account-association",
	Short: "Remove a third-party account association for an end user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_deleteAccountAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_deleteAccountAssociationCmd).Standalone()

		iotManagedIntegrations_deleteAccountAssociationCmd.Flags().String("account-association-id", "", "The unique identifier of the account association to be deleted.")
		iotManagedIntegrations_deleteAccountAssociationCmd.MarkFlagRequired("account-association-id")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_deleteAccountAssociationCmd)
}
