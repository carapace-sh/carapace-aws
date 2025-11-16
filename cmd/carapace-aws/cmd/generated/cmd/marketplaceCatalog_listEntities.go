package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceCatalog_listEntitiesCmd = &cobra.Command{
	Use:   "list-entities",
	Short: "Provides the list of entities of a given type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceCatalog_listEntitiesCmd).Standalone()

	marketplaceCatalog_listEntitiesCmd.Flags().String("catalog", "", "The catalog related to the request.")
	marketplaceCatalog_listEntitiesCmd.Flags().String("entity-type", "", "The type of entities to retrieve.")
	marketplaceCatalog_listEntitiesCmd.Flags().String("entity-type-filters", "", "A Union object containing filter shapes for all `EntityType`s.")
	marketplaceCatalog_listEntitiesCmd.Flags().String("entity-type-sort", "", "A Union object containing `Sort` shapes for all `EntityType`s.")
	marketplaceCatalog_listEntitiesCmd.Flags().String("filter-list", "", "An array of filter objects.")
	marketplaceCatalog_listEntitiesCmd.Flags().String("max-results", "", "Specifies the upper limit of the elements on a single page.")
	marketplaceCatalog_listEntitiesCmd.Flags().String("next-token", "", "The value of the next token, if it exists.")
	marketplaceCatalog_listEntitiesCmd.Flags().String("ownership-type", "", "Filters the returned set of entities based on their owner.")
	marketplaceCatalog_listEntitiesCmd.Flags().String("sort", "", "An object that contains two attributes, `SortBy` and `SortOrder`.")
	marketplaceCatalog_listEntitiesCmd.MarkFlagRequired("catalog")
	marketplaceCatalog_listEntitiesCmd.MarkFlagRequired("entity-type")
	marketplaceCatalogCmd.AddCommand(marketplaceCatalog_listEntitiesCmd)
}
