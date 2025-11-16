package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_updateFlowCmd = &cobra.Command{
	Use:   "update-flow",
	Short: "Modifies a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_updateFlowCmd).Standalone()

	bedrockAgent_updateFlowCmd.Flags().String("customer-encryption-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key to encrypt the flow.")
	bedrockAgent_updateFlowCmd.Flags().String("definition", "", "A definition of the nodes and the connections between the nodes in the flow.")
	bedrockAgent_updateFlowCmd.Flags().String("description", "", "A description for the flow.")
	bedrockAgent_updateFlowCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the service role with permissions to create and manage a flow.")
	bedrockAgent_updateFlowCmd.Flags().String("flow-identifier", "", "The unique identifier of the flow.")
	bedrockAgent_updateFlowCmd.Flags().String("name", "", "A name for the flow.")
	bedrockAgent_updateFlowCmd.MarkFlagRequired("execution-role-arn")
	bedrockAgent_updateFlowCmd.MarkFlagRequired("flow-identifier")
	bedrockAgent_updateFlowCmd.MarkFlagRequired("name")
	bedrockAgentCmd.AddCommand(bedrockAgent_updateFlowCmd)
}
