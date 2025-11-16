package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_updateAutomatedReasoningPolicyCmd = &cobra.Command{
	Use:   "update-automated-reasoning-policy",
	Short: "Updates an existing Automated Reasoning policy with new rules, variables, or configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_updateAutomatedReasoningPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_updateAutomatedReasoningPolicyCmd).Standalone()

		bedrock_updateAutomatedReasoningPolicyCmd.Flags().String("description", "", "The updated description for the Automated Reasoning policy.")
		bedrock_updateAutomatedReasoningPolicyCmd.Flags().String("name", "", "The updated name for the Automated Reasoning policy.")
		bedrock_updateAutomatedReasoningPolicyCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy to update.")
		bedrock_updateAutomatedReasoningPolicyCmd.Flags().String("policy-definition", "", "The updated policy definition containing the formal logic rules, variables, and types.")
		bedrock_updateAutomatedReasoningPolicyCmd.MarkFlagRequired("policy-arn")
		bedrock_updateAutomatedReasoningPolicyCmd.MarkFlagRequired("policy-definition")
	})
	bedrockCmd.AddCommand(bedrock_updateAutomatedReasoningPolicyCmd)
}
