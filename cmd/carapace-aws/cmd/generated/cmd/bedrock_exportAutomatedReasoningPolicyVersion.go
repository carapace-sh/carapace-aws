package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_exportAutomatedReasoningPolicyVersionCmd = &cobra.Command{
	Use:   "export-automated-reasoning-policy-version",
	Short: "Exports the policy definition for an Automated Reasoning policy version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_exportAutomatedReasoningPolicyVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_exportAutomatedReasoningPolicyVersionCmd).Standalone()

		bedrock_exportAutomatedReasoningPolicyVersionCmd.Flags().String("policy-arn", "", "The Amazon Resource Name (ARN) of the Automated Reasoning policy to export.")
		bedrock_exportAutomatedReasoningPolicyVersionCmd.MarkFlagRequired("policy-arn")
	})
	bedrockCmd.AddCommand(bedrock_exportAutomatedReasoningPolicyVersionCmd)
}
