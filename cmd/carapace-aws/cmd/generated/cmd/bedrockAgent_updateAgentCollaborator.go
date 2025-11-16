package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_updateAgentCollaboratorCmd = &cobra.Command{
	Use:   "update-agent-collaborator",
	Short: "Updates an agent's collaborator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_updateAgentCollaboratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_updateAgentCollaboratorCmd).Standalone()

		bedrockAgent_updateAgentCollaboratorCmd.Flags().String("agent-descriptor", "", "An agent descriptor for the agent collaborator.")
		bedrockAgent_updateAgentCollaboratorCmd.Flags().String("agent-id", "", "The agent's ID.")
		bedrockAgent_updateAgentCollaboratorCmd.Flags().String("agent-version", "", "The agent's version.")
		bedrockAgent_updateAgentCollaboratorCmd.Flags().String("collaboration-instruction", "", "Instruction for the collaborator.")
		bedrockAgent_updateAgentCollaboratorCmd.Flags().String("collaborator-id", "", "The collaborator's ID.")
		bedrockAgent_updateAgentCollaboratorCmd.Flags().String("collaborator-name", "", "The collaborator's name.")
		bedrockAgent_updateAgentCollaboratorCmd.Flags().String("relay-conversation-history", "", "A relay conversation history for the collaborator.")
		bedrockAgent_updateAgentCollaboratorCmd.MarkFlagRequired("agent-descriptor")
		bedrockAgent_updateAgentCollaboratorCmd.MarkFlagRequired("agent-id")
		bedrockAgent_updateAgentCollaboratorCmd.MarkFlagRequired("agent-version")
		bedrockAgent_updateAgentCollaboratorCmd.MarkFlagRequired("collaboration-instruction")
		bedrockAgent_updateAgentCollaboratorCmd.MarkFlagRequired("collaborator-id")
		bedrockAgent_updateAgentCollaboratorCmd.MarkFlagRequired("collaborator-name")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_updateAgentCollaboratorCmd)
}
