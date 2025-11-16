package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getLoggerDefinitionCmd = &cobra.Command{
	Use:   "get-logger-definition",
	Short: "Retrieves information about a logger definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getLoggerDefinitionCmd).Standalone()

	greengrass_getLoggerDefinitionCmd.Flags().String("logger-definition-id", "", "The ID of the logger definition.")
	greengrass_getLoggerDefinitionCmd.MarkFlagRequired("logger-definition-id")
	greengrassCmd.AddCommand(greengrass_getLoggerDefinitionCmd)
}
