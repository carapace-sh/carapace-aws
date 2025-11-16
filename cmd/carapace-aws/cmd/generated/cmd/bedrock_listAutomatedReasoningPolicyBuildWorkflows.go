package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listAutomatedReasoningPolicyBuildWorkflowsCmd = &cobra.Command{
	Use:   "list-automated-reasoning-policy-build-workflows",
	Short: "Lists all build workflows for an Automated Reasoning policy, showing the history of policy creation and modification attempts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listAutomatedReasoningPolicyBuildWorkflowsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_listAutomatedReasoningPolicyBuildWorkflowsCmd).Standalone()

		bedrock_listAutomatedReasoningPolicyBuildWorkflowsCmd.Flags().String("max-results", "", "The maximum number of build workflows to return in a single response.")
		bedrock_listAutomatedReasoningPolicyBuildWorkflowsCmd.Flags().String("next-token", "", "A pagination token from a previous request to continue listing build workflows from where the previous request left off.")
		bedrock_listAutomatedReasoningPolicyBuildWorkflowsCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy whose build workflows you want to list.")
		bedrock_listAutomatedReasoningPolicyBuildWorkflowsCmd.MarkFlagRequired("policy-arn")
	})
	bedrockCmd.AddCommand(bedrock_listAutomatedReasoningPolicyBuildWorkflowsCmd)
}
