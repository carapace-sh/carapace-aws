package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_createFlowVersionCmd = &cobra.Command{
	Use:   "create-flow-version",
	Short: "Creates a version of the flow that you can deploy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_createFlowVersionCmd).Standalone()

	bedrockAgent_createFlowVersionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
	bedrockAgent_createFlowVersionCmd.Flags().String("description", "", "A description of the version of the flow.")
	bedrockAgent_createFlowVersionCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow that you want to create a version of.")
	bedrockAgent_createFlowVersionCmd.MarkFlagRequired("flow-identifier")
	bedrockAgentCmd.AddCommand(bedrockAgent_createFlowVersionCmd)
}
