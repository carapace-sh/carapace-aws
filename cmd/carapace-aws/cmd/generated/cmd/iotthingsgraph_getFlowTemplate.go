package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_getFlowTemplateCmd = &cobra.Command{
	Use:   "get-flow-template",
	Short: "Gets the latest version of the `DefinitionDocument` and `FlowTemplateSummary` for the specified workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_getFlowTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotthingsgraph_getFlowTemplateCmd).Standalone()

		iotthingsgraph_getFlowTemplateCmd.Flags().String("id", "", "The ID of the workflow.")
		iotthingsgraph_getFlowTemplateCmd.Flags().String("revision-number", "", "The number of the workflow revision to retrieve.")
		iotthingsgraph_getFlowTemplateCmd.MarkFlagRequired("id")
	})
	iotthingsgraphCmd.AddCommand(iotthingsgraph_getFlowTemplateCmd)
}
