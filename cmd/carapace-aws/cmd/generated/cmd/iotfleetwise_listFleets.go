package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_listFleetsCmd = &cobra.Command{
	Use:   "list-fleets",
	Short: "Retrieves information for each created fleet in an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_listFleetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_listFleetsCmd).Standalone()

		iotfleetwise_listFleetsCmd.Flags().String("list-response-scope", "", "When you set the `listResponseScope` parameter to `METADATA_ONLY`, the list response includes: fleet ID, Amazon Resource Name (ARN), creation time, and last modification time.")
		iotfleetwise_listFleetsCmd.Flags().String("max-results", "", "The maximum number of items to return, between 1 and 100, inclusive.")
		iotfleetwise_listFleetsCmd.Flags().String("next-token", "", "A pagination token for the next set of results.")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_listFleetsCmd)
}
