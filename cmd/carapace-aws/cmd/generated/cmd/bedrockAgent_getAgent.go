package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_getAgentCmd = &cobra.Command{
	Use:   "get-agent",
	Short: "Gets information about an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_getAgentCmd).Standalone()

	bedrockAgent_getAgentCmd.Flags().String("agent-id", "", "The unique identifier of the agent.")
	bedrockAgent_getAgentCmd.MarkFlagRequired("agent-id")
	bedrockAgentCmd.AddCommand(bedrockAgent_getAgentCmd)
}
