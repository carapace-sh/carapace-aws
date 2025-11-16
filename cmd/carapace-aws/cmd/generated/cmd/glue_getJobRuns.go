package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getJobRunsCmd = &cobra.Command{
	Use:   "get-job-runs",
	Short: "Retrieves metadata for all runs of a given job definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getJobRunsCmd).Standalone()

	glue_getJobRunsCmd.Flags().String("job-name", "", "The name of the job definition for which to retrieve all job runs.")
	glue_getJobRunsCmd.Flags().String("max-results", "", "The maximum size of the response.")
	glue_getJobRunsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
	glue_getJobRunsCmd.MarkFlagRequired("job-name")
	glueCmd.AddCommand(glue_getJobRunsCmd)
}
