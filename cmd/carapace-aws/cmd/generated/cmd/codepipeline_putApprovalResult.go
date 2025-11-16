package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_putApprovalResultCmd = &cobra.Command{
	Use:   "put-approval-result",
	Short: "Provides the response to a manual approval request to CodePipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_putApprovalResultCmd).Standalone()

	codepipeline_putApprovalResultCmd.Flags().String("action-name", "", "The name of the action for which approval is requested.")
	codepipeline_putApprovalResultCmd.Flags().String("pipeline-name", "", "The name of the pipeline that contains the action.")
	codepipeline_putApprovalResultCmd.Flags().String("result", "", "Represents information about the result of the approval request.")
	codepipeline_putApprovalResultCmd.Flags().String("stage-name", "", "The name of the stage that contains the action.")
	codepipeline_putApprovalResultCmd.Flags().String("token", "", "The system-generated token used to identify a unique approval request.")
	codepipeline_putApprovalResultCmd.MarkFlagRequired("action-name")
	codepipeline_putApprovalResultCmd.MarkFlagRequired("pipeline-name")
	codepipeline_putApprovalResultCmd.MarkFlagRequired("result")
	codepipeline_putApprovalResultCmd.MarkFlagRequired("stage-name")
	codepipeline_putApprovalResultCmd.MarkFlagRequired("token")
	codepipelineCmd.AddCommand(codepipeline_putApprovalResultCmd)
}
