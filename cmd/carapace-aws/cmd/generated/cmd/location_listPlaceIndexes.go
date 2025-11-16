package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_listPlaceIndexesCmd = &cobra.Command{
	Use:   "list-place-indexes",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_listPlaceIndexesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_listPlaceIndexesCmd).Standalone()

		location_listPlaceIndexesCmd.Flags().String("max-results", "", "An optional limit for the maximum number of results returned in a single call.")
		location_listPlaceIndexesCmd.Flags().String("next-token", "", "The pagination token specifying which page of results to return in the response.")
	})
	locationCmd.AddCommand(location_listPlaceIndexesCmd)
}
