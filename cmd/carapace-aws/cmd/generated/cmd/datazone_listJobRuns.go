package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listJobRunsCmd = &cobra.Command{
	Use:   "list-job-runs",
	Short: "Lists job runs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listJobRunsCmd).Standalone()

	datazone_listJobRunsCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to list job runs.")
	datazone_listJobRunsCmd.Flags().String("job-identifier", "", "The ID of the job run.")
	datazone_listJobRunsCmd.Flags().String("max-results", "", "The maximum number of job runs to return in a single call to ListJobRuns.")
	datazone_listJobRunsCmd.Flags().String("next-token", "", "When the number of job runs is greater than the default value for the MaxResults parameter, or if you explicitly specify a value for MaxResults that is less than the number of job runs, the response includes a pagination token named NextToken.")
	datazone_listJobRunsCmd.Flags().String("sort-order", "", "Specifies the order in which job runs are to be sorted.")
	datazone_listJobRunsCmd.Flags().String("status", "", "The status of a job run.")
	datazone_listJobRunsCmd.MarkFlagRequired("domain-identifier")
	datazone_listJobRunsCmd.MarkFlagRequired("job-identifier")
	datazoneCmd.AddCommand(datazone_listJobRunsCmd)
}
