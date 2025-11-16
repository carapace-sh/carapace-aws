package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_cancelAutomatedReasoningPolicyBuildWorkflowCmd = &cobra.Command{
	Use:   "cancel-automated-reasoning-policy-build-workflow",
	Short: "Cancels a running Automated Reasoning policy build workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_cancelAutomatedReasoningPolicyBuildWorkflowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_cancelAutomatedReasoningPolicyBuildWorkflowCmd).Standalone()

		bedrock_cancelAutomatedReasoningPolicyBuildWorkflowCmd.Flags().String("build-workflow-id", "", "The unique identifier of the build workflow to cancel.")
		bedrock_cancelAutomatedReasoningPolicyBuildWorkflowCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy whose build workflow you want to cancel.")
		bedrock_cancelAutomatedReasoningPolicyBuildWorkflowCmd.MarkFlagRequired("build-workflow-id")
		bedrock_cancelAutomatedReasoningPolicyBuildWorkflowCmd.MarkFlagRequired("policy-arn")
	})
	bedrockCmd.AddCommand(bedrock_cancelAutomatedReasoningPolicyBuildWorkflowCmd)
}
