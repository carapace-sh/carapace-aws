package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_deleteFlowVersionCmd = &cobra.Command{
	Use:   "delete-flow-version",
	Short: "Deletes a version of a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_deleteFlowVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_deleteFlowVersionCmd).Standalone()

		bedrockAgent_deleteFlowVersionCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow whose version that you want to delete")
		bedrockAgent_deleteFlowVersionCmd.Flags().String("flow-version", "", "The version of the flow that you want to delete.")
		bedrockAgent_deleteFlowVersionCmd.Flags().Bool("no-skip-resource-in-use-check", false, "By default, this value is `false` and deletion is stopped if the resource is in use.")
		bedrockAgent_deleteFlowVersionCmd.Flags().Bool("skip-resource-in-use-check", false, "By default, this value is `false` and deletion is stopped if the resource is in use.")
		bedrockAgent_deleteFlowVersionCmd.MarkFlagRequired("flow-identifier")
		bedrockAgent_deleteFlowVersionCmd.MarkFlagRequired("flow-version")
		bedrockAgent_deleteFlowVersionCmd.Flag("no-skip-resource-in-use-check").Hidden = true
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_deleteFlowVersionCmd)
}
