package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_startAccountAssociationRefreshCmd = &cobra.Command{
	Use:   "start-account-association-refresh",
	Short: "Initiates a refresh of an existing account association to update its authorization and connection status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_startAccountAssociationRefreshCmd).Standalone()

	iotManagedIntegrations_startAccountAssociationRefreshCmd.Flags().String("account-association-id", "", "The unique identifier of the account association to refresh.")
	iotManagedIntegrations_startAccountAssociationRefreshCmd.MarkFlagRequired("account-association-id")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_startAccountAssociationRefreshCmd)
}
