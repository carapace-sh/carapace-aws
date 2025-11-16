package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_putThirdPartyJobFailureResultCmd = &cobra.Command{
	Use:   "put-third-party-job-failure-result",
	Short: "Represents the failure of a third party job as returned to the pipeline by a job worker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_putThirdPartyJobFailureResultCmd).Standalone()

	codepipeline_putThirdPartyJobFailureResultCmd.Flags().String("client-token", "", "The clientToken portion of the clientId and clientToken pair used to verify that the calling entity is allowed access to the job and its details.")
	codepipeline_putThirdPartyJobFailureResultCmd.Flags().String("failure-details", "", "Represents information about failure details.")
	codepipeline_putThirdPartyJobFailureResultCmd.Flags().String("job-id", "", "The ID of the job that failed.")
	codepipeline_putThirdPartyJobFailureResultCmd.MarkFlagRequired("client-token")
	codepipeline_putThirdPartyJobFailureResultCmd.MarkFlagRequired("failure-details")
	codepipeline_putThirdPartyJobFailureResultCmd.MarkFlagRequired("job-id")
	codepipelineCmd.AddCommand(codepipeline_putThirdPartyJobFailureResultCmd)
}
