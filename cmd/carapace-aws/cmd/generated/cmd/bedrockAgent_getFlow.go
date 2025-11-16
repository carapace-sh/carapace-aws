package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_getFlowCmd = &cobra.Command{
	Use:   "get-flow",
	Short: "Retrieves information about a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_getFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_getFlowCmd).Standalone()

		bedrockAgent_getFlowCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow.")
		bedrockAgent_getFlowCmd.MarkFlagRequired("flow-identifier")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_getFlowCmd)
}
