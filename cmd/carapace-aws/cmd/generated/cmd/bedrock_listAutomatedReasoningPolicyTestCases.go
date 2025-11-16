package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listAutomatedReasoningPolicyTestCasesCmd = &cobra.Command{
	Use:   "list-automated-reasoning-policy-test-cases",
	Short: "Lists tests for an Automated Reasoning policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listAutomatedReasoningPolicyTestCasesCmd).Standalone()

	bedrock_listAutomatedReasoningPolicyTestCasesCmd.Flags().String("max-results", "", "The maximum number of tests to return in a single call.")
	bedrock_listAutomatedReasoningPolicyTestCasesCmd.Flags().String("next-token", "", "The pagination token from a previous request to retrieve the next page of results.")
	bedrock_listAutomatedReasoningPolicyTestCasesCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy for which to list tests.")
	bedrock_listAutomatedReasoningPolicyTestCasesCmd.MarkFlagRequired("policy-arn")
	bedrockCmd.AddCommand(bedrock_listAutomatedReasoningPolicyTestCasesCmd)
}
