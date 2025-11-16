package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_cancelLegalHoldCmd = &cobra.Command{
	Use:   "cancel-legal-hold",
	Short: "Removes the specified legal hold on a recovery point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_cancelLegalHoldCmd).Standalone()

	backup_cancelLegalHoldCmd.Flags().String("cancel-description", "", "A string the describes the reason for removing the legal hold.")
	backup_cancelLegalHoldCmd.Flags().String("legal-hold-id", "", "The ID of the legal hold.")
	backup_cancelLegalHoldCmd.Flags().String("retain-record-in-days", "", "The integer amount, in days, after which to remove legal hold.")
	backup_cancelLegalHoldCmd.MarkFlagRequired("cancel-description")
	backup_cancelLegalHoldCmd.MarkFlagRequired("legal-hold-id")
	backupCmd.AddCommand(backup_cancelLegalHoldCmd)
}
