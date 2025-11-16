package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_updateAgentAliasCmd = &cobra.Command{
	Use:   "update-agent-alias",
	Short: "Updates configurations for an alias of an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_updateAgentAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_updateAgentAliasCmd).Standalone()

		bedrockAgent_updateAgentAliasCmd.Flags().String("agent-alias-id", "", "The unique identifier of the alias.")
		bedrockAgent_updateAgentAliasCmd.Flags().String("agent-alias-name", "", "Specifies a new name for the alias.")
		bedrockAgent_updateAgentAliasCmd.Flags().String("agent-id", "", "The unique identifier of the agent.")
		bedrockAgent_updateAgentAliasCmd.Flags().String("alias-invocation-state", "", "The invocation state for the agent alias.")
		bedrockAgent_updateAgentAliasCmd.Flags().String("description", "", "Specifies a new description for the alias.")
		bedrockAgent_updateAgentAliasCmd.Flags().String("routing-configuration", "", "Contains details about the routing configuration of the alias.")
		bedrockAgent_updateAgentAliasCmd.MarkFlagRequired("agent-alias-id")
		bedrockAgent_updateAgentAliasCmd.MarkFlagRequired("agent-alias-name")
		bedrockAgent_updateAgentAliasCmd.MarkFlagRequired("agent-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_updateAgentAliasCmd)
}
