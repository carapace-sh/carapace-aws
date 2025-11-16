package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_disassociateFromMasterAccountCmd = &cobra.Command{
	Use:   "disassociate-from-master-account",
	Short: "This method is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_disassociateFromMasterAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_disassociateFromMasterAccountCmd).Standalone()

	})
	securityhubCmd.AddCommand(securityhub_disassociateFromMasterAccountCmd)
}
