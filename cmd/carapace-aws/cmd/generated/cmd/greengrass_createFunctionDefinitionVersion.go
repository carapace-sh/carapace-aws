package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createFunctionDefinitionVersionCmd = &cobra.Command{
	Use:   "create-function-definition-version",
	Short: "Creates a version of a Lambda function definition that has already been defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createFunctionDefinitionVersionCmd).Standalone()

	greengrass_createFunctionDefinitionVersionCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
	greengrass_createFunctionDefinitionVersionCmd.Flags().String("default-config", "", "The default configuration that applies to all Lambda functions in this function definition version.")
	greengrass_createFunctionDefinitionVersionCmd.Flags().String("function-definition-id", "", "The ID of the Lambda function definition.")
	greengrass_createFunctionDefinitionVersionCmd.Flags().String("functions", "", "A list of Lambda functions in this function definition version.")
	greengrass_createFunctionDefinitionVersionCmd.MarkFlagRequired("function-definition-id")
	greengrassCmd.AddCommand(greengrass_createFunctionDefinitionVersionCmd)
}
