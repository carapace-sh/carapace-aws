package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listFunctionDefinitionVersionsCmd = &cobra.Command{
	Use:   "list-function-definition-versions",
	Short: "Lists the versions of a Lambda function definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listFunctionDefinitionVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_listFunctionDefinitionVersionsCmd).Standalone()

		greengrass_listFunctionDefinitionVersionsCmd.Flags().String("function-definition-id", "", "The ID of the Lambda function definition.")
		greengrass_listFunctionDefinitionVersionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		greengrass_listFunctionDefinitionVersionsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
		greengrass_listFunctionDefinitionVersionsCmd.MarkFlagRequired("function-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_listFunctionDefinitionVersionsCmd)
}
