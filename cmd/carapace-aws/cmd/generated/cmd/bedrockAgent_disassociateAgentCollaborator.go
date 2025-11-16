package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_disassociateAgentCollaboratorCmd = &cobra.Command{
	Use:   "disassociate-agent-collaborator",
	Short: "Disassociates an agent collaborator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_disassociateAgentCollaboratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_disassociateAgentCollaboratorCmd).Standalone()

		bedrockAgent_disassociateAgentCollaboratorCmd.Flags().String("agent-id", "", "An agent ID.")
		bedrockAgent_disassociateAgentCollaboratorCmd.Flags().String("agent-version", "", "The agent's version.")
		bedrockAgent_disassociateAgentCollaboratorCmd.Flags().String("collaborator-id", "", "The collaborator's ID.")
		bedrockAgent_disassociateAgentCollaboratorCmd.MarkFlagRequired("agent-id")
		bedrockAgent_disassociateAgentCollaboratorCmd.MarkFlagRequired("agent-version")
		bedrockAgent_disassociateAgentCollaboratorCmd.MarkFlagRequired("collaborator-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_disassociateAgentCollaboratorCmd)
}
