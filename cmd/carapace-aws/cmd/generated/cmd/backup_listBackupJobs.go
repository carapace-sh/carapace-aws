package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listBackupJobsCmd = &cobra.Command{
	Use:   "list-backup-jobs",
	Short: "Returns a list of existing backup jobs for an authenticated account for the last 30 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listBackupJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_listBackupJobsCmd).Standalone()

		backup_listBackupJobsCmd.Flags().String("by-account-id", "", "The account ID to list the jobs from.")
		backup_listBackupJobsCmd.Flags().String("by-backup-vault-name", "", "Returns only backup jobs that will be stored in the specified backup vault.")
		backup_listBackupJobsCmd.Flags().String("by-complete-after", "", "Returns only backup jobs completed after a date expressed in Unix format and Coordinated Universal Time (UTC).")
		backup_listBackupJobsCmd.Flags().String("by-complete-before", "", "Returns only backup jobs completed before a date expressed in Unix format and Coordinated Universal Time (UTC).")
		backup_listBackupJobsCmd.Flags().String("by-created-after", "", "Returns only backup jobs that were created after the specified date.")
		backup_listBackupJobsCmd.Flags().String("by-created-before", "", "Returns only backup jobs that were created before the specified date.")
		backup_listBackupJobsCmd.Flags().String("by-message-category", "", "This is an optional parameter that can be used to filter out jobs with a MessageCategory which matches the value you input.")
		backup_listBackupJobsCmd.Flags().String("by-parent-job-id", "", "This is a filter to list child (nested) jobs based on parent job ID.")
		backup_listBackupJobsCmd.Flags().String("by-resource-arn", "", "Returns only backup jobs that match the specified resource Amazon Resource Name (ARN).")
		backup_listBackupJobsCmd.Flags().String("by-resource-type", "", "Returns only backup jobs for the specified resources:")
		backup_listBackupJobsCmd.Flags().String("by-state", "", "Returns only backup jobs that are in the specified state.")
		backup_listBackupJobsCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
		backup_listBackupJobsCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
	})
	backupCmd.AddCommand(backup_listBackupJobsCmd)
}
