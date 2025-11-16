package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getAutomatedReasoningPolicyTestCaseCmd = &cobra.Command{
	Use:   "get-automated-reasoning-policy-test-case",
	Short: "Retrieves details about a specific Automated Reasoning policy test.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getAutomatedReasoningPolicyTestCaseCmd).Standalone()

	bedrock_getAutomatedReasoningPolicyTestCaseCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy that contains the test.")
	bedrock_getAutomatedReasoningPolicyTestCaseCmd.Flags().String("test-case-id", "", "The unique identifier of the test to retrieve.")
	bedrock_getAutomatedReasoningPolicyTestCaseCmd.MarkFlagRequired("policy-arn")
	bedrock_getAutomatedReasoningPolicyTestCaseCmd.MarkFlagRequired("test-case-id")
	bedrockCmd.AddCommand(bedrock_getAutomatedReasoningPolicyTestCaseCmd)
}
