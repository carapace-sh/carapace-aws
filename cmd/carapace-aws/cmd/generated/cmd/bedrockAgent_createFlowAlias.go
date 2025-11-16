package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_createFlowAliasCmd = &cobra.Command{
	Use:   "create-flow-alias",
	Short: "Creates an alias of a flow for deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_createFlowAliasCmd).Standalone()

	bedrockAgent_createFlowAliasCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
	bedrockAgent_createFlowAliasCmd.Flags().String("concurrency-configuration", "", "The configuration that specifies how nodes in the flow are executed in parallel.")
	bedrockAgent_createFlowAliasCmd.Flags().String("description", "", "A description for the alias.")
	bedrockAgent_createFlowAliasCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow for which to create an alias.")
	bedrockAgent_createFlowAliasCmd.Flags().String("name", "", "A name for the alias.")
	bedrockAgent_createFlowAliasCmd.Flags().String("routing-configuration", "", "Contains information about the version to which to map the alias.")
	bedrockAgent_createFlowAliasCmd.Flags().String("tags", "", "Any tags that you want to attach to the alias of the flow.")
	bedrockAgent_createFlowAliasCmd.MarkFlagRequired("flow-identifier")
	bedrockAgent_createFlowAliasCmd.MarkFlagRequired("name")
	bedrockAgent_createFlowAliasCmd.MarkFlagRequired("routing-configuration")
	bedrockAgentCmd.AddCommand(bedrockAgent_createFlowAliasCmd)
}
