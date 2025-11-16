package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanagerCmd = &cobra.Command{
	Use:   "auditmanager",
	Short: "Welcome to the Audit Manager API reference.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanagerCmd).Standalone()

	rootCmd.AddCommand(auditmanagerCmd)
}
