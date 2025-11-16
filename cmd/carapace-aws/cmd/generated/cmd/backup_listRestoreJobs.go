package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listRestoreJobsCmd = &cobra.Command{
	Use:   "list-restore-jobs",
	Short: "Returns a list of jobs that Backup initiated to restore a saved resource, including details about the recovery process.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listRestoreJobsCmd).Standalone()

	backup_listRestoreJobsCmd.Flags().String("by-account-id", "", "The account ID to list the jobs from.")
	backup_listRestoreJobsCmd.Flags().String("by-complete-after", "", "Returns only copy jobs completed after a date expressed in Unix format and Coordinated Universal Time (UTC).")
	backup_listRestoreJobsCmd.Flags().String("by-complete-before", "", "Returns only copy jobs completed before a date expressed in Unix format and Coordinated Universal Time (UTC).")
	backup_listRestoreJobsCmd.Flags().String("by-created-after", "", "Returns only restore jobs that were created after the specified date.")
	backup_listRestoreJobsCmd.Flags().String("by-created-before", "", "Returns only restore jobs that were created before the specified date.")
	backup_listRestoreJobsCmd.Flags().String("by-parent-job-id", "", "This is a filter to list child (nested) restore jobs based on parent restore job ID.")
	backup_listRestoreJobsCmd.Flags().String("by-resource-type", "", "Include this parameter to return only restore jobs for the specified resources:")
	backup_listRestoreJobsCmd.Flags().String("by-restore-testing-plan-arn", "", "This returns only restore testing jobs that match the specified resource Amazon Resource Name (ARN).")
	backup_listRestoreJobsCmd.Flags().String("by-status", "", "Returns only restore jobs associated with the specified job status.")
	backup_listRestoreJobsCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
	backup_listRestoreJobsCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
	backupCmd.AddCommand(backup_listRestoreJobsCmd)
}
