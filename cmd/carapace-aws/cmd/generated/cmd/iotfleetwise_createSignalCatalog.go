package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_createSignalCatalogCmd = &cobra.Command{
	Use:   "create-signal-catalog",
	Short: "Creates a collection of standardized signals that can be reused to create vehicle models.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_createSignalCatalogCmd).Standalone()

	iotfleetwise_createSignalCatalogCmd.Flags().String("description", "", "A brief description of the signal catalog.")
	iotfleetwise_createSignalCatalogCmd.Flags().String("name", "", "The name of the signal catalog to create.")
	iotfleetwise_createSignalCatalogCmd.Flags().String("nodes", "", "A list of information about nodes, which are a general abstraction of signals.")
	iotfleetwise_createSignalCatalogCmd.Flags().String("tags", "", "Metadata that can be used to manage the signal catalog.")
	iotfleetwise_createSignalCatalogCmd.MarkFlagRequired("name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_createSignalCatalogCmd)
}
