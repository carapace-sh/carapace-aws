package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_describeBackupJobCmd = &cobra.Command{
	Use:   "describe-backup-job",
	Short: "Returns backup job details for the specified `BackupJobId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_describeBackupJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_describeBackupJobCmd).Standalone()

		backup_describeBackupJobCmd.Flags().String("backup-job-id", "", "Uniquely identifies a request to Backup to back up a resource.")
		backup_describeBackupJobCmd.MarkFlagRequired("backup-job-id")
	})
	backupCmd.AddCommand(backup_describeBackupJobCmd)
}
