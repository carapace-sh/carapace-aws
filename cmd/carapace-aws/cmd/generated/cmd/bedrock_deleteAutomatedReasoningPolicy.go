package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_deleteAutomatedReasoningPolicyCmd = &cobra.Command{
	Use:   "delete-automated-reasoning-policy",
	Short: "Deletes an Automated Reasoning policy or policy version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_deleteAutomatedReasoningPolicyCmd).Standalone()

	bedrock_deleteAutomatedReasoningPolicyCmd.Flags().Bool("force", false, "Specifies whether to force delete the automated reasoning policy even if it has active resources.")
	bedrock_deleteAutomatedReasoningPolicyCmd.Flags().Bool("no-force", false, "Specifies whether to force delete the automated reasoning policy even if it has active resources.")
	bedrock_deleteAutomatedReasoningPolicyCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy to delete.")
	bedrock_deleteAutomatedReasoningPolicyCmd.Flag("no-force").Hidden = true
	bedrock_deleteAutomatedReasoningPolicyCmd.MarkFlagRequired("policy-arn")
	bedrockCmd.AddCommand(bedrock_deleteAutomatedReasoningPolicyCmd)
}
