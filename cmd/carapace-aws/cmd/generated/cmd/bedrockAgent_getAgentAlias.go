package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_getAgentAliasCmd = &cobra.Command{
	Use:   "get-agent-alias",
	Short: "Gets information about an alias of an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_getAgentAliasCmd).Standalone()

	bedrockAgent_getAgentAliasCmd.Flags().String("agent-alias-id", "", "The unique identifier of the alias for which to get information.")
	bedrockAgent_getAgentAliasCmd.Flags().String("agent-id", "", "The unique identifier of the agent to which the alias to get information belongs.")
	bedrockAgent_getAgentAliasCmd.MarkFlagRequired("agent-alias-id")
	bedrockAgent_getAgentAliasCmd.MarkFlagRequired("agent-id")
	bedrockAgentCmd.AddCommand(bedrockAgent_getAgentAliasCmd)
}
