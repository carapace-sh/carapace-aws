package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_getAdminAccountCmd = &cobra.Command{
	Use:   "get-admin-account",
	Short: "Returns the Organizations account that is associated with Firewall Manager as the Firewall Manager default administrator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_getAdminAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_getAdminAccountCmd).Standalone()

	})
	fmsCmd.AddCommand(fms_getAdminAccountCmd)
}
