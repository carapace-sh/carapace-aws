package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_deleteSignalCatalogCmd = &cobra.Command{
	Use:   "delete-signal-catalog",
	Short: "Deletes a signal catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_deleteSignalCatalogCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_deleteSignalCatalogCmd).Standalone()

		iotfleetwise_deleteSignalCatalogCmd.Flags().String("name", "", "The name of the signal catalog to delete.")
		iotfleetwise_deleteSignalCatalogCmd.MarkFlagRequired("name")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_deleteSignalCatalogCmd)
}
