package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_updateFlowAliasCmd = &cobra.Command{
	Use:   "update-flow-alias",
	Short: "Modifies the alias of a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_updateFlowAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_updateFlowAliasCmd).Standalone()

		bedrockAgent_updateFlowAliasCmd.Flags().String("alias-identifier", "", "The unique identifier of the alias.")
		bedrockAgent_updateFlowAliasCmd.Flags().String("concurrency-configuration", "", "The configuration that specifies how nodes in the flow are executed in parallel.")
		bedrockAgent_updateFlowAliasCmd.Flags().String("description", "", "A description for the alias.")
		bedrockAgent_updateFlowAliasCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow.")
		bedrockAgent_updateFlowAliasCmd.Flags().String("name", "", "The name of the alias.")
		bedrockAgent_updateFlowAliasCmd.Flags().String("routing-configuration", "", "Contains information about the version to which to map the alias.")
		bedrockAgent_updateFlowAliasCmd.MarkFlagRequired("alias-identifier")
		bedrockAgent_updateFlowAliasCmd.MarkFlagRequired("flow-identifier")
		bedrockAgent_updateFlowAliasCmd.MarkFlagRequired("name")
		bedrockAgent_updateFlowAliasCmd.MarkFlagRequired("routing-configuration")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_updateFlowAliasCmd)
}
