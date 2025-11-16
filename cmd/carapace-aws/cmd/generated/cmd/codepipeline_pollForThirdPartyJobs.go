package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_pollForThirdPartyJobsCmd = &cobra.Command{
	Use:   "poll-for-third-party-jobs",
	Short: "Determines whether there are any third party jobs for a job worker to act on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_pollForThirdPartyJobsCmd).Standalone()

	codepipeline_pollForThirdPartyJobsCmd.Flags().String("action-type-id", "", "Represents information about an action type.")
	codepipeline_pollForThirdPartyJobsCmd.Flags().String("max-batch-size", "", "The maximum number of jobs to return in a poll for jobs call.")
	codepipeline_pollForThirdPartyJobsCmd.MarkFlagRequired("action-type-id")
	codepipelineCmd.AddCommand(codepipeline_pollForThirdPartyJobsCmd)
}
