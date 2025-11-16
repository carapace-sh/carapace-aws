package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_updateGuardrailCmd = &cobra.Command{
	Use:   "update-guardrail",
	Short: "Updates a guardrail with the values you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_updateGuardrailCmd).Standalone()

	bedrock_updateGuardrailCmd.Flags().String("automated-reasoning-policy-config", "", "Updated configuration for Automated Reasoning policies associated with the guardrail.")
	bedrock_updateGuardrailCmd.Flags().String("blocked-input-messaging", "", "The message to return when the guardrail blocks a prompt.")
	bedrock_updateGuardrailCmd.Flags().String("blocked-outputs-messaging", "", "The message to return when the guardrail blocks a model response.")
	bedrock_updateGuardrailCmd.Flags().String("content-policy-config", "", "The content policy to configure for the guardrail.")
	bedrock_updateGuardrailCmd.Flags().String("contextual-grounding-policy-config", "", "The contextual grounding policy configuration used to update a guardrail.")
	bedrock_updateGuardrailCmd.Flags().String("cross-region-config", "", "The system-defined guardrail profile that you're using with your guardrail.")
	bedrock_updateGuardrailCmd.Flags().String("description", "", "A description of the guardrail.")
	bedrock_updateGuardrailCmd.Flags().String("guardrail-identifier", "", "The unique identifier of the guardrail.")
	bedrock_updateGuardrailCmd.Flags().String("kms-key-id", "", "The ARN of the KMS key with which to encrypt the guardrail.")
	bedrock_updateGuardrailCmd.Flags().String("name", "", "A name for the guardrail.")
	bedrock_updateGuardrailCmd.Flags().String("sensitive-information-policy-config", "", "The sensitive information policy to configure for the guardrail.")
	bedrock_updateGuardrailCmd.Flags().String("topic-policy-config", "", "The topic policy to configure for the guardrail.")
	bedrock_updateGuardrailCmd.Flags().String("word-policy-config", "", "The word policy to configure for the guardrail.")
	bedrock_updateGuardrailCmd.MarkFlagRequired("blocked-input-messaging")
	bedrock_updateGuardrailCmd.MarkFlagRequired("blocked-outputs-messaging")
	bedrock_updateGuardrailCmd.MarkFlagRequired("guardrail-identifier")
	bedrock_updateGuardrailCmd.MarkFlagRequired("name")
	bedrockCmd.AddCommand(bedrock_updateGuardrailCmd)
}
