package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_associateTrustStoreCmd = &cobra.Command{
	Use:   "associate-trust-store",
	Short: "Associates a trust store with a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_associateTrustStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_associateTrustStoreCmd).Standalone()

		workspacesWeb_associateTrustStoreCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
		workspacesWeb_associateTrustStoreCmd.Flags().String("trust-store-arn", "", "The ARN of the trust store.")
		workspacesWeb_associateTrustStoreCmd.MarkFlagRequired("portal-arn")
		workspacesWeb_associateTrustStoreCmd.MarkFlagRequired("trust-store-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_associateTrustStoreCmd)
}
