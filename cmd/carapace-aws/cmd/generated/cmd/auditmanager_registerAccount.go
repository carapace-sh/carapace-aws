package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_registerAccountCmd = &cobra.Command{
	Use:   "register-account",
	Short: "Enables Audit Manager for the specified Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_registerAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_registerAccountCmd).Standalone()

		auditmanager_registerAccountCmd.Flags().String("delegated-admin-account", "", "The delegated administrator account for Audit Manager.")
		auditmanager_registerAccountCmd.Flags().String("kms-key", "", "The KMS key details.")
	})
	auditmanagerCmd.AddCommand(auditmanager_registerAccountCmd)
}
