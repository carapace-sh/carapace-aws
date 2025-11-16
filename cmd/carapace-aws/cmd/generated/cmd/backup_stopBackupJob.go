package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_stopBackupJobCmd = &cobra.Command{
	Use:   "stop-backup-job",
	Short: "Attempts to cancel a job to create a one-time backup of a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_stopBackupJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_stopBackupJobCmd).Standalone()

		backup_stopBackupJobCmd.Flags().String("backup-job-id", "", "Uniquely identifies a request to Backup to back up a resource.")
		backup_stopBackupJobCmd.MarkFlagRequired("backup-job-id")
	})
	backupCmd.AddCommand(backup_stopBackupJobCmd)
}
