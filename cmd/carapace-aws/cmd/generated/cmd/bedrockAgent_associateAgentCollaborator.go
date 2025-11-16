package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_associateAgentCollaboratorCmd = &cobra.Command{
	Use:   "associate-agent-collaborator",
	Short: "Makes an agent a collaborator for another agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_associateAgentCollaboratorCmd).Standalone()

	bedrockAgent_associateAgentCollaboratorCmd.Flags().String("agent-descriptor", "", "The alias of the collaborator agent.")
	bedrockAgent_associateAgentCollaboratorCmd.Flags().String("agent-id", "", "The agent's ID.")
	bedrockAgent_associateAgentCollaboratorCmd.Flags().String("agent-version", "", "An agent version.")
	bedrockAgent_associateAgentCollaboratorCmd.Flags().String("client-token", "", "A client token.")
	bedrockAgent_associateAgentCollaboratorCmd.Flags().String("collaboration-instruction", "", "Instruction for the collaborator.")
	bedrockAgent_associateAgentCollaboratorCmd.Flags().String("collaborator-name", "", "A name for the collaborator.")
	bedrockAgent_associateAgentCollaboratorCmd.Flags().String("relay-conversation-history", "", "A relay conversation history for the collaborator.")
	bedrockAgent_associateAgentCollaboratorCmd.MarkFlagRequired("agent-descriptor")
	bedrockAgent_associateAgentCollaboratorCmd.MarkFlagRequired("agent-id")
	bedrockAgent_associateAgentCollaboratorCmd.MarkFlagRequired("agent-version")
	bedrockAgent_associateAgentCollaboratorCmd.MarkFlagRequired("collaboration-instruction")
	bedrockAgent_associateAgentCollaboratorCmd.MarkFlagRequired("collaborator-name")
	bedrockAgentCmd.AddCommand(bedrockAgent_associateAgentCollaboratorCmd)
}
