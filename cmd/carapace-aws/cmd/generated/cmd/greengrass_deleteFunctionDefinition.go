package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_deleteFunctionDefinitionCmd = &cobra.Command{
	Use:   "delete-function-definition",
	Short: "Deletes a Lambda function definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_deleteFunctionDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_deleteFunctionDefinitionCmd).Standalone()

		greengrass_deleteFunctionDefinitionCmd.Flags().String("function-definition-id", "", "The ID of the Lambda function definition.")
		greengrass_deleteFunctionDefinitionCmd.MarkFlagRequired("function-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_deleteFunctionDefinitionCmd)
}
