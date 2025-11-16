package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_listCopyJobSummariesCmd = &cobra.Command{
	Use:   "list-copy-job-summaries",
	Short: "This request obtains a list of copy jobs created or running within the the most recent 30 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_listCopyJobSummariesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_listCopyJobSummariesCmd).Standalone()

		backup_listCopyJobSummariesCmd.Flags().String("account-id", "", "Returns the job count for the specified account.")
		backup_listCopyJobSummariesCmd.Flags().String("aggregation-period", "", "The period for the returned results.")
		backup_listCopyJobSummariesCmd.Flags().String("max-results", "", "This parameter sets the maximum number of items to be returned.")
		backup_listCopyJobSummariesCmd.Flags().String("message-category", "", "This parameter returns the job count for the specified message category.")
		backup_listCopyJobSummariesCmd.Flags().String("next-token", "", "The next item following a partial list of returned resources.")
		backup_listCopyJobSummariesCmd.Flags().String("resource-type", "", "Returns the job count for the specified resource type.")
		backup_listCopyJobSummariesCmd.Flags().String("state", "", "This parameter returns the job count for jobs with the specified state.")
	})
	backupCmd.AddCommand(backup_listCopyJobSummariesCmd)
}
