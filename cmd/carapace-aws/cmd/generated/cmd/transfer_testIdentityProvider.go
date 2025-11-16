package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_testIdentityProviderCmd = &cobra.Command{
	Use:   "test-identity-provider",
	Short: "If the `IdentityProviderType` of a file transfer protocol-enabled server is `AWS_DIRECTORY_SERVICE` or `API_Gateway`, tests whether your identity provider is set up successfully.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_testIdentityProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_testIdentityProviderCmd).Standalone()

		transfer_testIdentityProviderCmd.Flags().String("server-id", "", "A system-assigned identifier for a specific server.")
		transfer_testIdentityProviderCmd.Flags().String("server-protocol", "", "The type of file transfer protocol to be tested.")
		transfer_testIdentityProviderCmd.Flags().String("source-ip", "", "The source IP address of the account to be tested.")
		transfer_testIdentityProviderCmd.Flags().String("user-name", "", "The name of the account to be tested.")
		transfer_testIdentityProviderCmd.Flags().String("user-password", "", "The password of the account to be tested.")
		transfer_testIdentityProviderCmd.MarkFlagRequired("server-id")
		transfer_testIdentityProviderCmd.MarkFlagRequired("user-name")
	})
	transferCmd.AddCommand(transfer_testIdentityProviderCmd)
}
