package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_getThirdPartyJobDetailsCmd = &cobra.Command{
	Use:   "get-third-party-job-details",
	Short: "Requests the details of a job for a third party action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_getThirdPartyJobDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_getThirdPartyJobDetailsCmd).Standalone()

		codepipeline_getThirdPartyJobDetailsCmd.Flags().String("client-token", "", "The clientToken portion of the clientId and clientToken pair used to verify that the calling entity is allowed access to the job and its details.")
		codepipeline_getThirdPartyJobDetailsCmd.Flags().String("job-id", "", "The unique system-generated ID used for identifying the job.")
		codepipeline_getThirdPartyJobDetailsCmd.MarkFlagRequired("client-token")
		codepipeline_getThirdPartyJobDetailsCmd.MarkFlagRequired("job-id")
	})
	codepipelineCmd.AddCommand(codepipeline_getThirdPartyJobDetailsCmd)
}
