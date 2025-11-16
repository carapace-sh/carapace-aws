package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_deleteFlowTemplateCmd = &cobra.Command{
	Use:   "delete-flow-template",
	Short: "Deletes a workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_deleteFlowTemplateCmd).Standalone()

	iotthingsgraph_deleteFlowTemplateCmd.Flags().String("id", "", "The ID of the workflow to be deleted.")
	iotthingsgraph_deleteFlowTemplateCmd.MarkFlagRequired("id")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_deleteFlowTemplateCmd)
}
