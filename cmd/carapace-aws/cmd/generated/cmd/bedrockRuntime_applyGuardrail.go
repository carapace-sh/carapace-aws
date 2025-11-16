package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockRuntime_applyGuardrailCmd = &cobra.Command{
	Use:   "apply-guardrail",
	Short: "The action to apply a guardrail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockRuntime_applyGuardrailCmd).Standalone()

	bedrockRuntime_applyGuardrailCmd.Flags().String("content", "", "The content details used in the request to apply the guardrail.")
	bedrockRuntime_applyGuardrailCmd.Flags().String("guardrail-identifier", "", "The guardrail identifier used in the request to apply the guardrail.")
	bedrockRuntime_applyGuardrailCmd.Flags().String("guardrail-version", "", "The guardrail version used in the request to apply the guardrail.")
	bedrockRuntime_applyGuardrailCmd.Flags().String("output-scope", "", "Specifies the scope of the output that you get in the response.")
	bedrockRuntime_applyGuardrailCmd.Flags().String("source", "", "The source of data used in the request to apply the guardrail.")
	bedrockRuntime_applyGuardrailCmd.MarkFlagRequired("content")
	bedrockRuntime_applyGuardrailCmd.MarkFlagRequired("guardrail-identifier")
	bedrockRuntime_applyGuardrailCmd.MarkFlagRequired("guardrail-version")
	bedrockRuntime_applyGuardrailCmd.MarkFlagRequired("source")
	bedrockRuntimeCmd.AddCommand(bedrockRuntime_applyGuardrailCmd)
}
