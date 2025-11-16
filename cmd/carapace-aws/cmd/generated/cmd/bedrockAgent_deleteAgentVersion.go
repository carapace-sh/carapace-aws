package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_deleteAgentVersionCmd = &cobra.Command{
	Use:   "delete-agent-version",
	Short: "Deletes a version of an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_deleteAgentVersionCmd).Standalone()

	bedrockAgent_deleteAgentVersionCmd.Flags().String("agent-id", "", "The unique identifier of the agent that the version belongs to.")
	bedrockAgent_deleteAgentVersionCmd.Flags().String("agent-version", "", "The version of the agent to delete.")
	bedrockAgent_deleteAgentVersionCmd.Flags().Bool("no-skip-resource-in-use-check", false, "By default, this value is `false` and deletion is stopped if the resource is in use.")
	bedrockAgent_deleteAgentVersionCmd.Flags().Bool("skip-resource-in-use-check", false, "By default, this value is `false` and deletion is stopped if the resource is in use.")
	bedrockAgent_deleteAgentVersionCmd.MarkFlagRequired("agent-id")
	bedrockAgent_deleteAgentVersionCmd.MarkFlagRequired("agent-version")
	bedrockAgent_deleteAgentVersionCmd.Flag("no-skip-resource-in-use-check").Hidden = true
	bedrockAgentCmd.AddCommand(bedrockAgent_deleteAgentVersionCmd)
}
