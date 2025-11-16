package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getAccountStatusCmd = &cobra.Command{
	Use:   "get-account-status",
	Short: "Gets the registration status of an account in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getAccountStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_getAccountStatusCmd).Standalone()

	})
	auditmanagerCmd.AddCommand(auditmanager_getAccountStatusCmd)
}
