package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getAccountAssociationCmd = &cobra.Command{
	Use:   "get-account-association",
	Short: "Get an account association for an Amazon Web Services account linked to a customer-managed destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getAccountAssociationCmd).Standalone()

	iotManagedIntegrations_getAccountAssociationCmd.Flags().String("account-association-id", "", "The unique identifier of the account association to retrieve.")
	iotManagedIntegrations_getAccountAssociationCmd.MarkFlagRequired("account-association-id")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getAccountAssociationCmd)
}
