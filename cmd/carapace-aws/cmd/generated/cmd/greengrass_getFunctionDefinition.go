package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getFunctionDefinitionCmd = &cobra.Command{
	Use:   "get-function-definition",
	Short: "Retrieves information about a Lambda function definition, including its creation time and latest version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getFunctionDefinitionCmd).Standalone()

	greengrass_getFunctionDefinitionCmd.Flags().String("function-definition-id", "", "The ID of the Lambda function definition.")
	greengrass_getFunctionDefinitionCmd.MarkFlagRequired("function-definition-id")
	greengrassCmd.AddCommand(greengrass_getFunctionDefinitionCmd)
}
