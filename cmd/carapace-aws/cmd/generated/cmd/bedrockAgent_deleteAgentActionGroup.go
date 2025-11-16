package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_deleteAgentActionGroupCmd = &cobra.Command{
	Use:   "delete-agent-action-group",
	Short: "Deletes an action group in an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_deleteAgentActionGroupCmd).Standalone()

	bedrockAgent_deleteAgentActionGroupCmd.Flags().String("action-group-id", "", "The unique identifier of the action group to delete.")
	bedrockAgent_deleteAgentActionGroupCmd.Flags().String("agent-id", "", "The unique identifier of the agent that the action group belongs to.")
	bedrockAgent_deleteAgentActionGroupCmd.Flags().String("agent-version", "", "The version of the agent that the action group belongs to.")
	bedrockAgent_deleteAgentActionGroupCmd.Flags().Bool("no-skip-resource-in-use-check", false, "By default, this value is `false` and deletion is stopped if the resource is in use.")
	bedrockAgent_deleteAgentActionGroupCmd.Flags().Bool("skip-resource-in-use-check", false, "By default, this value is `false` and deletion is stopped if the resource is in use.")
	bedrockAgent_deleteAgentActionGroupCmd.MarkFlagRequired("action-group-id")
	bedrockAgent_deleteAgentActionGroupCmd.MarkFlagRequired("agent-id")
	bedrockAgent_deleteAgentActionGroupCmd.MarkFlagRequired("agent-version")
	bedrockAgent_deleteAgentActionGroupCmd.Flag("no-skip-resource-in-use-check").Hidden = true
	bedrockAgentCmd.AddCommand(bedrockAgent_deleteAgentActionGroupCmd)
}
