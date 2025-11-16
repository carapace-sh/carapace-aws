package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listCoreDefinitionsCmd = &cobra.Command{
	Use:   "list-core-definitions",
	Short: "Retrieves a list of core definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listCoreDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_listCoreDefinitionsCmd).Standalone()

		greengrass_listCoreDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		greengrass_listCoreDefinitionsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	})
	greengrassCmd.AddCommand(greengrass_listCoreDefinitionsCmd)
}
