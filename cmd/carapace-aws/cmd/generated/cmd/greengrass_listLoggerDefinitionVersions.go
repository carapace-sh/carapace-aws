package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listLoggerDefinitionVersionsCmd = &cobra.Command{
	Use:   "list-logger-definition-versions",
	Short: "Lists the versions of a logger definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listLoggerDefinitionVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_listLoggerDefinitionVersionsCmd).Standalone()

		greengrass_listLoggerDefinitionVersionsCmd.Flags().String("logger-definition-id", "", "The ID of the logger definition.")
		greengrass_listLoggerDefinitionVersionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
		greengrass_listLoggerDefinitionVersionsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
		greengrass_listLoggerDefinitionVersionsCmd.MarkFlagRequired("logger-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_listLoggerDefinitionVersionsCmd)
}
