package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createAutomatedReasoningPolicyCmd = &cobra.Command{
	Use:   "create-automated-reasoning-policy",
	Short: "Creates an Automated Reasoning policy for Amazon Bedrock Guardrails.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createAutomatedReasoningPolicyCmd).Standalone()

	bedrock_createAutomatedReasoningPolicyCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the operation completes no more than once.")
	bedrock_createAutomatedReasoningPolicyCmd.Flags().String("description", "", "A description of the Automated Reasoning policy.")
	bedrock_createAutomatedReasoningPolicyCmd.Flags().String("kms-key-id", "", "The identifier of the KMS key to use for encrypting the automated reasoning policy and its associated artifacts.")
	bedrock_createAutomatedReasoningPolicyCmd.Flags().String("name", "", "A unique name for the Automated Reasoning policy.")
	bedrock_createAutomatedReasoningPolicyCmd.Flags().String("policy-definition", "", "The policy definition that contains the formal logic rules, variables, and custom variable types used to validate foundation model responses in your application.")
	bedrock_createAutomatedReasoningPolicyCmd.Flags().String("tags", "", "A list of tags to associate with the Automated Reasoning policy.")
	bedrock_createAutomatedReasoningPolicyCmd.MarkFlagRequired("name")
	bedrockCmd.AddCommand(bedrock_createAutomatedReasoningPolicyCmd)
}
