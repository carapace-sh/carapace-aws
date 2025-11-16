package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_deprecateFlowTemplateCmd = &cobra.Command{
	Use:   "deprecate-flow-template",
	Short: "Deprecates the specified workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_deprecateFlowTemplateCmd).Standalone()

	iotthingsgraph_deprecateFlowTemplateCmd.Flags().String("id", "", "The ID of the workflow to be deleted.")
	iotthingsgraph_deprecateFlowTemplateCmd.MarkFlagRequired("id")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_deprecateFlowTemplateCmd)
}
