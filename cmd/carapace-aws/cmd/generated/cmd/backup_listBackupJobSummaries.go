package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listBackupJobSummariesCmd = &cobra.Command{
	Use:   "list-backup-job-summaries",
	Short: "This is a request for a summary of backup jobs created or running within the most recent 30 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listBackupJobSummariesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_listBackupJobSummariesCmd).Standalone()

		backup_listBackupJobSummariesCmd.Flags().String("account-id", "", "Returns the job count for the specified account.")
		backup_listBackupJobSummariesCmd.Flags().String("aggregation-period", "", "The period for the returned results.")
		backup_listBackupJobSummariesCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
		backup_listBackupJobSummariesCmd.Flags().String("message-category", "", "This parameter returns the job count for the specified message category.")
		backup_listBackupJobSummariesCmd.Flags().String("next-token", "", "The next item following a partial list of returned resources.")
		backup_listBackupJobSummariesCmd.Flags().String("resource-type", "", "Returns the job count for the specified resource type.")
		backup_listBackupJobSummariesCmd.Flags().String("state", "", "This parameter returns the job count for jobs with the specified state.")
	})
	backupCmd.AddCommand(backup_listBackupJobSummariesCmd)
}
