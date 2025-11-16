package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_updateTrustStoreCmd = &cobra.Command{
	Use:   "update-trust-store",
	Short: "Updates the trust store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_updateTrustStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_updateTrustStoreCmd).Standalone()

		workspacesWeb_updateTrustStoreCmd.Flags().String("certificates-to-add", "", "A list of CA certificates to add to the trust store.")
		workspacesWeb_updateTrustStoreCmd.Flags().String("certificates-to-delete", "", "A list of CA certificates to delete from a trust store.")
		workspacesWeb_updateTrustStoreCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		workspacesWeb_updateTrustStoreCmd.Flags().String("trust-store-arn", "", "The ARN of the trust store.")
		workspacesWeb_updateTrustStoreCmd.MarkFlagRequired("trust-store-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_updateTrustStoreCmd)
}
