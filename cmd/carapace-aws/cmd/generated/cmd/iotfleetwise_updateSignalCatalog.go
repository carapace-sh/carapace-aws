package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_updateSignalCatalogCmd = &cobra.Command{
	Use:   "update-signal-catalog",
	Short: "Updates a signal catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_updateSignalCatalogCmd).Standalone()

	iotfleetwise_updateSignalCatalogCmd.Flags().String("description", "", "A brief description of the signal catalog to update.")
	iotfleetwise_updateSignalCatalogCmd.Flags().String("name", "", "The name of the signal catalog to update.")
	iotfleetwise_updateSignalCatalogCmd.Flags().String("nodes-to-add", "", "A list of information about nodes to add to the signal catalog.")
	iotfleetwise_updateSignalCatalogCmd.Flags().String("nodes-to-remove", "", "A list of `fullyQualifiedName` of nodes to remove from the signal catalog.")
	iotfleetwise_updateSignalCatalogCmd.Flags().String("nodes-to-update", "", "A list of information about nodes to update in the signal catalog.")
	iotfleetwise_updateSignalCatalogCmd.MarkFlagRequired("name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_updateSignalCatalogCmd)
}
