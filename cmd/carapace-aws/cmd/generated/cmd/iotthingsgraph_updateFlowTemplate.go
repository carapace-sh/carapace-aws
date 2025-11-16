package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_updateFlowTemplateCmd = &cobra.Command{
	Use:   "update-flow-template",
	Short: "Updates the specified workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_updateFlowTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotthingsgraph_updateFlowTemplateCmd).Standalone()

		iotthingsgraph_updateFlowTemplateCmd.Flags().String("compatible-namespace-version", "", "The version of the user's namespace.")
		iotthingsgraph_updateFlowTemplateCmd.Flags().String("definition", "", "The `DefinitionDocument` that contains the updated workflow definition.")
		iotthingsgraph_updateFlowTemplateCmd.Flags().String("id", "", "The ID of the workflow to be updated.")
		iotthingsgraph_updateFlowTemplateCmd.MarkFlagRequired("definition")
		iotthingsgraph_updateFlowTemplateCmd.MarkFlagRequired("id")
	})
	iotthingsgraphCmd.AddCommand(iotthingsgraph_updateFlowTemplateCmd)
}
