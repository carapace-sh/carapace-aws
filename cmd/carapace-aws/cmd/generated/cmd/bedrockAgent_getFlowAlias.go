package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_getFlowAliasCmd = &cobra.Command{
	Use:   "get-flow-alias",
	Short: "Retrieves information about a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_getFlowAliasCmd).Standalone()

	bedrockAgent_getFlowAliasCmd.Flags().String("alias-identifier", "", "The unique identifier of the alias for which to retrieve information.")
	bedrockAgent_getFlowAliasCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow that the alias belongs to.")
	bedrockAgent_getFlowAliasCmd.MarkFlagRequired("alias-identifier")
	bedrockAgent_getFlowAliasCmd.MarkFlagRequired("flow-identifier")
	bedrockAgentCmd.AddCommand(bedrockAgent_getFlowAliasCmd)
}
