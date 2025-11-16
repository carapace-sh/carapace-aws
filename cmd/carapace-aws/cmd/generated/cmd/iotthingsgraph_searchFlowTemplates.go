package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_searchFlowTemplatesCmd = &cobra.Command{
	Use:   "search-flow-templates",
	Short: "Searches for summary information about workflows.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_searchFlowTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotthingsgraph_searchFlowTemplatesCmd).Standalone()

		iotthingsgraph_searchFlowTemplatesCmd.Flags().String("filters", "", "An array of objects that limit the result set.")
		iotthingsgraph_searchFlowTemplatesCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		iotthingsgraph_searchFlowTemplatesCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
	})
	iotthingsgraphCmd.AddCommand(iotthingsgraph_searchFlowTemplatesCmd)
}
