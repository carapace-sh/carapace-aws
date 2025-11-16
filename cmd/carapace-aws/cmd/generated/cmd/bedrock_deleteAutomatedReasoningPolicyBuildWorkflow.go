package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_deleteAutomatedReasoningPolicyBuildWorkflowCmd = &cobra.Command{
	Use:   "delete-automated-reasoning-policy-build-workflow",
	Short: "Deletes an Automated Reasoning policy build workflow and its associated artifacts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_deleteAutomatedReasoningPolicyBuildWorkflowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_deleteAutomatedReasoningPolicyBuildWorkflowCmd).Standalone()

		bedrock_deleteAutomatedReasoningPolicyBuildWorkflowCmd.Flags().String("build-workflow-id", "", "The unique identifier of the build workflow to delete.")
		bedrock_deleteAutomatedReasoningPolicyBuildWorkflowCmd.Flags().String("last-updated-at", "", "The timestamp when the build workflow was last updated.")
		bedrock_deleteAutomatedReasoningPolicyBuildWorkflowCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy whose build workflow you want to delete.")
		bedrock_deleteAutomatedReasoningPolicyBuildWorkflowCmd.MarkFlagRequired("build-workflow-id")
		bedrock_deleteAutomatedReasoningPolicyBuildWorkflowCmd.MarkFlagRequired("last-updated-at")
		bedrock_deleteAutomatedReasoningPolicyBuildWorkflowCmd.MarkFlagRequired("policy-arn")
	})
	bedrockCmd.AddCommand(bedrock_deleteAutomatedReasoningPolicyBuildWorkflowCmd)
}
