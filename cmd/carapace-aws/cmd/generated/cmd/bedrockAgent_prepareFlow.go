package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_prepareFlowCmd = &cobra.Command{
	Use:   "prepare-flow",
	Short: "Prepares the `DRAFT` version of a flow so that it can be invoked.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_prepareFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_prepareFlowCmd).Standalone()

		bedrockAgent_prepareFlowCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow.")
		bedrockAgent_prepareFlowCmd.MarkFlagRequired("flow-identifier")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_prepareFlowCmd)
}
