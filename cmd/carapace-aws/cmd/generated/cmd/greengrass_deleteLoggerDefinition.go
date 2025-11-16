package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_deleteLoggerDefinitionCmd = &cobra.Command{
	Use:   "delete-logger-definition",
	Short: "Deletes a logger definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_deleteLoggerDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_deleteLoggerDefinitionCmd).Standalone()

		greengrass_deleteLoggerDefinitionCmd.Flags().String("logger-definition-id", "", "The ID of the logger definition.")
		greengrass_deleteLoggerDefinitionCmd.MarkFlagRequired("logger-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_deleteLoggerDefinitionCmd)
}
