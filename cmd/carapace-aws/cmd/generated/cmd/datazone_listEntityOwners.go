package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listEntityOwnersCmd = &cobra.Command{
	Use:   "list-entity-owners",
	Short: "Lists the entity (domain units) owners.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listEntityOwnersCmd).Standalone()

	datazone_listEntityOwnersCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to list entity owners.")
	datazone_listEntityOwnersCmd.Flags().String("entity-identifier", "", "The ID of the entity that you want to list.")
	datazone_listEntityOwnersCmd.Flags().String("entity-type", "", "The type of the entity that you want to list.")
	datazone_listEntityOwnersCmd.Flags().String("max-results", "", "The maximum number of entities to return in a single call to `ListEntityOwners`.")
	datazone_listEntityOwnersCmd.Flags().String("next-token", "", "When the number of entities is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of entities, the response includes a pagination token named `NextToken`.")
	datazone_listEntityOwnersCmd.MarkFlagRequired("domain-identifier")
	datazone_listEntityOwnersCmd.MarkFlagRequired("entity-identifier")
	datazone_listEntityOwnersCmd.MarkFlagRequired("entity-type")
	datazoneCmd.AddCommand(datazone_listEntityOwnersCmd)
}
