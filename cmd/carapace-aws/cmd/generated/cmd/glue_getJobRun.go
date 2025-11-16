package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getJobRunCmd = &cobra.Command{
	Use:   "get-job-run",
	Short: "Retrieves the metadata for a given job run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getJobRunCmd).Standalone()

	glue_getJobRunCmd.Flags().String("job-name", "", "Name of the job definition being run.")
	glue_getJobRunCmd.Flags().String("predecessors-included", "", "True if a list of predecessor runs should be returned.")
	glue_getJobRunCmd.Flags().String("run-id", "", "The ID of the job run.")
	glue_getJobRunCmd.MarkFlagRequired("job-name")
	glue_getJobRunCmd.MarkFlagRequired("run-id")
	glueCmd.AddCommand(glue_getJobRunCmd)
}
