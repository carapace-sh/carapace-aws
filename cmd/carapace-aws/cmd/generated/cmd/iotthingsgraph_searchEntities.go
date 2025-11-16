package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_searchEntitiesCmd = &cobra.Command{
	Use:   "search-entities",
	Short: "Searches for entities of the specified type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_searchEntitiesCmd).Standalone()

	iotthingsgraph_searchEntitiesCmd.Flags().String("entity-types", "", "The entity types for which to search.")
	iotthingsgraph_searchEntitiesCmd.Flags().String("filters", "", "Optional filter to apply to the search.")
	iotthingsgraph_searchEntitiesCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	iotthingsgraph_searchEntitiesCmd.Flags().String("namespace-version", "", "The version of the user's namespace.")
	iotthingsgraph_searchEntitiesCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
	iotthingsgraph_searchEntitiesCmd.MarkFlagRequired("entity-types")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_searchEntitiesCmd)
}
