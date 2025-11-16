package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_registerAccountAssociationCmd = &cobra.Command{
	Use:   "register-account-association",
	Short: "Registers an account association with a managed thing, establishing a connection between a device and a third-party account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_registerAccountAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_registerAccountAssociationCmd).Standalone()

		iotManagedIntegrations_registerAccountAssociationCmd.Flags().String("account-association-id", "", "The identifier of the account association to register with the managed thing.")
		iotManagedIntegrations_registerAccountAssociationCmd.Flags().String("device-discovery-id", "", "The identifier of the device discovery job associated with this registration.")
		iotManagedIntegrations_registerAccountAssociationCmd.Flags().String("managed-thing-id", "", "The identifier of the managed thing to register with the account association.")
		iotManagedIntegrations_registerAccountAssociationCmd.MarkFlagRequired("account-association-id")
		iotManagedIntegrations_registerAccountAssociationCmd.MarkFlagRequired("device-discovery-id")
		iotManagedIntegrations_registerAccountAssociationCmd.MarkFlagRequired("managed-thing-id")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_registerAccountAssociationCmd)
}
