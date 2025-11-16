package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_createCredentialLockerCmd = &cobra.Command{
	Use:   "create-credential-locker",
	Short: "Create a credential locker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_createCredentialLockerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_createCredentialLockerCmd).Standalone()

		iotManagedIntegrations_createCredentialLockerCmd.Flags().String("client-token", "", "An idempotency token.")
		iotManagedIntegrations_createCredentialLockerCmd.Flags().String("name", "", "The name of the credential locker.")
		iotManagedIntegrations_createCredentialLockerCmd.Flags().String("tags", "", "A set of key/value pairs that are used to manage the credential locker.")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_createCredentialLockerCmd)
}
