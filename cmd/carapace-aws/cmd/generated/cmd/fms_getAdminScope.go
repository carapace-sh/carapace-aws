package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_getAdminScopeCmd = &cobra.Command{
	Use:   "get-admin-scope",
	Short: "Returns information about the specified account's administrative scope.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_getAdminScopeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_getAdminScopeCmd).Standalone()

		fms_getAdminScopeCmd.Flags().String("admin-account", "", "The administrator account that you want to get the details for.")
		fms_getAdminScopeCmd.MarkFlagRequired("admin-account")
	})
	fmsCmd.AddCommand(fms_getAdminScopeCmd)
}
