package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getGuardrailCmd = &cobra.Command{
	Use:   "get-guardrail",
	Short: "Gets details about a guardrail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getGuardrailCmd).Standalone()

	bedrock_getGuardrailCmd.Flags().String("guardrail-identifier", "", "The unique identifier of the guardrail for which to get details.")
	bedrock_getGuardrailCmd.Flags().String("guardrail-version", "", "The version of the guardrail for which to get details.")
	bedrock_getGuardrailCmd.MarkFlagRequired("guardrail-identifier")
	bedrockCmd.AddCommand(bedrock_getGuardrailCmd)
}
