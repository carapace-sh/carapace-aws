package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_restoreBackupCmd = &cobra.Command{
	Use:   "restore-backup",
	Short: "Restores a specified CloudHSM backup that is in the `PENDING_DELETION` state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_restoreBackupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudhsmv2_restoreBackupCmd).Standalone()

		cloudhsmv2_restoreBackupCmd.Flags().String("backup-id", "", "The ID of the backup to be restored.")
		cloudhsmv2_restoreBackupCmd.MarkFlagRequired("backup-id")
	})
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_restoreBackupCmd)
}
