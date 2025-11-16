package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_importSignalCatalogCmd = &cobra.Command{
	Use:   "import-signal-catalog",
	Short: "Creates a signal catalog using your existing VSS formatted content from your local device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_importSignalCatalogCmd).Standalone()

	iotfleetwise_importSignalCatalogCmd.Flags().String("description", "", "A brief description of the signal catalog.")
	iotfleetwise_importSignalCatalogCmd.Flags().String("name", "", "The name of the signal catalog to import.")
	iotfleetwise_importSignalCatalogCmd.Flags().String("tags", "", "Metadata that can be used to manage the signal catalog.")
	iotfleetwise_importSignalCatalogCmd.Flags().String("vss", "", "The contents of the Vehicle Signal Specification (VSS) configuration.")
	iotfleetwise_importSignalCatalogCmd.MarkFlagRequired("name")
	iotfleetwiseCmd.AddCommand(iotfleetwise_importSignalCatalogCmd)
}
