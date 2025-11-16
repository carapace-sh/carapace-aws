package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listGuardrailsCmd = &cobra.Command{
	Use:   "list-guardrails",
	Short: "Lists details about all the guardrails in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listGuardrailsCmd).Standalone()

	bedrock_listGuardrailsCmd.Flags().String("guardrail-identifier", "", "The unique identifier of the guardrail.")
	bedrock_listGuardrailsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	bedrock_listGuardrailsCmd.Flags().String("next-token", "", "If there are more results than were returned in the response, the response returns a `nextToken` that you can send in another `ListGuardrails` request to see the next batch of results.")
	bedrockCmd.AddCommand(bedrock_listGuardrailsCmd)
}
