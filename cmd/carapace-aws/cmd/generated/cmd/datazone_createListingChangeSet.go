package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createListingChangeSetCmd = &cobra.Command{
	Use:   "create-listing-change-set",
	Short: "Publishes a listing (a record of an asset at a given time) or removes a listing from the catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createListingChangeSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createListingChangeSetCmd).Standalone()

		datazone_createListingChangeSetCmd.Flags().String("action", "", "Specifies whether to publish or unpublish a listing.")
		datazone_createListingChangeSetCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_createListingChangeSetCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain.")
		datazone_createListingChangeSetCmd.Flags().String("entity-identifier", "", "The ID of the asset.")
		datazone_createListingChangeSetCmd.Flags().String("entity-revision", "", "The revision of an asset.")
		datazone_createListingChangeSetCmd.Flags().String("entity-type", "", "The type of an entity.")
		datazone_createListingChangeSetCmd.MarkFlagRequired("action")
		datazone_createListingChangeSetCmd.MarkFlagRequired("domain-identifier")
		datazone_createListingChangeSetCmd.MarkFlagRequired("entity-identifier")
		datazone_createListingChangeSetCmd.MarkFlagRequired("entity-type")
	})
	datazoneCmd.AddCommand(datazone_createListingChangeSetCmd)
}
