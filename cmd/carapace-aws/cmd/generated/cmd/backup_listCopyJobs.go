package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listCopyJobsCmd = &cobra.Command{
	Use:   "list-copy-jobs",
	Short: "Returns metadata about your copy jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listCopyJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_listCopyJobsCmd).Standalone()

		backup_listCopyJobsCmd.Flags().String("by-account-id", "", "The account ID to list the jobs from.")
		backup_listCopyJobsCmd.Flags().String("by-complete-after", "", "Returns only copy jobs completed after a date expressed in Unix format and Coordinated Universal Time (UTC).")
		backup_listCopyJobsCmd.Flags().String("by-complete-before", "", "Returns only copy jobs completed before a date expressed in Unix format and Coordinated Universal Time (UTC).")
		backup_listCopyJobsCmd.Flags().String("by-created-after", "", "Returns only copy jobs that were created after the specified date.")
		backup_listCopyJobsCmd.Flags().String("by-created-before", "", "Returns only copy jobs that were created before the specified date.")
		backup_listCopyJobsCmd.Flags().String("by-destination-vault-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies a source backup vault to copy from; for example, `arn:aws:backup:us-east-1:123456789012:backup-vault:aBackupVault`.")
		backup_listCopyJobsCmd.Flags().String("by-message-category", "", "This is an optional parameter that can be used to filter out jobs with a MessageCategory which matches the value you input.")
		backup_listCopyJobsCmd.Flags().String("by-parent-job-id", "", "This is a filter to list child (nested) jobs based on parent job ID.")
		backup_listCopyJobsCmd.Flags().String("by-resource-arn", "", "Returns only copy jobs that match the specified resource Amazon Resource Name (ARN).")
		backup_listCopyJobsCmd.Flags().String("by-resource-type", "", "Returns only backup jobs for the specified resources:")
		backup_listCopyJobsCmd.Flags().String("by-state", "", "Returns only copy jobs that are in the specified state.")
		backup_listCopyJobsCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
		backup_listCopyJobsCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
	})
	backupCmd.AddCommand(backup_listCopyJobsCmd)
}
