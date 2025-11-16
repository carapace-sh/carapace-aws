package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getLoggerDefinitionVersionCmd = &cobra.Command{
	Use:   "get-logger-definition-version",
	Short: "Retrieves information about a logger definition version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getLoggerDefinitionVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_getLoggerDefinitionVersionCmd).Standalone()

		greengrass_getLoggerDefinitionVersionCmd.Flags().String("logger-definition-id", "", "The ID of the logger definition.")
		greengrass_getLoggerDefinitionVersionCmd.Flags().String("logger-definition-version-id", "", "The ID of the logger definition version.")
		greengrass_getLoggerDefinitionVersionCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
		greengrass_getLoggerDefinitionVersionCmd.MarkFlagRequired("logger-definition-id")
		greengrass_getLoggerDefinitionVersionCmd.MarkFlagRequired("logger-definition-version-id")
	})
	greengrassCmd.AddCommand(greengrass_getLoggerDefinitionVersionCmd)
}
