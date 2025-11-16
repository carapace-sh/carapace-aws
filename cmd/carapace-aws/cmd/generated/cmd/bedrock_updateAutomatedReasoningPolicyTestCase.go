package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_updateAutomatedReasoningPolicyTestCaseCmd = &cobra.Command{
	Use:   "update-automated-reasoning-policy-test-case",
	Short: "Updates an existing Automated Reasoning policy test.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_updateAutomatedReasoningPolicyTestCaseCmd).Standalone()

	bedrock_updateAutomatedReasoningPolicyTestCaseCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the operation completes no more than one time.")
	bedrock_updateAutomatedReasoningPolicyTestCaseCmd.Flags().String("confidence-threshold", "", "The updated minimum confidence level for logic validation.")
	bedrock_updateAutomatedReasoningPolicyTestCaseCmd.Flags().String("expected-aggregated-findings-result", "", "The updated expected result of the Automated Reasoning check.")
	bedrock_updateAutomatedReasoningPolicyTestCaseCmd.Flags().String("guard-content", "", "The updated content to be validated by the Automated Reasoning policy.")
	bedrock_updateAutomatedReasoningPolicyTestCaseCmd.Flags().String("last-updated-at", "", "The timestamp when the test was last updated.")
	bedrock_updateAutomatedReasoningPolicyTestCaseCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy that contains the test.")
	bedrock_updateAutomatedReasoningPolicyTestCaseCmd.Flags().String("query-content", "", "The updated input query or prompt that generated the content.")
	bedrock_updateAutomatedReasoningPolicyTestCaseCmd.Flags().String("test-case-id", "", "The unique identifier of the test to update.")
	bedrock_updateAutomatedReasoningPolicyTestCaseCmd.MarkFlagRequired("expected-aggregated-findings-result")
	bedrock_updateAutomatedReasoningPolicyTestCaseCmd.MarkFlagRequired("guard-content")
	bedrock_updateAutomatedReasoningPolicyTestCaseCmd.MarkFlagRequired("last-updated-at")
	bedrock_updateAutomatedReasoningPolicyTestCaseCmd.MarkFlagRequired("policy-arn")
	bedrock_updateAutomatedReasoningPolicyTestCaseCmd.MarkFlagRequired("test-case-id")
	bedrockCmd.AddCommand(bedrock_updateAutomatedReasoningPolicyTestCaseCmd)
}
