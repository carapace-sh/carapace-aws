package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_deleteAgentCmd = &cobra.Command{
	Use:   "delete-agent",
	Short: "Deletes an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_deleteAgentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_deleteAgentCmd).Standalone()

		bedrockAgent_deleteAgentCmd.Flags().String("agent-id", "", "The unique identifier of the agent to delete.")
		bedrockAgent_deleteAgentCmd.Flags().Bool("no-skip-resource-in-use-check", false, "By default, this value is `false` and deletion is stopped if the resource is in use.")
		bedrockAgent_deleteAgentCmd.Flags().Bool("skip-resource-in-use-check", false, "By default, this value is `false` and deletion is stopped if the resource is in use.")
		bedrockAgent_deleteAgentCmd.MarkFlagRequired("agent-id")
		bedrockAgent_deleteAgentCmd.Flag("no-skip-resource-in-use-check").Hidden = true
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_deleteAgentCmd)
}
