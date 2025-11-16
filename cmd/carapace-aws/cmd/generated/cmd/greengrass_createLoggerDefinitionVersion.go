package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createLoggerDefinitionVersionCmd = &cobra.Command{
	Use:   "create-logger-definition-version",
	Short: "Creates a version of a logger definition that has already been defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createLoggerDefinitionVersionCmd).Standalone()

	greengrass_createLoggerDefinitionVersionCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
	greengrass_createLoggerDefinitionVersionCmd.Flags().String("logger-definition-id", "", "The ID of the logger definition.")
	greengrass_createLoggerDefinitionVersionCmd.Flags().String("loggers", "", "A list of loggers.")
	greengrass_createLoggerDefinitionVersionCmd.MarkFlagRequired("logger-definition-id")
	greengrassCmd.AddCommand(greengrass_createLoggerDefinitionVersionCmd)
}
