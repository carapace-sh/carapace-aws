package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_listMapsCmd = &cobra.Command{
	Use:   "list-maps",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_listMapsCmd).Standalone()

	location_listMapsCmd.Flags().String("max-results", "", "An optional limit for the number of resources returned in a single call.")
	location_listMapsCmd.Flags().String("next-token", "", "The pagination token specifying which page of results to return in the response.")
	locationCmd.AddCommand(location_listMapsCmd)
}
