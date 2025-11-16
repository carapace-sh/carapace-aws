package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_deleteFlowAliasCmd = &cobra.Command{
	Use:   "delete-flow-alias",
	Short: "Deletes an alias of a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_deleteFlowAliasCmd).Standalone()

	bedrockAgent_deleteFlowAliasCmd.Flags().String("alias-identifier", "", "The unique identifier of the alias to be deleted.")
	bedrockAgent_deleteFlowAliasCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow that the alias belongs to.")
	bedrockAgent_deleteFlowAliasCmd.MarkFlagRequired("alias-identifier")
	bedrockAgent_deleteFlowAliasCmd.MarkFlagRequired("flow-identifier")
	bedrockAgentCmd.AddCommand(bedrockAgent_deleteFlowAliasCmd)
}
