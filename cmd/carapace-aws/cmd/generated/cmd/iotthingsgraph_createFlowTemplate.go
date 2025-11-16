package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_createFlowTemplateCmd = &cobra.Command{
	Use:   "create-flow-template",
	Short: "Creates a workflow template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_createFlowTemplateCmd).Standalone()

	iotthingsgraph_createFlowTemplateCmd.Flags().String("compatible-namespace-version", "", "The namespace version in which the workflow is to be created.")
	iotthingsgraph_createFlowTemplateCmd.Flags().String("definition", "", "The workflow `DefinitionDocument`.")
	iotthingsgraph_createFlowTemplateCmd.MarkFlagRequired("definition")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_createFlowTemplateCmd)
}
