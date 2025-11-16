package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_updateLoggerDefinitionCmd = &cobra.Command{
	Use:   "update-logger-definition",
	Short: "Updates a logger definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_updateLoggerDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_updateLoggerDefinitionCmd).Standalone()

		greengrass_updateLoggerDefinitionCmd.Flags().String("logger-definition-id", "", "The ID of the logger definition.")
		greengrass_updateLoggerDefinitionCmd.Flags().String("name", "", "The name of the definition.")
		greengrass_updateLoggerDefinitionCmd.MarkFlagRequired("logger-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_updateLoggerDefinitionCmd)
}
