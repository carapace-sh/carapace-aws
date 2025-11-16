package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listAutomatedReasoningPolicyTestResultsCmd = &cobra.Command{
	Use:   "list-automated-reasoning-policy-test-results",
	Short: "Lists test results for an Automated Reasoning policy, showing how the policy performed against various test scenarios and validation checks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listAutomatedReasoningPolicyTestResultsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_listAutomatedReasoningPolicyTestResultsCmd).Standalone()

		bedrock_listAutomatedReasoningPolicyTestResultsCmd.Flags().String("build-workflow-id", "", "The unique identifier of the build workflow whose test results you want to list.")
		bedrock_listAutomatedReasoningPolicyTestResultsCmd.Flags().String("max-results", "", "The maximum number of test results to return in a single response.")
		bedrock_listAutomatedReasoningPolicyTestResultsCmd.Flags().String("next-token", "", "A pagination token from a previous request to continue listing test results from where the previous request left off.")
		bedrock_listAutomatedReasoningPolicyTestResultsCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy whose test results you want to list.")
		bedrock_listAutomatedReasoningPolicyTestResultsCmd.MarkFlagRequired("build-workflow-id")
		bedrock_listAutomatedReasoningPolicyTestResultsCmd.MarkFlagRequired("policy-arn")
	})
	bedrockCmd.AddCommand(bedrock_listAutomatedReasoningPolicyTestResultsCmd)
}
