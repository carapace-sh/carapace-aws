package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_deleteAutomatedReasoningPolicyTestCaseCmd = &cobra.Command{
	Use:   "delete-automated-reasoning-policy-test-case",
	Short: "Deletes an Automated Reasoning policy test.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_deleteAutomatedReasoningPolicyTestCaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_deleteAutomatedReasoningPolicyTestCaseCmd).Standalone()

		bedrock_deleteAutomatedReasoningPolicyTestCaseCmd.Flags().String("last-updated-at", "", "The timestamp when the test was last updated.")
		bedrock_deleteAutomatedReasoningPolicyTestCaseCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy that contains the test.")
		bedrock_deleteAutomatedReasoningPolicyTestCaseCmd.Flags().String("test-case-id", "", "The unique identifier of the test to delete.")
		bedrock_deleteAutomatedReasoningPolicyTestCaseCmd.MarkFlagRequired("last-updated-at")
		bedrock_deleteAutomatedReasoningPolicyTestCaseCmd.MarkFlagRequired("policy-arn")
		bedrock_deleteAutomatedReasoningPolicyTestCaseCmd.MarkFlagRequired("test-case-id")
	})
	bedrockCmd.AddCommand(bedrock_deleteAutomatedReasoningPolicyTestCaseCmd)
}
