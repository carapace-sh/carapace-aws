package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_getAgentVersionCmd = &cobra.Command{
	Use:   "get-agent-version",
	Short: "Gets details about a version of an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_getAgentVersionCmd).Standalone()

	bedrockAgent_getAgentVersionCmd.Flags().String("agent-id", "", "The unique identifier of the agent.")
	bedrockAgent_getAgentVersionCmd.Flags().String("agent-version", "", "The version of the agent.")
	bedrockAgent_getAgentVersionCmd.MarkFlagRequired("agent-id")
	bedrockAgent_getAgentVersionCmd.MarkFlagRequired("agent-version")
	bedrockAgentCmd.AddCommand(bedrockAgent_getAgentVersionCmd)
}
