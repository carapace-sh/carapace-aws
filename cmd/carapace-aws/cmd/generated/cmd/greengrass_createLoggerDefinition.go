package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createLoggerDefinitionCmd = &cobra.Command{
	Use:   "create-logger-definition",
	Short: "Creates a logger definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createLoggerDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_createLoggerDefinitionCmd).Standalone()

		greengrass_createLoggerDefinitionCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
		greengrass_createLoggerDefinitionCmd.Flags().String("initial-version", "", "Information about the initial version of the logger definition.")
		greengrass_createLoggerDefinitionCmd.Flags().String("name", "", "The name of the logger definition.")
		greengrass_createLoggerDefinitionCmd.Flags().String("tags", "", "Tag(s) to add to the new resource.")
	})
	greengrassCmd.AddCommand(greengrass_createLoggerDefinitionCmd)
}
