package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getAdministratorAccountCmd = &cobra.Command{
	Use:   "get-administrator-account",
	Short: "Provides the details for the Security Hub administrator account for the current member account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getAdministratorAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_getAdministratorAccountCmd).Standalone()

	})
	securityhubCmd.AddCommand(securityhub_getAdministratorAccountCmd)
}
