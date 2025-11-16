package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listConnectorDefinitionsCmd = &cobra.Command{
	Use:   "list-connector-definitions",
	Short: "Retrieves a list of connector definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listConnectorDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_listConnectorDefinitionsCmd).Standalone()

		greengrass_listConnectorDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		greengrass_listConnectorDefinitionsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	})
	greengrassCmd.AddCommand(greengrass_listConnectorDefinitionsCmd)
}
