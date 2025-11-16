package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createAutomatedReasoningPolicyTestCaseCmd = &cobra.Command{
	Use:   "create-automated-reasoning-policy-test-case",
	Short: "Creates a test for an Automated Reasoning policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createAutomatedReasoningPolicyTestCaseCmd).Standalone()

	bedrock_createAutomatedReasoningPolicyTestCaseCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the operation completes no more than one time.")
	bedrock_createAutomatedReasoningPolicyTestCaseCmd.Flags().String("confidence-threshold", "", "The minimum confidence level for logic validation.")
	bedrock_createAutomatedReasoningPolicyTestCaseCmd.Flags().String("expected-aggregated-findings-result", "", "The expected result of the Automated Reasoning check.")
	bedrock_createAutomatedReasoningPolicyTestCaseCmd.Flags().String("guard-content", "", "The output content that's validated by the Automated Reasoning policy.")
	bedrock_createAutomatedReasoningPolicyTestCaseCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy for which to create the test.")
	bedrock_createAutomatedReasoningPolicyTestCaseCmd.Flags().String("query-content", "", "The input query or prompt that generated the content.")
	bedrock_createAutomatedReasoningPolicyTestCaseCmd.MarkFlagRequired("expected-aggregated-findings-result")
	bedrock_createAutomatedReasoningPolicyTestCaseCmd.MarkFlagRequired("guard-content")
	bedrock_createAutomatedReasoningPolicyTestCaseCmd.MarkFlagRequired("policy-arn")
	bedrockCmd.AddCommand(bedrock_createAutomatedReasoningPolicyTestCaseCmd)
}
