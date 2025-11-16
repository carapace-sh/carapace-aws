package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_getFlowVersionCmd = &cobra.Command{
	Use:   "get-flow-version",
	Short: "Retrieves information about a version of a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_getFlowVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_getFlowVersionCmd).Standalone()

		bedrockAgent_getFlowVersionCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow for which to get information.")
		bedrockAgent_getFlowVersionCmd.Flags().String("flow-version", "", "The version of the flow for which to get information.")
		bedrockAgent_getFlowVersionCmd.MarkFlagRequired("flow-identifier")
		bedrockAgent_getFlowVersionCmd.MarkFlagRequired("flow-version")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_getFlowVersionCmd)
}
