package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_getDelegatedAdminAccountCmd = &cobra.Command{
	Use:   "get-delegated-admin-account",
	Short: "Retrieves information about the Amazon Inspector delegated administrator for your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_getDelegatedAdminAccountCmd).Standalone()

	inspector2Cmd.AddCommand(inspector2_getDelegatedAdminAccountCmd)
}
