package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_deleteAgentAliasCmd = &cobra.Command{
	Use:   "delete-agent-alias",
	Short: "Deletes an alias of an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_deleteAgentAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_deleteAgentAliasCmd).Standalone()

		bedrockAgent_deleteAgentAliasCmd.Flags().String("agent-alias-id", "", "The unique identifier of the alias to delete.")
		bedrockAgent_deleteAgentAliasCmd.Flags().String("agent-id", "", "The unique identifier of the agent that the alias belongs to.")
		bedrockAgent_deleteAgentAliasCmd.MarkFlagRequired("agent-alias-id")
		bedrockAgent_deleteAgentAliasCmd.MarkFlagRequired("agent-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_deleteAgentAliasCmd)
}
