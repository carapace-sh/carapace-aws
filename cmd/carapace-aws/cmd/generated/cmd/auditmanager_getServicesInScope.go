package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getServicesInScopeCmd = &cobra.Command{
	Use:   "get-services-in-scope",
	Short: "Gets a list of the Amazon Web Services services from which Audit Manager can collect evidence.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getServicesInScopeCmd).Standalone()

	auditmanagerCmd.AddCommand(auditmanager_getServicesInScopeCmd)
}
