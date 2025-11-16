package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_acknowledgeThirdPartyJobCmd = &cobra.Command{
	Use:   "acknowledge-third-party-job",
	Short: "Confirms a job worker has received the specified job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_acknowledgeThirdPartyJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_acknowledgeThirdPartyJobCmd).Standalone()

		codepipeline_acknowledgeThirdPartyJobCmd.Flags().String("client-token", "", "The clientToken portion of the clientId and clientToken pair used to verify that the calling entity is allowed access to the job and its details.")
		codepipeline_acknowledgeThirdPartyJobCmd.Flags().String("job-id", "", "The unique system-generated ID of the job.")
		codepipeline_acknowledgeThirdPartyJobCmd.Flags().String("nonce", "", "A system-generated random number that CodePipeline uses to ensure that the job is being worked on by only one job worker.")
		codepipeline_acknowledgeThirdPartyJobCmd.MarkFlagRequired("client-token")
		codepipeline_acknowledgeThirdPartyJobCmd.MarkFlagRequired("job-id")
		codepipeline_acknowledgeThirdPartyJobCmd.MarkFlagRequired("nonce")
	})
	codepipelineCmd.AddCommand(codepipeline_acknowledgeThirdPartyJobCmd)
}
