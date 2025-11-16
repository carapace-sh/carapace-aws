package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_deleteGuardrailCmd = &cobra.Command{
	Use:   "delete-guardrail",
	Short: "Deletes a guardrail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_deleteGuardrailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_deleteGuardrailCmd).Standalone()

		bedrock_deleteGuardrailCmd.Flags().String("guardrail-identifier", "", "The unique identifier of the guardrail.")
		bedrock_deleteGuardrailCmd.Flags().String("guardrail-version", "", "The version of the guardrail.")
		bedrock_deleteGuardrailCmd.MarkFlagRequired("guardrail-identifier")
	})
	bedrockCmd.AddCommand(bedrock_deleteGuardrailCmd)
}
