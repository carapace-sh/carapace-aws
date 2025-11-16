package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getCredentialLockerCmd = &cobra.Command{
	Use:   "get-credential-locker",
	Short: "Get information on an existing credential locker",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getCredentialLockerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_getCredentialLockerCmd).Standalone()

		iotManagedIntegrations_getCredentialLockerCmd.Flags().String("identifier", "", "The identifier of the credential locker.")
		iotManagedIntegrations_getCredentialLockerCmd.MarkFlagRequired("identifier")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getCredentialLockerCmd)
}
