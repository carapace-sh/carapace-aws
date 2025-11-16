package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_disassociateFromAdministratorAccountCmd = &cobra.Command{
	Use:   "disassociate-from-administrator-account",
	Short: "Disassociates a member account from its Amazon Macie administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_disassociateFromAdministratorAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_disassociateFromAdministratorAccountCmd).Standalone()

	})
	macie2Cmd.AddCommand(macie2_disassociateFromAdministratorAccountCmd)
}
