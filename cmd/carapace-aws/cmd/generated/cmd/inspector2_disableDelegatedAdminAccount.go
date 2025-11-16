package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_disableDelegatedAdminAccountCmd = &cobra.Command{
	Use:   "disable-delegated-admin-account",
	Short: "Disables the Amazon Inspector delegated administrator for your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_disableDelegatedAdminAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_disableDelegatedAdminAccountCmd).Standalone()

		inspector2_disableDelegatedAdminAccountCmd.Flags().String("delegated-admin-account-id", "", "The Amazon Web Services account ID of the current Amazon Inspector delegated administrator.")
		inspector2_disableDelegatedAdminAccountCmd.MarkFlagRequired("delegated-admin-account-id")
	})
	inspector2Cmd.AddCommand(inspector2_disableDelegatedAdminAccountCmd)
}
