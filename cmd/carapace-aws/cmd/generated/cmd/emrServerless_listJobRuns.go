package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_listJobRunsCmd = &cobra.Command{
	Use:   "list-job-runs",
	Short: "Lists job runs based on a set of parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_listJobRunsCmd).Standalone()

	emrServerless_listJobRunsCmd.Flags().String("application-id", "", "The ID of the application for which to list the job run.")
	emrServerless_listJobRunsCmd.Flags().String("created-at-after", "", "The lower bound of the option to filter by creation date and time.")
	emrServerless_listJobRunsCmd.Flags().String("created-at-before", "", "The upper bound of the option to filter by creation date and time.")
	emrServerless_listJobRunsCmd.Flags().String("max-results", "", "The maximum number of job runs that can be listed.")
	emrServerless_listJobRunsCmd.Flags().String("mode", "", "The mode of the job runs to list.")
	emrServerless_listJobRunsCmd.Flags().String("next-token", "", "The token for the next set of job run results.")
	emrServerless_listJobRunsCmd.Flags().String("states", "", "An optional filter for job run states.")
	emrServerless_listJobRunsCmd.MarkFlagRequired("application-id")
	emrServerlessCmd.AddCommand(emrServerless_listJobRunsCmd)
}
