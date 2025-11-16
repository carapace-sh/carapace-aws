package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getAutomatedReasoningPolicyTestResultCmd = &cobra.Command{
	Use:   "get-automated-reasoning-policy-test-result",
	Short: "Retrieves the test result for a specific Automated Reasoning policy test.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getAutomatedReasoningPolicyTestResultCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_getAutomatedReasoningPolicyTestResultCmd).Standalone()

		bedrock_getAutomatedReasoningPolicyTestResultCmd.Flags().String("build-workflow-id", "", "The build workflow identifier.")
		bedrock_getAutomatedReasoningPolicyTestResultCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy.")
		bedrock_getAutomatedReasoningPolicyTestResultCmd.Flags().String("test-case-id", "", "The unique identifier of the test for which to retrieve results.")
		bedrock_getAutomatedReasoningPolicyTestResultCmd.MarkFlagRequired("build-workflow-id")
		bedrock_getAutomatedReasoningPolicyTestResultCmd.MarkFlagRequired("policy-arn")
		bedrock_getAutomatedReasoningPolicyTestResultCmd.MarkFlagRequired("test-case-id")
	})
	bedrockCmd.AddCommand(bedrock_getAutomatedReasoningPolicyTestResultCmd)
}
