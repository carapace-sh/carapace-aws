package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listAutomatedReasoningPoliciesCmd = &cobra.Command{
	Use:   "list-automated-reasoning-policies",
	Short: "Lists all Automated Reasoning policies in your account, with optional filtering by policy ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listAutomatedReasoningPoliciesCmd).Standalone()

	bedrock_listAutomatedReasoningPoliciesCmd.Flags().String("max-results", "", "The maximum number of policies to return in a single call.")
	bedrock_listAutomatedReasoningPoliciesCmd.Flags().String("next-token", "", "The pagination token from a previous request to retrieve the next page of results.")
	bedrock_listAutomatedReasoningPoliciesCmd.Flags().String("policy-arn", "", "Optional filter to list only the policy versions with the specified Amazon Resource Name (ARN).")
	bedrockCmd.AddCommand(bedrock_listAutomatedReasoningPoliciesCmd)
}
