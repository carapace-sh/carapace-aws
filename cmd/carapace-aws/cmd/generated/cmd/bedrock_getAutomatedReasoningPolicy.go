package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getAutomatedReasoningPolicyCmd = &cobra.Command{
	Use:   "get-automated-reasoning-policy",
	Short: "Retrieves details about an Automated Reasoning policy or policy version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getAutomatedReasoningPolicyCmd).Standalone()

	bedrock_getAutomatedReasoningPolicyCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy to retrieve.")
	bedrock_getAutomatedReasoningPolicyCmd.MarkFlagRequired("policy-arn")
	bedrockCmd.AddCommand(bedrock_getAutomatedReasoningPolicyCmd)
}
