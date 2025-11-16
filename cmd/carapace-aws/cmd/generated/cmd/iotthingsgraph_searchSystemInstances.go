package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_searchSystemInstancesCmd = &cobra.Command{
	Use:   "search-system-instances",
	Short: "Searches for system instances in the user's account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_searchSystemInstancesCmd).Standalone()

	iotthingsgraph_searchSystemInstancesCmd.Flags().String("filters", "", "Optional filter to apply to the search.")
	iotthingsgraph_searchSystemInstancesCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	iotthingsgraph_searchSystemInstancesCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_searchSystemInstancesCmd)
}
