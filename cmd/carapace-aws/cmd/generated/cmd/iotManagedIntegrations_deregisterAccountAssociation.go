package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_deregisterAccountAssociationCmd = &cobra.Command{
	Use:   "deregister-account-association",
	Short: "Deregister an account association from a managed thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_deregisterAccountAssociationCmd).Standalone()

	iotManagedIntegrations_deregisterAccountAssociationCmd.Flags().String("account-association-id", "", "The unique identifier of the account association to be deregistered.")
	iotManagedIntegrations_deregisterAccountAssociationCmd.Flags().String("managed-thing-id", "", "The identifier of the managed thing to be deregistered from the account association.")
	iotManagedIntegrations_deregisterAccountAssociationCmd.MarkFlagRequired("account-association-id")
	iotManagedIntegrations_deregisterAccountAssociationCmd.MarkFlagRequired("managed-thing-id")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_deregisterAccountAssociationCmd)
}
