package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_deleteFlowCmd = &cobra.Command{
	Use:   "delete-flow",
	Short: "Deletes a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_deleteFlowCmd).Standalone()

	bedrockAgent_deleteFlowCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow.")
	bedrockAgent_deleteFlowCmd.Flags().Bool("no-skip-resource-in-use-check", false, "By default, this value is `false` and deletion is stopped if the resource is in use.")
	bedrockAgent_deleteFlowCmd.Flags().Bool("skip-resource-in-use-check", false, "By default, this value is `false` and deletion is stopped if the resource is in use.")
	bedrockAgent_deleteFlowCmd.MarkFlagRequired("flow-identifier")
	bedrockAgent_deleteFlowCmd.Flag("no-skip-resource-in-use-check").Hidden = true
	bedrockAgentCmd.AddCommand(bedrockAgent_deleteFlowCmd)
}
