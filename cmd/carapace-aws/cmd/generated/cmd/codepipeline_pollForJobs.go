package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_pollForJobsCmd = &cobra.Command{
	Use:   "poll-for-jobs",
	Short: "Returns information about any jobs for CodePipeline to act on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_pollForJobsCmd).Standalone()

	codepipeline_pollForJobsCmd.Flags().String("action-type-id", "", "Represents information about an action type.")
	codepipeline_pollForJobsCmd.Flags().String("max-batch-size", "", "The maximum number of jobs to return in a poll for jobs call.")
	codepipeline_pollForJobsCmd.Flags().String("query-param", "", "A map of property names and values.")
	codepipeline_pollForJobsCmd.MarkFlagRequired("action-type-id")
	codepipelineCmd.AddCommand(codepipeline_pollForJobsCmd)
}
