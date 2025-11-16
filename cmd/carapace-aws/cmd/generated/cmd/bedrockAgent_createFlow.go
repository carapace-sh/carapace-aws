package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_createFlowCmd = &cobra.Command{
	Use:   "create-flow",
	Short: "Creates a prompt flow that you can use to send an input through various steps to yield an output.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_createFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_createFlowCmd).Standalone()

		bedrockAgent_createFlowCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
		bedrockAgent_createFlowCmd.Flags().String("customer-encryption-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key to encrypt the flow.")
		bedrockAgent_createFlowCmd.Flags().String("definition", "", "A definition of the nodes and connections between nodes in the flow.")
		bedrockAgent_createFlowCmd.Flags().String("description", "", "A description for the flow.")
		bedrockAgent_createFlowCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the service role with permissions to create and manage a flow.")
		bedrockAgent_createFlowCmd.Flags().String("name", "", "A name for the flow.")
		bedrockAgent_createFlowCmd.Flags().String("tags", "", "Any tags that you want to attach to the flow.")
		bedrockAgent_createFlowCmd.MarkFlagRequired("execution-role-arn")
		bedrockAgent_createFlowCmd.MarkFlagRequired("name")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_createFlowCmd)
}
