package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_createCodeInterpreterCmd = &cobra.Command{
	Use:   "create-code-interpreter",
	Short: "Creates a custom code interpreter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_createCodeInterpreterCmd).Standalone()

	bedrockAgentcoreControl_createCodeInterpreterCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the operation completes no more than one time.")
	bedrockAgentcoreControl_createCodeInterpreterCmd.Flags().String("description", "", "The description of the code interpreter.")
	bedrockAgentcoreControl_createCodeInterpreterCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that provides permissions for the code interpreter to access Amazon Web Services services.")
	bedrockAgentcoreControl_createCodeInterpreterCmd.Flags().String("name", "", "The name of the code interpreter.")
	bedrockAgentcoreControl_createCodeInterpreterCmd.Flags().String("network-configuration", "", "The network configuration for the code interpreter.")
	bedrockAgentcoreControl_createCodeInterpreterCmd.Flags().String("tags", "", "A map of tag keys and values to assign to the code interpreter.")
	bedrockAgentcoreControl_createCodeInterpreterCmd.MarkFlagRequired("name")
	bedrockAgentcoreControl_createCodeInterpreterCmd.MarkFlagRequired("network-configuration")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_createCodeInterpreterCmd)
}
