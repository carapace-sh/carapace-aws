package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchGetJobsCmd = &cobra.Command{
	Use:   "batch-get-jobs",
	Short: "Returns a list of resource metadata for a given list of job names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchGetJobsCmd).Standalone()

	glue_batchGetJobsCmd.Flags().String("job-names", "", "A list of job names, which might be the names returned from the `ListJobs` operation.")
	glue_batchGetJobsCmd.MarkFlagRequired("job-names")
	glueCmd.AddCommand(glue_batchGetJobsCmd)
}
