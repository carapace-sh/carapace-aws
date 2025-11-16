package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_removeEntityOwnerCmd = &cobra.Command{
	Use:   "remove-entity-owner",
	Short: "Removes an owner from an entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_removeEntityOwnerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_removeEntityOwnerCmd).Standalone()

		datazone_removeEntityOwnerCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_removeEntityOwnerCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to remove an owner from an entity.")
		datazone_removeEntityOwnerCmd.Flags().String("entity-identifier", "", "The ID of the entity from which you want to remove an owner.")
		datazone_removeEntityOwnerCmd.Flags().String("entity-type", "", "The type of the entity from which you want to remove an owner.")
		datazone_removeEntityOwnerCmd.Flags().String("owner", "", "The owner that you want to remove from an entity.")
		datazone_removeEntityOwnerCmd.MarkFlagRequired("domain-identifier")
		datazone_removeEntityOwnerCmd.MarkFlagRequired("entity-identifier")
		datazone_removeEntityOwnerCmd.MarkFlagRequired("entity-type")
		datazone_removeEntityOwnerCmd.MarkFlagRequired("owner")
	})
	datazoneCmd.AddCommand(datazone_removeEntityOwnerCmd)
}
