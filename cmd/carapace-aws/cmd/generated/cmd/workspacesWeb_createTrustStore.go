package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_createTrustStoreCmd = &cobra.Command{
	Use:   "create-trust-store",
	Short: "Creates a trust store that can be associated with a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_createTrustStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_createTrustStoreCmd).Standalone()

		workspacesWeb_createTrustStoreCmd.Flags().String("certificate-list", "", "A list of CA certificates to be added to the trust store.")
		workspacesWeb_createTrustStoreCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		workspacesWeb_createTrustStoreCmd.Flags().String("tags", "", "The tags to add to the trust store.")
		workspacesWeb_createTrustStoreCmd.MarkFlagRequired("certificate-list")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_createTrustStoreCmd)
}
