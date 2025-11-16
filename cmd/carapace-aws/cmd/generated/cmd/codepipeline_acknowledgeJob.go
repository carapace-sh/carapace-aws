package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_acknowledgeJobCmd = &cobra.Command{
	Use:   "acknowledge-job",
	Short: "Returns information about a specified job and whether that job has been received by the job worker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_acknowledgeJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_acknowledgeJobCmd).Standalone()

		codepipeline_acknowledgeJobCmd.Flags().String("job-id", "", "The unique system-generated ID of the job for which you want to confirm receipt.")
		codepipeline_acknowledgeJobCmd.Flags().String("nonce", "", "A system-generated random number that CodePipeline uses to ensure that the job is being worked on by only one job worker.")
		codepipeline_acknowledgeJobCmd.MarkFlagRequired("job-id")
		codepipeline_acknowledgeJobCmd.MarkFlagRequired("nonce")
	})
	codepipelineCmd.AddCommand(codepipeline_acknowledgeJobCmd)
}
