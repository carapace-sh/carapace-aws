package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_updateFunctionDefinitionCmd = &cobra.Command{
	Use:   "update-function-definition",
	Short: "Updates a Lambda function definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_updateFunctionDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_updateFunctionDefinitionCmd).Standalone()

		greengrass_updateFunctionDefinitionCmd.Flags().String("function-definition-id", "", "The ID of the Lambda function definition.")
		greengrass_updateFunctionDefinitionCmd.Flags().String("name", "", "The name of the definition.")
		greengrass_updateFunctionDefinitionCmd.MarkFlagRequired("function-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_updateFunctionDefinitionCmd)
}
