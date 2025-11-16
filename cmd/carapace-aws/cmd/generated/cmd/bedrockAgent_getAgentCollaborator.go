package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_getAgentCollaboratorCmd = &cobra.Command{
	Use:   "get-agent-collaborator",
	Short: "Retrieves information about an agent's collaborator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_getAgentCollaboratorCmd).Standalone()

	bedrockAgent_getAgentCollaboratorCmd.Flags().String("agent-id", "", "The agent's ID.")
	bedrockAgent_getAgentCollaboratorCmd.Flags().String("agent-version", "", "The agent's version.")
	bedrockAgent_getAgentCollaboratorCmd.Flags().String("collaborator-id", "", "The collaborator's ID.")
	bedrockAgent_getAgentCollaboratorCmd.MarkFlagRequired("agent-id")
	bedrockAgent_getAgentCollaboratorCmd.MarkFlagRequired("agent-version")
	bedrockAgent_getAgentCollaboratorCmd.MarkFlagRequired("collaborator-id")
	bedrockAgentCmd.AddCommand(bedrockAgent_getAgentCollaboratorCmd)
}
