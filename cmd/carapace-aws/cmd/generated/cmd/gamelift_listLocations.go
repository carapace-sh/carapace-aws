package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_listLocationsCmd = &cobra.Command{
	Use:   "list-locations",
	Short: "**This API works with the following fleet types:** Anywhere\n\nLists all custom and Amazon Web Services locations where Amazon GameLift Servers can host game servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_listLocationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_listLocationsCmd).Standalone()

		gamelift_listLocationsCmd.Flags().String("filters", "", "Filters the list for `AWS` or `CUSTOM` locations.")
		gamelift_listLocationsCmd.Flags().String("limit", "", "The maximum number of results to return.")
		gamelift_listLocationsCmd.Flags().String("next-token", "", "A token that indicates the start of the next sequential page of results.")
	})
	gameliftCmd.AddCommand(gamelift_listLocationsCmd)
}
