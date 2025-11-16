package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_validateFlowDefinitionCmd = &cobra.Command{
	Use:   "validate-flow-definition",
	Short: "Validates the definition of a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_validateFlowDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_validateFlowDefinitionCmd).Standalone()

		bedrockAgent_validateFlowDefinitionCmd.Flags().String("definition", "", "The definition of a flow to validate.")
		bedrockAgent_validateFlowDefinitionCmd.MarkFlagRequired("definition")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_validateFlowDefinitionCmd)
}
