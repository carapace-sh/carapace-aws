package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_listTrackersCmd = &cobra.Command{
	Use:   "list-trackers",
	Short: "Lists tracker resources in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_listTrackersCmd).Standalone()

	location_listTrackersCmd.Flags().String("max-results", "", "An optional limit for the number of resources returned in a single call.")
	location_listTrackersCmd.Flags().String("next-token", "", "The pagination token specifying which page of results to return in the response.")
	locationCmd.AddCommand(location_listTrackersCmd)
}
