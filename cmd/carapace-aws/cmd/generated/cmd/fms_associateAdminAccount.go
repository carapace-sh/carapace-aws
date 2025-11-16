package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_associateAdminAccountCmd = &cobra.Command{
	Use:   "associate-admin-account",
	Short: "Sets a Firewall Manager default administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_associateAdminAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_associateAdminAccountCmd).Standalone()

		fms_associateAdminAccountCmd.Flags().String("admin-account", "", "The Amazon Web Services account ID to associate with Firewall Manager as the Firewall Manager default administrator account.")
		fms_associateAdminAccountCmd.MarkFlagRequired("admin-account")
	})
	fmsCmd.AddCommand(fms_associateAdminAccountCmd)
}
