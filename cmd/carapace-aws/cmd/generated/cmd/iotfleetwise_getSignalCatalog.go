package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_getSignalCatalogCmd = &cobra.Command{
	Use:   "get-signal-catalog",
	Short: "Retrieves information about a signal catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_getSignalCatalogCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_getSignalCatalogCmd).Standalone()

		iotfleetwise_getSignalCatalogCmd.Flags().String("name", "", "The name of the signal catalog to retrieve information about.")
		iotfleetwise_getSignalCatalogCmd.MarkFlagRequired("name")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_getSignalCatalogCmd)
}
