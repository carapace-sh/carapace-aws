package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_disassociateFromAdministratorAccountCmd = &cobra.Command{
	Use:   "disassociate-from-administrator-account",
	Short: "Disassociates the current Security Hub member account from the associated administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_disassociateFromAdministratorAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_disassociateFromAdministratorAccountCmd).Standalone()

	})
	securityhubCmd.AddCommand(securityhub_disassociateFromAdministratorAccountCmd)
}
