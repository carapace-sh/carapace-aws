package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_searchThingsCmd = &cobra.Command{
	Use:   "search-things",
	Short: "Searches for things associated with the specified entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_searchThingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotthingsgraph_searchThingsCmd).Standalone()

		iotthingsgraph_searchThingsCmd.Flags().String("entity-id", "", "The ID of the entity to which the things are associated.")
		iotthingsgraph_searchThingsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		iotthingsgraph_searchThingsCmd.Flags().String("namespace-version", "", "The version of the user's namespace.")
		iotthingsgraph_searchThingsCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
		iotthingsgraph_searchThingsCmd.MarkFlagRequired("entity-id")
	})
	iotthingsgraphCmd.AddCommand(iotthingsgraph_searchThingsCmd)
}
