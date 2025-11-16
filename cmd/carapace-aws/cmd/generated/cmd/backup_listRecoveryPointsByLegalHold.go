package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listRecoveryPointsByLegalHoldCmd = &cobra.Command{
	Use:   "list-recovery-points-by-legal-hold",
	Short: "This action returns recovery point ARNs (Amazon Resource Names) of the specified legal hold.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listRecoveryPointsByLegalHoldCmd).Standalone()

	backup_listRecoveryPointsByLegalHoldCmd.Flags().String("legal-hold-id", "", "The ID of the legal hold.")
	backup_listRecoveryPointsByLegalHoldCmd.Flags().String("max-results", "", "The maximum number of resource list items to be returned.")
	backup_listRecoveryPointsByLegalHoldCmd.Flags().String("next-token", "", "The next item following a partial list of returned resources.")
	backup_listRecoveryPointsByLegalHoldCmd.MarkFlagRequired("legal-hold-id")
	backupCmd.AddCommand(backup_listRecoveryPointsByLegalHoldCmd)
}
