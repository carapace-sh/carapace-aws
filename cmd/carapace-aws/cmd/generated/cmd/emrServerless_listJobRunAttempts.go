package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_listJobRunAttemptsCmd = &cobra.Command{
	Use:   "list-job-run-attempts",
	Short: "Lists all attempt of a job run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_listJobRunAttemptsCmd).Standalone()

	emrServerless_listJobRunAttemptsCmd.Flags().String("application-id", "", "The ID of the application for which to list job runs.")
	emrServerless_listJobRunAttemptsCmd.Flags().String("job-run-id", "", "The ID of the job run to list.")
	emrServerless_listJobRunAttemptsCmd.Flags().String("max-results", "", "The maximum number of job run attempts to list.")
	emrServerless_listJobRunAttemptsCmd.Flags().String("next-token", "", "The token for the next set of job run attempt results.")
	emrServerless_listJobRunAttemptsCmd.MarkFlagRequired("application-id")
	emrServerless_listJobRunAttemptsCmd.MarkFlagRequired("job-run-id")
	emrServerlessCmd.AddCommand(emrServerless_listJobRunAttemptsCmd)
}
