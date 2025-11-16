package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createGuardrailCmd = &cobra.Command{
	Use:   "create-guardrail",
	Short: "Creates a guardrail to block topics and to implement safeguards for your generative AI applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createGuardrailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_createGuardrailCmd).Standalone()

		bedrock_createGuardrailCmd.Flags().String("automated-reasoning-policy-config", "", "Optional configuration for integrating Automated Reasoning policies with the new guardrail.")
		bedrock_createGuardrailCmd.Flags().String("blocked-input-messaging", "", "The message to return when the guardrail blocks a prompt.")
		bedrock_createGuardrailCmd.Flags().String("blocked-outputs-messaging", "", "The message to return when the guardrail blocks a model response.")
		bedrock_createGuardrailCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than once.")
		bedrock_createGuardrailCmd.Flags().String("content-policy-config", "", "The content filter policies to configure for the guardrail.")
		bedrock_createGuardrailCmd.Flags().String("contextual-grounding-policy-config", "", "The contextual grounding policy configuration used to create a guardrail.")
		bedrock_createGuardrailCmd.Flags().String("cross-region-config", "", "The system-defined guardrail profile that you're using with your guardrail.")
		bedrock_createGuardrailCmd.Flags().String("description", "", "A description of the guardrail.")
		bedrock_createGuardrailCmd.Flags().String("kms-key-id", "", "The ARN of the KMS key that you use to encrypt the guardrail.")
		bedrock_createGuardrailCmd.Flags().String("name", "", "The name to give the guardrail.")
		bedrock_createGuardrailCmd.Flags().String("sensitive-information-policy-config", "", "The sensitive information policy to configure for the guardrail.")
		bedrock_createGuardrailCmd.Flags().String("tags", "", "The tags that you want to attach to the guardrail.")
		bedrock_createGuardrailCmd.Flags().String("topic-policy-config", "", "The topic policies to configure for the guardrail.")
		bedrock_createGuardrailCmd.Flags().String("word-policy-config", "", "The word policy you configure for the guardrail.")
		bedrock_createGuardrailCmd.MarkFlagRequired("blocked-input-messaging")
		bedrock_createGuardrailCmd.MarkFlagRequired("blocked-outputs-messaging")
		bedrock_createGuardrailCmd.MarkFlagRequired("name")
	})
	bedrockCmd.AddCommand(bedrock_createGuardrailCmd)
}
