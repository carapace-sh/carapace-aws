package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_disassociateFromMasterAccountCmd = &cobra.Command{
	Use:   "disassociate-from-master-account",
	Short: "(Deprecated) Disassociates a member account from its Amazon Macie administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_disassociateFromMasterAccountCmd).Standalone()

	macie2Cmd.AddCommand(macie2_disassociateFromMasterAccountCmd)
}
