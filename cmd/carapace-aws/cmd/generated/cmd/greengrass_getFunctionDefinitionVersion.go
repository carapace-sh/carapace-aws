package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getFunctionDefinitionVersionCmd = &cobra.Command{
	Use:   "get-function-definition-version",
	Short: "Retrieves information about a Lambda function definition version, including which Lambda functions are included in the version and their configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getFunctionDefinitionVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_getFunctionDefinitionVersionCmd).Standalone()

		greengrass_getFunctionDefinitionVersionCmd.Flags().String("function-definition-id", "", "The ID of the Lambda function definition.")
		greengrass_getFunctionDefinitionVersionCmd.Flags().String("function-definition-version-id", "", "The ID of the function definition version.")
		greengrass_getFunctionDefinitionVersionCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
		greengrass_getFunctionDefinitionVersionCmd.MarkFlagRequired("function-definition-id")
		greengrass_getFunctionDefinitionVersionCmd.MarkFlagRequired("function-definition-version-id")
	})
	greengrassCmd.AddCommand(greengrass_getFunctionDefinitionVersionCmd)
}
