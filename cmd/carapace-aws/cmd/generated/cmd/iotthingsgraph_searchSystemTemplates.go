package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_searchSystemTemplatesCmd = &cobra.Command{
	Use:   "search-system-templates",
	Short: "Searches for summary information about systems in the user's account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_searchSystemTemplatesCmd).Standalone()

	iotthingsgraph_searchSystemTemplatesCmd.Flags().String("filters", "", "An array of filters that limit the result set.")
	iotthingsgraph_searchSystemTemplatesCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	iotthingsgraph_searchSystemTemplatesCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_searchSystemTemplatesCmd)
}
