package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_listSignalCatalogsCmd = &cobra.Command{
	Use:   "list-signal-catalogs",
	Short: "Lists all the created signal catalogs in an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_listSignalCatalogsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_listSignalCatalogsCmd).Standalone()

		iotfleetwise_listSignalCatalogsCmd.Flags().String("max-results", "", "The maximum number of items to return, between 1 and 100, inclusive.")
		iotfleetwise_listSignalCatalogsCmd.Flags().String("next-token", "", "A pagination token for the next set of results.")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_listSignalCatalogsCmd)
}
