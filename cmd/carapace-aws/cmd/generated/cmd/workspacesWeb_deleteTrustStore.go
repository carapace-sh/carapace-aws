package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_deleteTrustStoreCmd = &cobra.Command{
	Use:   "delete-trust-store",
	Short: "Deletes the trust store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_deleteTrustStoreCmd).Standalone()

	workspacesWeb_deleteTrustStoreCmd.Flags().String("trust-store-arn", "", "The ARN of the trust store.")
	workspacesWeb_deleteTrustStoreCmd.MarkFlagRequired("trust-store-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_deleteTrustStoreCmd)
}
