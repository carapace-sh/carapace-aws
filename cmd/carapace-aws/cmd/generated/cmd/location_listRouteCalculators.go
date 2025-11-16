package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_listRouteCalculatorsCmd = &cobra.Command{
	Use:   "list-route-calculators",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_listRouteCalculatorsCmd).Standalone()

	location_listRouteCalculatorsCmd.Flags().String("max-results", "", "An optional maximum number of results returned in a single call.")
	location_listRouteCalculatorsCmd.Flags().String("next-token", "", "The pagination token specifying which page of results to return in the response.")
	locationCmd.AddCommand(location_listRouteCalculatorsCmd)
}
