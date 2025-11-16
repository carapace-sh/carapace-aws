package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getAutomatedReasoningPolicyNextScenarioCmd = &cobra.Command{
	Use:   "get-automated-reasoning-policy-next-scenario",
	Short: "Retrieves the next test scenario for validating an Automated Reasoning policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getAutomatedReasoningPolicyNextScenarioCmd).Standalone()

	bedrock_getAutomatedReasoningPolicyNextScenarioCmd.Flags().String("build-workflow-id", "", "The unique identifier of the build workflow associated with the test scenarios.")
	bedrock_getAutomatedReasoningPolicyNextScenarioCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy for which you want to get the next test scenario.")
	bedrock_getAutomatedReasoningPolicyNextScenarioCmd.MarkFlagRequired("build-workflow-id")
	bedrock_getAutomatedReasoningPolicyNextScenarioCmd.MarkFlagRequired("policy-arn")
	bedrockCmd.AddCommand(bedrock_getAutomatedReasoningPolicyNextScenarioCmd)
}
