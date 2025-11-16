package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_deleteCredentialLockerCmd = &cobra.Command{
	Use:   "delete-credential-locker",
	Short: "Delete a credential locker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_deleteCredentialLockerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_deleteCredentialLockerCmd).Standalone()

		iotManagedIntegrations_deleteCredentialLockerCmd.Flags().String("identifier", "", "The identifier of the credential locker.")
		iotManagedIntegrations_deleteCredentialLockerCmd.MarkFlagRequired("identifier")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_deleteCredentialLockerCmd)
}
