package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listLoggerDefinitionsCmd = &cobra.Command{
	Use:   "list-logger-definitions",
	Short: "Retrieves a list of logger definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listLoggerDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_listLoggerDefinitionsCmd).Standalone()

		greengrass_listLoggerDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		greengrass_listLoggerDefinitionsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	})
	greengrassCmd.AddCommand(greengrass_listLoggerDefinitionsCmd)
}
