package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_putJobFailureResultCmd = &cobra.Command{
	Use:   "put-job-failure-result",
	Short: "Represents the failure of a job as returned to the pipeline by a job worker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_putJobFailureResultCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_putJobFailureResultCmd).Standalone()

		codepipeline_putJobFailureResultCmd.Flags().String("failure-details", "", "The details about the failure of a job.")
		codepipeline_putJobFailureResultCmd.Flags().String("job-id", "", "The unique system-generated ID of the job that failed.")
		codepipeline_putJobFailureResultCmd.MarkFlagRequired("failure-details")
		codepipeline_putJobFailureResultCmd.MarkFlagRequired("job-id")
	})
	codepipelineCmd.AddCommand(codepipeline_putJobFailureResultCmd)
}
