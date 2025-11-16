package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createFunctionDefinitionCmd = &cobra.Command{
	Use:   "create-function-definition",
	Short: "Creates a Lambda function definition which contains a list of Lambda functions and their configurations to be used in a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createFunctionDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_createFunctionDefinitionCmd).Standalone()

		greengrass_createFunctionDefinitionCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
		greengrass_createFunctionDefinitionCmd.Flags().String("initial-version", "", "Information about the initial version of the function definition.")
		greengrass_createFunctionDefinitionCmd.Flags().String("name", "", "The name of the function definition.")
		greengrass_createFunctionDefinitionCmd.Flags().String("tags", "", "Tag(s) to add to the new resource.")
	})
	greengrassCmd.AddCommand(greengrass_createFunctionDefinitionCmd)
}
