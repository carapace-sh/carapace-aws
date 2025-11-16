package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getAutomatedReasoningPolicyBuildWorkflowCmd = &cobra.Command{
	Use:   "get-automated-reasoning-policy-build-workflow",
	Short: "Retrieves detailed information about an Automated Reasoning policy build workflow, including its status, configuration, and metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getAutomatedReasoningPolicyBuildWorkflowCmd).Standalone()

	bedrock_getAutomatedReasoningPolicyBuildWorkflowCmd.Flags().String("build-workflow-id", "", "The unique identifier of the build workflow to retrieve.")
	bedrock_getAutomatedReasoningPolicyBuildWorkflowCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy whose build workflow you want to retrieve.")
	bedrock_getAutomatedReasoningPolicyBuildWorkflowCmd.MarkFlagRequired("build-workflow-id")
	bedrock_getAutomatedReasoningPolicyBuildWorkflowCmd.MarkFlagRequired("policy-arn")
	bedrockCmd.AddCommand(bedrock_getAutomatedReasoningPolicyBuildWorkflowCmd)
}
