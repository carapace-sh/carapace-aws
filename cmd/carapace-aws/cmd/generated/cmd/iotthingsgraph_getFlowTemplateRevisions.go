package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_getFlowTemplateRevisionsCmd = &cobra.Command{
	Use:   "get-flow-template-revisions",
	Short: "Gets revisions of the specified workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_getFlowTemplateRevisionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotthingsgraph_getFlowTemplateRevisionsCmd).Standalone()

		iotthingsgraph_getFlowTemplateRevisionsCmd.Flags().String("id", "", "The ID of the workflow.")
		iotthingsgraph_getFlowTemplateRevisionsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		iotthingsgraph_getFlowTemplateRevisionsCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
		iotthingsgraph_getFlowTemplateRevisionsCmd.MarkFlagRequired("id")
	})
	iotthingsgraphCmd.AddCommand(iotthingsgraph_getFlowTemplateRevisionsCmd)
}
