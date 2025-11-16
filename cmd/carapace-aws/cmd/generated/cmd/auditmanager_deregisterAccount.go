package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_deregisterAccountCmd = &cobra.Command{
	Use:   "deregister-account",
	Short: "Deregisters an account in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_deregisterAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_deregisterAccountCmd).Standalone()

	})
	auditmanagerCmd.AddCommand(auditmanager_deregisterAccountCmd)
}
