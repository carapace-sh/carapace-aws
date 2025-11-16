package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_enableDelegatedAdminAccountCmd = &cobra.Command{
	Use:   "enable-delegated-admin-account",
	Short: "Enables the Amazon Inspector delegated administrator for your Organizations organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_enableDelegatedAdminAccountCmd).Standalone()

	inspector2_enableDelegatedAdminAccountCmd.Flags().String("client-token", "", "The idempotency token for the request.")
	inspector2_enableDelegatedAdminAccountCmd.Flags().String("delegated-admin-account-id", "", "The Amazon Web Services account ID of the Amazon Inspector delegated administrator.")
	inspector2_enableDelegatedAdminAccountCmd.MarkFlagRequired("delegated-admin-account-id")
	inspector2Cmd.AddCommand(inspector2_enableDelegatedAdminAccountCmd)
}
