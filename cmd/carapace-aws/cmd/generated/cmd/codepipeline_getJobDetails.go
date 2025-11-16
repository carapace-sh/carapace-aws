package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_getJobDetailsCmd = &cobra.Command{
	Use:   "get-job-details",
	Short: "Returns information about a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_getJobDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_getJobDetailsCmd).Standalone()

		codepipeline_getJobDetailsCmd.Flags().String("job-id", "", "The unique system-generated ID for the job.")
		codepipeline_getJobDetailsCmd.MarkFlagRequired("job-id")
	})
	codepipelineCmd.AddCommand(codepipeline_getJobDetailsCmd)
}
