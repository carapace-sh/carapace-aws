package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_createAgentAliasCmd = &cobra.Command{
	Use:   "create-agent-alias",
	Short: "Creates an alias of an agent that can be used to deploy the agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_createAgentAliasCmd).Standalone()

	bedrockAgent_createAgentAliasCmd.Flags().String("agent-alias-name", "", "The name of the alias.")
	bedrockAgent_createAgentAliasCmd.Flags().String("agent-id", "", "The unique identifier of the agent.")
	bedrockAgent_createAgentAliasCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
	bedrockAgent_createAgentAliasCmd.Flags().String("description", "", "A description of the alias of the agent.")
	bedrockAgent_createAgentAliasCmd.Flags().String("routing-configuration", "", "Contains details about the routing configuration of the alias.")
	bedrockAgent_createAgentAliasCmd.Flags().String("tags", "", "Any tags that you want to attach to the alias of the agent.")
	bedrockAgent_createAgentAliasCmd.MarkFlagRequired("agent-alias-name")
	bedrockAgent_createAgentAliasCmd.MarkFlagRequired("agent-id")
	bedrockAgentCmd.AddCommand(bedrockAgent_createAgentAliasCmd)
}
