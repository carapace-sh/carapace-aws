package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_disassociateAdminAccountCmd = &cobra.Command{
	Use:   "disassociate-admin-account",
	Short: "Disassociates an Firewall Manager administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_disassociateAdminAccountCmd).Standalone()

	fmsCmd.AddCommand(fms_disassociateAdminAccountCmd)
}
