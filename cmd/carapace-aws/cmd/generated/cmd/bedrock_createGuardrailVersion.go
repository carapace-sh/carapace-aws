package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createGuardrailVersionCmd = &cobra.Command{
	Use:   "create-guardrail-version",
	Short: "Creates a version of the guardrail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createGuardrailVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_createGuardrailVersionCmd).Standalone()

		bedrock_createGuardrailVersionCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than once.")
		bedrock_createGuardrailVersionCmd.Flags().String("description", "", "A description of the guardrail version.")
		bedrock_createGuardrailVersionCmd.Flags().String("guardrail-identifier", "", "The unique identifier of the guardrail.")
		bedrock_createGuardrailVersionCmd.MarkFlagRequired("guardrail-identifier")
	})
	bedrockCmd.AddCommand(bedrock_createGuardrailVersionCmd)
}
