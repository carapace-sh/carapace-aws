package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_createLegalHoldCmd = &cobra.Command{
	Use:   "create-legal-hold",
	Short: "Creates a legal hold on a recovery point (backup).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_createLegalHoldCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_createLegalHoldCmd).Standalone()

		backup_createLegalHoldCmd.Flags().String("description", "", "The description of the legal hold.")
		backup_createLegalHoldCmd.Flags().String("idempotency-token", "", "This is a user-chosen string used to distinguish between otherwise identical calls.")
		backup_createLegalHoldCmd.Flags().String("recovery-point-selection", "", "The criteria to assign a set of resources, such as resource types or backup vaults.")
		backup_createLegalHoldCmd.Flags().String("tags", "", "Optional tags to include.")
		backup_createLegalHoldCmd.Flags().String("title", "", "The title of the legal hold.")
		backup_createLegalHoldCmd.MarkFlagRequired("description")
		backup_createLegalHoldCmd.MarkFlagRequired("title")
	})
	backupCmd.AddCommand(backup_createLegalHoldCmd)
}
