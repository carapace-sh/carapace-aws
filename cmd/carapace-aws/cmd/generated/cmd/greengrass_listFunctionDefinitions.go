package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listFunctionDefinitionsCmd = &cobra.Command{
	Use:   "list-function-definitions",
	Short: "Retrieves a list of Lambda function definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listFunctionDefinitionsCmd).Standalone()

	greengrass_listFunctionDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	greengrass_listFunctionDefinitionsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	greengrassCmd.AddCommand(greengrass_listFunctionDefinitionsCmd)
}
