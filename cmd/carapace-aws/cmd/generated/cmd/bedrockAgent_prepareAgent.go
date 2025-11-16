package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_prepareAgentCmd = &cobra.Command{
	Use:   "prepare-agent",
	Short: "Creates a `DRAFT` version of the agent that can be used for internal testing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_prepareAgentCmd).Standalone()

	bedrockAgent_prepareAgentCmd.Flags().String("agent-id", "", "The unique identifier of the agent for which to create a `DRAFT` version.")
	bedrockAgent_prepareAgentCmd.MarkFlagRequired("agent-id")
	bedrockAgentCmd.AddCommand(bedrockAgent_prepareAgentCmd)
}
