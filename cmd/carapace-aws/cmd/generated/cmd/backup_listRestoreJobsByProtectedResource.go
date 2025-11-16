package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listRestoreJobsByProtectedResourceCmd = &cobra.Command{
	Use:   "list-restore-jobs-by-protected-resource",
	Short: "This returns restore jobs that contain the specified protected resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listRestoreJobsByProtectedResourceCmd).Standalone()

	backup_listRestoreJobsByProtectedResourceCmd.Flags().String("by-recovery-point-creation-date-after", "", "Returns only restore jobs of recovery points that were created after the specified date.")
	backup_listRestoreJobsByProtectedResourceCmd.Flags().String("by-recovery-point-creation-date-before", "", "Returns only restore jobs of recovery points that were created before the specified date.")
	backup_listRestoreJobsByProtectedResourceCmd.Flags().String("by-status", "", "Returns only restore jobs associated with the specified job status.")
	backup_listRestoreJobsByProtectedResourceCmd.Flags().String("max-results", "", "The maximum number of items to be returned.")
	backup_listRestoreJobsByProtectedResourceCmd.Flags().String("next-token", "", "The next item following a partial list of returned items.")
	backup_listRestoreJobsByProtectedResourceCmd.Flags().String("resource-arn", "", "Returns only restore jobs that match the specified resource Amazon Resource Name (ARN).")
	backup_listRestoreJobsByProtectedResourceCmd.MarkFlagRequired("resource-arn")
	backupCmd.AddCommand(backup_listRestoreJobsByProtectedResourceCmd)
}
