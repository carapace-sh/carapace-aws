package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_getTrustStoreCmd = &cobra.Command{
	Use:   "get-trust-store",
	Short: "Gets the trust store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_getTrustStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_getTrustStoreCmd).Standalone()

		workspacesWeb_getTrustStoreCmd.Flags().String("trust-store-arn", "", "The ARN of the trust store.")
		workspacesWeb_getTrustStoreCmd.MarkFlagRequired("trust-store-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_getTrustStoreCmd)
}
