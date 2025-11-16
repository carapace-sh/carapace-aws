package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_addEntityOwnerCmd = &cobra.Command{
	Use:   "add-entity-owner",
	Short: "Adds the owner of an entity (a domain unit).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_addEntityOwnerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_addEntityOwnerCmd).Standalone()

		datazone_addEntityOwnerCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_addEntityOwnerCmd.Flags().String("domain-identifier", "", "The ID of the domain in which you want to add the entity owner.")
		datazone_addEntityOwnerCmd.Flags().String("entity-identifier", "", "The ID of the entity to which you want to add an owner.")
		datazone_addEntityOwnerCmd.Flags().String("entity-type", "", "The type of an entity.")
		datazone_addEntityOwnerCmd.Flags().String("owner", "", "The owner that you want to add to the entity.")
		datazone_addEntityOwnerCmd.MarkFlagRequired("domain-identifier")
		datazone_addEntityOwnerCmd.MarkFlagRequired("entity-identifier")
		datazone_addEntityOwnerCmd.MarkFlagRequired("entity-type")
		datazone_addEntityOwnerCmd.MarkFlagRequired("owner")
	})
	datazoneCmd.AddCommand(datazone_addEntityOwnerCmd)
}
