package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_updateSystemTemplateCmd = &cobra.Command{
	Use:   "update-system-template",
	Short: "Updates the specified system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_updateSystemTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotthingsgraph_updateSystemTemplateCmd).Standalone()

		iotthingsgraph_updateSystemTemplateCmd.Flags().String("compatible-namespace-version", "", "The version of the user's namespace.")
		iotthingsgraph_updateSystemTemplateCmd.Flags().String("definition", "", "The `DefinitionDocument` that contains the updated system definition.")
		iotthingsgraph_updateSystemTemplateCmd.Flags().String("id", "", "The ID of the system to be updated.")
		iotthingsgraph_updateSystemTemplateCmd.MarkFlagRequired("definition")
		iotthingsgraph_updateSystemTemplateCmd.MarkFlagRequired("id")
	})
	iotthingsgraphCmd.AddCommand(iotthingsgraph_updateSystemTemplateCmd)
}
