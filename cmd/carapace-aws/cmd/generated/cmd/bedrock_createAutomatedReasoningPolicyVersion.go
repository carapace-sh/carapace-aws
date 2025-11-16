package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createAutomatedReasoningPolicyVersionCmd = &cobra.Command{
	Use:   "create-automated-reasoning-policy-version",
	Short: "Creates a new version of an existing Automated Reasoning policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createAutomatedReasoningPolicyVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_createAutomatedReasoningPolicyVersionCmd).Standalone()

		bedrock_createAutomatedReasoningPolicyVersionCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the operation completes no more than one time.")
		bedrock_createAutomatedReasoningPolicyVersionCmd.Flags().String("last-updated-definition-hash", "", "The hash of the current policy definition used as a concurrency token to ensure the policy hasn't been modified since you last retrieved it.")
		bedrock_createAutomatedReasoningPolicyVersionCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy for which to create a version.")
		bedrock_createAutomatedReasoningPolicyVersionCmd.Flags().String("tags", "", "A list of tags to associate with the policy version.")
		bedrock_createAutomatedReasoningPolicyVersionCmd.MarkFlagRequired("last-updated-definition-hash")
		bedrock_createAutomatedReasoningPolicyVersionCmd.MarkFlagRequired("policy-arn")
	})
	bedrockCmd.AddCommand(bedrock_createAutomatedReasoningPolicyVersionCmd)
}
